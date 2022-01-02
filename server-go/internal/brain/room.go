package brain

import (
	"fmt"
	"github.com/google/uuid"
)

// Room maintains the set of active clients and broadcasts messages to the
// clients.
type Room struct {
	info     *RoomID
	messages Messages

	chatter Chatter

	// Inbound messages from the clients.
	broadcast chan *Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewRoom(name string, roomType string, chatter Chatter) *Room {
	return &Room{
		info: &RoomID{
			ID: ID{
				UUID: uuid.NewString(),
				Name: name,
			},
			Type: Type(roomType),
		},
		messages:   NewMessages(),
		broadcast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		chatter:    chatter,
	}
}

func (h *Room) Run() {
	fmt.Println("Starting room ", h.info.Name)
	for {
		select {
		case client := <-h.register:
			h.registerClient(client)

		case client := <-h.unregister:
			h.unregisterClient(client)

		case message := <-h.broadcast:
			h.broadcastMessage(message.encode())
		}
	}
}

func (h *Room) broadcastMessage(message []byte) {
	fmt.Println("----- room -- broadcastMessage")
	for client := range h.chatter.RoomClients(h).List() {
		select {
		case client.send <- message:
		default:
			// ???
			h.closeClient(client)
		}
	}
}

func (h *Room) closeClient(client *Client) {
	fmt.Println("----- room -- closeClient")
	h.chatter.RoomClients(h).Delete(client)
	close(client.send)
}

func (h *Room) unregisterClient(client *Client) {
	fmt.Println("----- room -- unregisterClient")

	find := h.chatter.RoomClients(h).Find(*client)
	if find != nil {
		h.closeClient(client)
	}
}

func (h *Room) registerClient(client *Client) {
	fmt.Println("----- room -- registerClient")
	h.chatter.RoomClients(h).Add(client)
}

func (h *Room) Name() string {
	return h.info.Name
}

func (h *Room) Info() *RoomID {
	return h.info
}
