package services

import (
	"github.com/gorilla/websocket"
	"github.com/iarsham/websocket-chat/pkg/constans"
	"github.com/iarsham/websocket-chat/pkg/utils"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"net/http"
	"slices"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  constans.UpgraderBufferSize,
	WriteBufferSize: constans.UpgraderBufferSize,
	CheckOrigin: func(r *http.Request) bool {
		if constans.Mode {
			return true
		}
		return slices.Contains(
			utils.GetListEnv(constans.ORIGINS),
			r.Header.Get(constans.Origin),
		)
	},
}

type WsService struct {
	pool   *PoolService
	broker *BrokerService
	log    *zap.Logger
}

func NewWsService(pool *PoolService, chnl *amqp091.Channel, log *zap.Logger) *WsService {
	return &WsService{
		pool:   pool,
		broker: NewBrokerService(chnl, log),
		log:    log,
	}
}

func (ws *WsService) ServeWs(w http.ResponseWriter, r *http.Request, userID, userName string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		ws.log.Warn(err.Error())
	}
	client := &Client{
		ID:       userID,
		Username: userName,
		Conn:     conn,
		Pool:     ws.pool,
		log:      ws.log,
	}
	ws.pool.Register <- client
	reqBody := make(chan []byte)
	go client.Read(reqBody)
	go ws.broker.Read(ws.pool)
	go ws.broker.Publish(reqBody)
}
