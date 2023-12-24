package main

import (
	"fmt"
	"net/http"
	// "encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

// Represents a simulated ticket in the system
type Ticket struct {
	Id       int    `json:"id"`
	ClientId string `json:"clientId"`
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// Store key-value pairs of WebSocket connection and the respective connected client
	clients = make(map[*websocket.Conn]string)
	// List of Tickets (ticket information)
	tickets []Ticket
	// Channel to share ticket updates
	ticket_update_channel = make(chan []Ticket)
	// Initialize ticket ID
	next_ticket_id = 1
)

func ws(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, Go Server")

  conn, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    log.Println(err)
    return
  }

}

func main() {
  http.HandleFunc("/ws", ws)
  http.ListenAndServe(":8080", nil)
}


