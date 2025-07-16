package call

import "github.com/gorilla/websocket"

type CallHub struct {
	clients   map[int]*CallClient
	broadcast chan SignalMsg
	register  chan *CallClient
	unreg     chan *CallClient
}

type CallClient struct {
	uid  int
	conn *websocket.Conn
	hub  *CallHub
}

func NewCallHub() *CallHub {
	h := &CallHub{
		clients:   make(map[int]*CallClient),
		broadcast: make(chan SignalMsg),
		register:  make(chan *CallClient),
		unreg:     make(chan *CallClient),
	}
	go h.run()
	return h
}

func (h *CallHub) run() {
	for {
		select {
		case c := <-h.register:
			h.clients[c.uid] = c
		case c := <-h.unreg:
			delete(h.clients, c.uid)
		case m := <-h.broadcast:
			if cli, ok := h.clients[m.To]; ok {
				cli.conn.WriteJSON(m)
			}
		}
	}
}

func (c *CallClient) readLoop() {
	for {
		var msg SignalMsg
		if err := c.conn.ReadJSON(&msg); err != nil {
			c.hub.unreg <- c
			break
		}
		msg.From = c.uid
		c.hub.broadcast <- msg
	}
}
