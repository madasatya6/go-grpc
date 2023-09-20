package wshandler

import (
	"context"
	"encoding/json"
	"strconv"

	"go_grpc"
	"go_grpc/lib/logger"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	Backend *backend.Backend
}

func NewHub(backend *backend.Backend) *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		Backend:      backend,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			found := 0
			for client := range h.clients {
				if client.deviceID == message.deviceID {
					found++
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}

			if found > 0 {
				payloadString, err := json.Marshal(message.payload)
				if err == nil {
					logger.Info(context.Background(), "Broadcast websocket"+", Device ID:"+message.deviceID+", Payload: "+string(payloadString)+", Found:"+strconv.Itoa(found), map[string]interface{}{
						"tags": []string{"websocket", "broadcast"},
					})
				}

				err = h.Backend.Usecase.WSSetPublishedMeeting(context.Background(), message.payload.ID)

				if err != nil {
					logger.Error(context.Background(), "Broadcast websocket, Set published meeting"+", Device ID:"+message.deviceID+", Payload: "+string(payloadString)+", Found:"+strconv.Itoa(found), map[string]interface{}{
						"tags": []string{"websocket", "broadcast"},
					})
				}
			} else {
				payloadString, err := json.Marshal(message.payload)
				if err == nil {
					logger.Error(context.Background(), "Broadcast websocket"+", Device ID:"+message.deviceID+", Payload: "+string(payloadString)+", Found:"+strconv.Itoa(found), map[string]interface{}{
						"tags": []string{"websocket", "broadcast"},
					})
				}
			}
		}
	}
}
