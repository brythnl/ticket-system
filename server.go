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
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
  }
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


