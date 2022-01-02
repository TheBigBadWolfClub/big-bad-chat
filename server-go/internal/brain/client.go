package brain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	info *UserId

	connection SocketConnection

	// Buffered channel of outbound messages.
	send chan []byte

	//
	chatter Chatter
}

func NewClient(name string, chatter Chatter) *Client {
	fmt.Println("client --- NewClient ")

	client := &Client{
		info: &UserId{
			ID: ID{
				Name: name,
				UUID: uuid.NewString(),
			},
			LastConnection: time.Now(),
		},
		connection: nil,
		send:       make(chan []byte, 256),
		chatter:    chatter,
	}

	return client

}

func (c *Client) disconnect() {
	fmt.Println("DISCONNET")
	for room := range c.chatter.ClientRooms(c).List() {
		room.unregister <- c
	}
	_ = c.connection.Close()
}

func (c *Client) newMessage(jsonMessage []byte) {
	fmt.Println("client --- newMessage ")
	var message Message
	if err := json.Unmarshal(jsonMessage, &message); err != nil {
		log.Printf("Error on unmarshal JSON message %s", err)
	}

	fmt.Println(message)
	message.Sender = c
	switch message.Action {
	case NewTextAction:
		c.newTextAction(&message)

	case JoinRoomAction:
		c.joinRoomAction(&message)

	case LeaveRoomAction:
		c.leaveRoomAction(&message)
	}
}

func (c *Client) newTextAction(msg *Message) {
	room := c.chatter.ClientRooms(c).Find(msg.ChatMessage.RoomUuid)
	if room != nil {
		room.messages.Append(msg.ChatMessage)
		room.broadcast <- msg
	}
}

func (c *Client) joinRoomAction(msg *Message) {
	room := c.chatter.AllRooms().Find(msg.RoomId.UUID)
	if room != nil {
		c.chatter.ClientRooms(c).Add(room)
		room.register <- c
	}
}

func (c *Client) leaveRoomAction(msg *Message) {
	fmt.Println("client --- leaveRoomAction ")

	room := c.chatter.ClientRooms(c).Find(msg.RoomId.UUID)
	if room != nil {
		c.chatter.ClientRooms(c).Delete(room)
		room.unregister <- c
	}
}

// ReadPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) ReadPump() {

	// Start endless read loop, waiting for messages from client
	c.connection.SetupRead()
	for {
		message, err := c.connection.ReadMessage()
		if err != nil {
			c.disconnect()
			break
		}

		message = bytes.TrimSpace(bytes.ReplaceAll(message, newline, space))
		c.newMessage(message)
	}

	fmt.Println("ReadPump exit: ", c.info.Name+", ID: ", c.connection.Uuid())
}

func (c *Client) WritePump() {
	c.connection.WritePump(c.info.Name, c.send)
}

func (c *Client) Info() *UserId {
	return c.info
}

func (c *Client) ConnectionUuid() string {
	if c.connection == nil {
		return ""
	}
	return c.connection.Uuid()
}

func (c *Client) AddConnection(conn *websocket.Conn, uuid string) {
	c.connection = NewSocketConnection(conn, uuid)
	c.connection.SetCloseHandler(func(code int, text string) error {
		c.info.LastConnection = time.Now()
		return nil
	})
}

func (c *Client) IsConnected() bool {
	return c.connection != nil
}
