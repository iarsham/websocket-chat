package routers

import (
	"github.com/gorilla/mux"
	"github.com/iarsham/websocket-chat/internal/controllers"
	"github.com/iarsham/websocket-chat/internal/middleware"
	"github.com/iarsham/websocket-chat/internal/services"
	"github.com/iarsham/websocket-chat/pkg/constans"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

func wsGroup(r *mux.Router, log *zap.Logger, chnl *amqp091.Channel, m *middleware.Middleware) {
	pool := services.NewPoolService(log)
	go pool.Start()
	wsRepo := services.NewWsService(pool, chnl, log)
	c := &controllers.WsController{
		Service: wsRepo,
	}
	wsAPI := r.PathPrefix(constans.EmptyStr).Subrouter()
	wsAPI.Use(m.Authenticate)
	wsAPI.HandleFunc("/ws", c.WsHandler)
}
