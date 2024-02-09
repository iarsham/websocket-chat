package main

import (
	"github.com/iarsham/websocket-chat/internal/common"
	"github.com/iarsham/websocket-chat/internal/db"
)

func main() {
	logger := common.ZapLogger()
	defer logger.Sync()
	dbInstance := db.InitDB(logger)
	defer db.CloseDB(logger)
	rdsInstance := db.ConnRedis(logger)
	defer db.DisConnRedis(logger)
	RunServer(dbInstance, rdsInstance, logger)
}
