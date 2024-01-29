package db

import (
	"database/sql"
	"fmt"
	"github.com/iarsham/websocket-chat/pkg/constans"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"os"
)

var (
	HOST     = os.Getenv("PG_HOST")
	USER     = os.Getenv("PG_USER")
	PASSWORD = os.Getenv("PG_PASS")
	Name     = os.Getenv("PG_DB")
	PORT     = os.Getenv("PG_PORT")
	DB       *sql.DB
)

func InitDB(log *zap.Logger) *sql.DB {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		HOST, USER, PASSWORD, Name, PORT,
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

func CloseDB(log *zap.Logger) {
	err := DB.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Info(constans.PostgresClosed)
}
