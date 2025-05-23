package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	liveDbSync "railway-go-app/db_live_sync"
	"railway-go-app/env"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
		// return []string{"http://localhost:3000"}.Contains(r.Header.Get("Origin"))
	},
}

var clients = make([]*websocket.Conn, 100)
var people_store liveDbSync.LiveDbSync

var db_conn *pgx.Conn

func init() {

	var err error
	err = env.Load_env(".env")
	if err != nil {
		panic(err)
	}

	db_conn, err = pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	liveDbSync.Config(db_conn)

	people_store = liveDbSync.LiveDbSync{
		Query: "select id, name, image, gender, is_descendant, parent_id, spouse_id from person where removed = false",
		Update_query: func(last_update_check float64) string {
			return "select id, name, image, gender, is_descendant, parent_id, spouse_id from person where removed = false and last_updated > " + fmt.Sprintf("%f", last_update_check)
		},
	}
	people_store.Load_data()
}

func wsHandler(r *gin.Context) {
	conn, err := upgrader.Upgrade(r.Writer, r.Request, nil)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": "WebSocket upgrade failed"})
		return
	}
	defer conn.Close()

	clients = append(clients, conn)

	people_store.On_listener_join(conn)

	for {
		_, p, err := conn.ReadMessage()

		if err != nil {
			remove_client(conn)
			break
		}
		broadcast(string(p))
		fmt.Printf("Received: %s\n", p)

	}
}

func remove_client(conn *websocket.Conn) {
	for i, client := range clients {
		if client == conn {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}
}

func broadcast(message string) {
	for _, client := range clients {
		if err := client.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			remove_client(client)
		}
	}
}

func main() {

	gin.SetMode(gin.ReleaseMode) // Switch to release mode

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	r.LoadHTMLGlob("frontend/dist/index.html")
	r.Static("/assets", "frontend/dist/assets")

	r.GET("/ping", func(c *gin.Context) {
		fmt.Printf("this is the ping route")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.File("frontend/dist/index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	r.GET("/port", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"port": port,
		})
	})

	r.GET("/new-route", func(c *gin.Context) {
		fmt.Printf("new route")
		c.JSON(200, gin.H{
			"message": "new route",
		})
	})

	r.GET("/todos", func(c *gin.Context) {
		c.JSON(200, todos)
	})
	r.GET("/todo/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}
		var todo Todo
		for _, todo := range todos {
			if todo.Id == id {
				c.JSON(200, todo)
				return
			}
		}
		c.JSON(200, todo)
	})

	r.PATCH("/todo/:id/completed", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}
		var todo Todo
		for _, todo := range todos {
			if todo.Id == id {
				todo.Completed = !todo.Completed
				c.JSON(200, todo)
				return
			}
		}
		c.JSON(200, todo)
	})

	r.GET("/people", func(c *gin.Context) {
		converted := make(map[string]map[string]interface{})

		for key, value := range people_store.Rows {
			converted[fmt.Sprintf("%v", key)] = value
		}

		c.JSON(200, converted)
	})

	r.GET("/ws", wsHandler) // Don't use WrapH here, just register the handler directly

	// for _, route := range r.Routes() {
	// 	fmt.Printf("%s %s\n", route.Method, route.Path)
	// 	fmt.Printf("%s %s\n", route.Handler, route.Path)
	// }
	print("http://localhost:" + port)

	if err := r.Run("0.0.0.0:" + port); err != nil {
		panic(err)
	}
}
