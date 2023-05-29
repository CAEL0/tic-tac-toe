package hub

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CAEL0/tic-tac-toe/server/board"
	"github.com/gorilla/websocket"
)

type Client struct {
	hub   *Hub
	conn  *websocket.Conn
	send  chan []byte
	board *board.Board
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ServeWebsocket(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Fail to upgrade: %v", err)
	}
	client := &Client{
		hub:   hub,
		conn:  conn,
		send:  make(chan []byte, 256),
		board: board.New(),
	}
	hub.register <- client

	go client.readPump()
	go client.writePump()
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("Fail to read message: %v", err)
			break
		}
		c.hub.broadcast <- message
	}
}

func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			fmt.Println(message)
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		}

	}
}
