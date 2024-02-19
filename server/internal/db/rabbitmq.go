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
		constans.AmqpStr,
		constans.RbtUSER, constans.RbtPASS, constans.RabHOST, constans.RbtPORT,
	)
	RabbitConn, err := amqp091.Dial(dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer RabbitConn.Close()
	RabbitChan, err := RabbitConn.Channel()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer RabbitChan.Close()
	log.Info(constans.RabbitMQConnected)
	return RabbitConn, RabbitChan
}
