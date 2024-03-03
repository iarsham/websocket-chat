package services

import (
	"github.com/gorilla/websocket"
	"github.com/iarsham/websocket-chat/pkg/constans"
	"go.uber.org/zap"
)

type PoolService struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Message
	log        *zap.Logger
}

func NewPoolService(log *zap.Logger) *PoolService {
	return &PoolService{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan Message),
		log:        log,
	}
}

func (p *PoolService) Start() error {
	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = true
			p.log.Info(constans.NewUserJoined, zap.Int(constans.Members, len(p.Clients)))
			for c := range p.Clients {
				err := c.Conn.WriteJSON(Message{Type: websocket.TextMessage, Body: Body{Message: constans.NewUserJoined}})
				if err != nil {
					p.log.Fatal(err.Error())
				}
			}

		case client := <-p.Unregister:
			delete(p.Clients, client)
			p.log.Info(constans.UserDisconnected, zap.Int(constans.Members, len(p.Clients)))
			for c := range p.Clients {
				err := c.Conn.WriteJSON(Message{Type: websocket.TextMessage, Body: Body{Message: constans.UserDisconnected}})
				if err != nil {
					p.log.Fatal(err.Error())
				}
			}

		case msg := <-p.Broadcast:
			p.log.Info(constans.MsgToClients)
			for c := range p.Clients {
				err := c.Conn.WriteJSON(msg)
				if err != nil {
					p.log.Fatal(err.Error())
				}
			}
		}

	}
}
