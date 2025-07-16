package chat

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
)

type Client struct {
	conn   *websocket.Conn
	userID int64
	send   chan []byte
	hub    *Hub
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
	mu         sync.Mutex
	db         *sqlx.DB // burada db ekleniyor
}

func NewHub(db *sqlx.DB) *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		db:         db, // db atanıyor
	}
}
func (h *Hub) BroadcastToChannel(chID int, msg Message) {
	h.mu.Lock()
	for client := range h.clients {
		// tüm kanal üyelerine gönder
		select {
		case client.send <- encode(msg):
		default:
			close(client.send)
			delete(h.clients, client)
		}
	}
	h.mu.Unlock()
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
				if client.userID == int64(msg.ToUID) || client.userID == msg.FromUID {
					select {
					case client.send <- encode(msg):
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
			h.mu.Unlock()

			go func() {
				if err := SaveMessage(h.db, msg); err != nil {
					log.Println("SaveMessage error:", err)
				}
			}()
		}
	}
}

func encode(v any) []byte {
	b, _ := json.Marshal(v)
	return b
}
