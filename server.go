package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

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
	// Ticket ID counter
	next_ticket_id = 1
)

func main() {
	http.HandleFunc("/ws", ws)

	// Display Tickets
	// Set option for new ticket (create_ticket) or quit
	fmt.Println("Server gestartet")
	http.ListenAndServe(":8080", nil)
}

func ws(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go Server")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	go read_client_messages(conn)
}

func main() {
  http.HandleFunc("/ws", ws)
  http.ListenAndServe(":8080", nil)
}


