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
	tickets = make([]Ticket, 0)
	// Channel to share ticket updates
	ticket_update_channel = make(chan []Ticket)
	// Ticket ID counter
	next_ticket_id = 1
)

func main() {
	go broadcast_ticket_updates()
	go interact()

	http.HandleFunc("/ws", ws)

	fmt.Println("Server gestartet")
	http.ListenAndServe(":8080", nil)
}

func interact() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		display_tickets()
		fmt.Println("n: neues Ticket, q: quit")

		if scanner.Scan() {
			input := scanner.Text()

			switch input {
			case "n":
				create_ticket()
			case "q":
				fmt.Println("Server beendet")
				os.Exit(0)
			default:
				fmt.Println("Ungültige Eingabe")
			}
		}
	}
}

func display_tickets() {
	if len(tickets) == 0 {
		fmt.Println("Keine Tickets")
		return
	}

	for _, ticket := range tickets {
		fmt.Printf("Ticket ID: %d, zugewiesen von: %s\n", ticket.Id, ticket.ClientId)
	}
}

func broadcast_ticket_updates() {
	for {
		new_tickets := <-ticket_update_channel
		// Broadcast updated ticket information to all connected clients
		for client := range clients {
			if err := client.WriteJSON(struct {
				MessageType string   `json:"message_type"`
				Tickets     []Ticket `json:"tickets"`
			}{
				MessageType: "ticket_info",
				Tickets:     new_tickets,
			}); err != nil {
				fmt.Println(err)
			}
		}
	}
	// for new_tickets := range ticket_update_channel {
	// 	// Broadcast updated ticket information to all connected clients
	// 	for client := range clients {
	// 		if err := client.WriteJSON(new_tickets); err != nil {
	// 			fmt.Println(err)
	// 		}
	// 	}
	// }
}

func create_ticket() {
	new_ticket := Ticket{
		Id:       next_ticket_id,
		ClientId: "",
	}

	tickets = append(tickets, new_ticket)
	next_ticket_id++

	ticket_update_channel <- tickets
}

func ws(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Send (initial) current tickets to client
	conn.WriteJSON(struct {
		MessageType string   `json:"message_type"`
		Tickets     []Ticket `json:"tickets"`
	}{
		MessageType: "ticket_info",
		Tickets:     tickets,
	})

	go read_client_messages(conn)
}

func main() {
  http.HandleFunc("/ws", ws)
  http.ListenAndServe(":8080", nil)
}


