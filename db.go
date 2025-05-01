package main

import (
	"context"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5"
)

var db_conn *pgx.Conn

type LiveDbSync struct {
	query           string
	update_query    string
	rows            map[interface{}]map[string]interface{}
	code_listeners  []func(map[string]interface{})
	listeners       []*websocket.Conn
	message_id_upto int
}

func (store *LiveDbSync) load_data() {
	store.rows = make(map[interface{}]map[string]interface{})
	rows, err := ScanRowsToMapSlice(context.Background(), db_conn, store.query)
	if err != nil {
		panic(err)
	}
	for _, row := range rows {
		pk := row["id"]
		store.rows[pk] = row
	}
}

func (store *LiveDbSync) on_listener_join(conn *websocket.Conn) {
	fmt.Println("Listener joined")
	store.listeners = append(store.listeners, conn)
	// message_id_upto doesn't get updated, this user will start keeping track of message_id from this id and on (while users that are here from the begging will keep track all the way from 0 and so they know if their in sync with the server)
	message := map[string]interface{}{
		"type":            "store-join",
		"rows":            store.rows,
		"length":          len(store.rows),
		"message_id_upto": store.message_id_upto,
	}
	fmt.Printf("Sending store-join: %v\n", message)
	conn.WriteJSON(message)
}

func (store *LiveDbSync) on_code_listener_join(code_listener func(map[string]interface{})) {
	store.code_listeners = append(store.code_listeners, code_listener)
}

func (store *LiveDbSync) broadcast_json(message map[string]interface{}) {
	message["message_id_upto"] = store.message_id_upto
	store.message_id_upto++
	for _, conn := range store.listeners {
		err := conn.WriteJSON(message)
		if err != nil {
			remove_client(conn)
		}
	}
	for _, listener := range store.code_listeners {
		listener(message)
	}

}

func (store *LiveDbSync) update_data() {
	rows, err := ScanRowsToMapSlice(context.Background(), db_conn, store.update_query)
	if err != nil {
		panic(err)
	}
	for _, row := range rows {
		pk := row["id"]
		if _, ok := store.rows[pk]; ok {
			store.broadcast_json(map[string]interface{}{
				"type":   "row-updated",
				"row":    row,
				"id":     pk,
				"length": len(store.rows),
			})
		} else {
			store.broadcast_json(map[string]interface{}{
				"type":   "row-added",
				"row":    row,
				"id":     pk,
				"length": len(store.rows),
			})
		}
		store.rows[pk] = row

	}
}
