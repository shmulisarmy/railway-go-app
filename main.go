package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // Switch to release mode
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := "8000"
	// port := os.Getenv("PORT")

	r.GET("/port", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"port": port,
		})
	})

	// @Summary      Get all todos
	// @Description  get all todos
	// @Tags         todos
	// @Accept       json
	// @Produce      json
	// @Success      200  {object}  []Todo
	// @Failure      400  {object}  ErrorResponse
	// @Router       /todos [get]

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

	for _, route := range r.Routes() {
		fmt.Printf("%s %s\n", route.Method, route.Path)
		fmt.Printf("%s %s\n", route.Handler, route.Path)
	}
	print("http://localhost:" + port)

	r.Run("0.0.0.0:" + port) // listen and serve on 0.0.0.0:8080
}
