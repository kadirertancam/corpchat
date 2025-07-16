package chat

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn   *websocket.Conn
	userID int
	send   chan []byte
	hub    *Hub
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
	mu         sync.Mutex
}

type Message struct {
	FromUID int    `json:"from_uid"`
	ToUID   int    `json:"to_uid"`
	Body    string `json:"body"`
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()
			log.Printf("user %d connected", client.userID)

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			h.mu.Unlock()

		case msg := <-h.broadcast:
			h.mu.Lock()
			for client := range h.clients {
				if client.userID == msg.ToUID || client.userID == msg.FromUID {
					select {
					case client.send <- encode(msg):
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
			h.mu.Unlock()

			go SaveMessage(db, msg)
		}
	}
}

func encode(v any) []byte {
	b, _ := json.Marshal(v)
	return b
}