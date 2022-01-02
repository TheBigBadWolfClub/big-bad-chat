package brain

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type SocketConnection interface {
	Close() error
	WritePump(string, chan []byte)
	ReadMessage() ([]byte, error)
	SetupRead()
	Uuid() string
	SetCloseHandler(handler func(code int, text string) error)
}

type connection struct {
	// The websocket connection.
	conn *websocket.Conn
	uuid string
}

func NewSocketConnection(conn *websocket.Conn, uuid string) SocketConnection {
	return &connection{
		conn: conn,
		uuid: uuid,
	}
}

func (c connection) Uuid() string {
	return c.uuid
}

func (c connection) Close() error {
	return c.conn.Close()
}

func (c connection) SetCloseHandler(handler func(code int, text string) error) {
	c.conn.SetCloseHandler(handler)
}

func (c connection) SetupRead() {
	c.conn.SetReadLimit(maxMessageSize)
	_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
}

func (c *connection) ReadMessage() ([]byte, error) {

	_, message, err := c.conn.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Printf("error: %v", err)

		}
		return nil, err
	}

	message = bytes.TrimSpace(bytes.ReplaceAll(message, newline, space))
	return message, nil
}

// WritePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *connection) WritePump(clientName string, send chan []byte) {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		fmt.Println("WritePump exit: ", clientName)
		_ = c.conn.Close()
		// c.conn = nil
	}()

	for {
		select {
		case message, ok := <-send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, _ = w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(send)
			for i := 0; i < n; i++ {
				_, _ = w.Write(newline)
				_, _ = w.Write(<-send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
