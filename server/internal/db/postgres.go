package db

import (
	"database/sql"
	"fmt"
	"github.com/iarsham/websocket-chat/pkg/constans"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var DB *sql.DB

func ConnDB(log *zap.Logger) *sql.DB {
	var err error
	dsn := fmt.Sprintf(
		constans.PgStr,
		constans.PgHost, constans.PgUSER, constans.PgPASSWORD, constans.PgName, constans.PgPORT,
	)
	DB, err = sql.Open(constans.DbName, dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	if err = DB.Ping(); err != nil {
		log.Fatal(err.Error())
	}
	log.Info(constans.PostgresConnected)
	return DB
}

func CloseDB() {
	defer DB.Close()
}
