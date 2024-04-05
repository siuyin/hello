package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/siuyin/dflt"
)

func main() {
	http.HandleFunc("/ws", wsHandler)
	// exercise color API with:
	// curl -X POST -d'{"color":"green"}' http://localhost:8080/color
	// curl -X POST -d'{"color":"red"}' http://localhost:8080/color
	http.HandleFunc("/color", colorHandler)
	http.Handle("/", http.FileServer(http.Dir("./internal/public")))
	// http.Handle("/", http.FileServer(http.FS(public.Content)))

	log.Fatal(http.ListenAndServe(":"+dflt.EnvString("HTTP_PORT", "8080"), nil))
}

type colorChange struct {
	Color string `json:"color"`
}

var currentColor = "grey"
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var connections = make(map[*websocket.Conn]bool)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	connections[conn] = true
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			delete(connections, conn)
			break
		}
		if messageType == websocket.CloseMessage {
			break
		}
		log.Printf("msg: %s", message)
	}
}

func colorHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var change colorChange
	err := decoder.Decode(&change)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid color request: %v", err)
		return
	}


	currentColor = change.Color
	broadcastColor()
}

func broadcastColor() {
	for conn := range connections {
		err := conn.WriteMessage(websocket.TextMessage, []byte(currentColor))
		if err != nil {
			delete(connections, conn)
			log.Println(err)
		}
	}
}
