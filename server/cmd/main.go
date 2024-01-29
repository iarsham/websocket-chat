package main

import (
	"github.com/iarsham/websocket-chat/internal/common"
	"github.com/iarsham/websocket-chat/internal/db"
)

func main() {
	logger := common.ZapLogger()
	defer logger.Sync()

	_ = db.InitDB(logger)
	defer db.CloseDB(logger)
}
