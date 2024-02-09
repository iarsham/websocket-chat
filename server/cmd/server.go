package main

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"github.com/google/uuid"
	"github.com/iarsham/websocket-chat/internal/routers"
	"github.com/iarsham/websocket-chat/pkg/constans"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func RunServer(db *sql.DB, rds *redis.Client, log *zap.Logger) {
	mux := routers.SetupRoutes(db, rds, log)
	log.Info(fmt.Sprintf("Server started on port %s...", constans.SrvPort))
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%s", constans.SrvPort),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func init() {
	gob.Register(uuid.UUID{})
}
