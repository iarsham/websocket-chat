package db

import (
	"context"
	"fmt"
	"github.com/iarsham/websocket-chat/pkg/constans"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var RDS *redis.Client

func ConnRedis(log *zap.Logger) *redis.Client {
	RDS := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf(constans.RedisStr, constans.RdsHOST, constans.RdsPORT),
		Password: constans.RdsPassword,
		DB:       0,
	})
	defer RDS.Close()
	_, err := RDS.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	log.Info(constans.RedisConnected)
	return RDS
}
