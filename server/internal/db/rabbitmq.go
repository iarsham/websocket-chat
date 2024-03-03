package db

import (
	"fmt"
	"github.com/iarsham/websocket-chat/pkg/constans"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

var (
	Conn *amqp091.Connection
	Chan *amqp091.Channel
)

func ConnRabbit(log *zap.Logger) (*amqp091.Connection, *amqp091.Channel) {
	dsn := fmt.Sprintf(
		constans.AmqpStr,
		constans.RbtUSER, constans.RbtPASS, constans.RabHOST, constans.RbtPORT,
	)
	Conn, err := amqp091.Dial(dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	Chan, err := Conn.Channel()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Info(constans.RabbitMQConnected)
	return Conn, Chan
}

func CloseRabbit() {
	defer func() {
		Conn.Close()
		Chan.Close()
	}()
}
