package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/coder/websocket"
	"github.com/siuyin/dflt"
	"github.com/siuyin/hello/cmd/websocket/internal/public"
)

type colorChange struct {
	Color string `json:"color"`
}

var currentColor = "grey"

var connections = make(map[*websocket.Conn]struct{})

func main() {
	go webServer()

	wsPort := dflt.EnvString("WS_PORT", "8081")
	fmt.Printf("starting websocket server on port %s\n", wsPort)
	log.Fatal(http.ListenAndServe(":"+wsPort, http.HandlerFunc(websocketHandler)))
}

func webServer() {
	http.HandleFunc("/color", colorHandler)
	http.Handle("/", http.FileServer(http.FS(public.Content)))

	port := dflt.EnvString("PORT", "8080")
	fmt.Printf("starting webserver on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r,
		&websocket.AcceptOptions{
			OriginPatterns: []string{"localhost:8080"},
		},
	)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close(websocket.StatusNormalClosure, "normal closure")

	connections[c] = struct{}{}
	for {
		_, message, err := c.Read(context.Background())
		//messageType, message, err := conn.ReadMessage()
		if err != nil {
			delete(connections, c)
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
	log.Println("broadcasting color:", currentColor)
	for conn := range connections {
		err := conn.Write(context.Background(), websocket.MessageText, []byte(currentColor))
		if err != nil {
			delete(connections, conn)
			log.Println(err)
		}
	}
}
