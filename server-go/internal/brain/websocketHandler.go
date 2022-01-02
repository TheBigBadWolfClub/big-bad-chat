package brain

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ClientServer interface {
	NewClientConnection(w http.ResponseWriter, r *http.Request)
}

type websocketConn struct {
	// Registered clients.
	chatter Chatter
}

func NewClientServer(chat Chatter) ClientServer {
	return &websocketConn{
		chatter: chat,
	}
}

// NewClientConnection handles websocket requests from the peer.
func (server *websocketConn) NewClientConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	username := r.URL.Query().Get("user")
	uuid := r.URL.Query().Get("uuid")
	fmt.Println("Accepting client: ", username, ", conn uuid: ", uuid)
	client := server.chatter.AddUser(username)
	client.AddConnection(conn, uuid)
	server.chatter.UserHasConnected(client)

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
}
