package services

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/iarsham/websocket-chat/pkg/constans"
	"go.uber.org/zap"
)

type Client struct {
	ID       string
	UserID   string
	Username string
	Pool     *PoolService
	Conn     *websocket.Conn
	log      *zap.Logger
}

type Body struct {
	RoomID   int    `json:"room_id" validate:"omitempty"`
	RoomName string `json:"room_name" validate:"omitempty"`
	Message  string `json:"message" validate:"omitempty"`
	User     string `json:"user" validate:"omitempty"`
}

type Message struct {
	Type int  `json:"type"`
	Body Body `json:"body"`
}

func (c *Client) Read(body chan []byte) {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()
	for {
		var body Body
		msgType, msg, err := c.Conn.ReadMessage()
		if err != nil {
			c.log.Fatal(err.Error())
		}
		err = json.Unmarshal(msg, &body)
		if err != nil {
			c.log.Fatal(err.Error())
		}
		body.User = c.Username
		message := Message{Type: msgType, Body: body}
		c.Pool.Broadcast <- message
		c.log.Info(constans.MessageReceived, zap.Any(constans.MessageType, msgType))
	}
}
