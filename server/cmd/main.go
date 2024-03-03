package main

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"github.com/google/uuid"
	"github.com/iarsham/websocket-chat/internal/common"
	"github.com/iarsham/websocket-chat/internal/db"
	"github.com/iarsham/websocket-chat/internal/routers"
	"github.com/iarsham/websocket-chat/pkg/constans"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// @title			Websocket-Chat-API
// @version		0.1.0
// @description	This is chat API server.
// @termsOfService	http://swagger.io/terms/
// @contact.name	Arsham Roshannejad
// @contact.url	https://www.linkedin.com/in/arsham-roshannejad
// @contact.email	arshamdev2001@gmail.com
// @license.name	MIT
// @license.url	https://www.mit.edu/~amini/LICENSE.md
// @host			localhost:8000
// @BasePath		/api
// @schemes		http https
func main() {
	log := common.ZapLogger()
	defer log.Sync()
	pg := db.ConnDB(log)
	defer db.CloseDB()
	_, chnl := db.ConnRabbit(log)
	defer db.CloseRabbit()
	RunServer(pg, chnl, log)
}

func RunServer(db *sql.DB, chnl *amqp091.Channel, log *zap.Logger) {
	mux := routers.SetupRoutes(db, chnl, log)
	log.Info(fmt.Sprintf(constans.StartSrvLog, constans.SrvPort))
	srv := &http.Server{
		Addr:           fmt.Sprintf(constans.SrvStr, constans.SrvPort),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(srv.ListenAndServe().Error())
}

func init() {
	gob.Register(uuid.UUID{})
}
