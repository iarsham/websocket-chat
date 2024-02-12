package main

import (
	"github.com/iarsham/websocket-chat/internal/common"
	"github.com/iarsham/websocket-chat/internal/db"
)

//	@title			Websocket-Chat-API
//	@version		0.1.0
//	@description	This is chat API server.
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	Arsham Roshannejad
//	@contact.url	https://www.linkedin.com/in/arsham-roshannejad
//	@contact.email	arshamdev2001@gmail.com
//	@license.name	MIT
//	@license.url	https://www.mit.edu/~amini/LICENSE.md
//	@host			localhost:8000
//	@BasePath		/api
//	@schemes		http https
func main() {
	logger := common.ZapLogger()
	defer logger.Sync()
	dbInstance := db.InitDB(logger)
	defer db.CloseDB(logger)
	rdsInstance := db.ConnRedis(logger)
	defer db.DisConnRedis(logger)
	RunServer(dbInstance, rdsInstance, logger)
}
