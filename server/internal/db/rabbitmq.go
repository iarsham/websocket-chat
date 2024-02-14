package db

import (
	"fmt"
	"github.com/iarsham/websocket-chat/pkg/constans"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

var (
	RabbitConn *amqp091.Connection
	RabbitChan *amqp091.Channel
)

func ConnRabbit(log *zap.Logger) (*amqp091.Connection, *amqp091.Channel) {
	dsn := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		constans.RbtUSER, constans.RbtPASS, constans.RabHOST, constans.RbtPORT,
	)
	RabbitConn, err := amqp091.Dial(dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	RabbitChan, err := RabbitConn.Channel()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Info(constans.RabbitMQConnected)
	return RabbitConn, RabbitChan
}

func DisConnRabbit(log *zap.Logger) {
	if err := RabbitConn.Close(); err != nil{
		log.Fatal(err.Error())
	}
	if err := RabbitChan.Close(); err != nil{
		log.Fatal(err.Error())
	}
	log.Info(constans.RabbitMQClosed)
}
