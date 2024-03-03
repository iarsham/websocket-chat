package services

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/iarsham/websocket-chat/pkg/constans"
	"github.com/iarsham/websocket-chat/pkg/responses"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"time"
)

type BrokerService struct {
	Receiver  amqp.Queue
	Publisher amqp.Queue
	Channel   *amqp.Channel
	log       *zap.Logger
}

func NewBrokerService(ch *amqp.Channel, log *zap.Logger) *BrokerService {
	q1, err := ch.QueueDeclare(constans.ReceiveQueueStr, false, false, false, false, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	q2, err := ch.QueueDeclare(constans.PublishQueueStr, false, false, false, false, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &BrokerService{
		Receiver:  q1,
		Publisher: q2,
		Channel:   ch,
		log:       log,
	}
}

func (b *BrokerService) Publish(requestBody chan []byte) {
	for body := range requestBody {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		msg := amqp.Publishing{
			Body:        body,
			ContentType: constans.TextContentType,
		}
		err := b.Channel.PublishWithContext(ctx, "", b.Publisher.Name, false, false, msg)
		cancel()
		if err != nil {
			b.log.Fatal(err.Error())
		}
		b.log.Info(constans.MessageSent, zap.String(constans.Message, string(body)))
	}
}

func (b *BrokerService) Read(pool *PoolService) {
	msg, err := b.Channel.Consume(b.Receiver.Name, "", true, false, false, false, nil)
	if err != nil {
		b.log.Fatal(err.Error())
	}
	qrResponse := make(chan responses.QueueResponse)
	go b.transformer(msg, qrResponse)
	go b.processResponse(qrResponse, pool)
}

func (b *BrokerService) transformer(entries <-chan amqp.Delivery, msgs chan responses.QueueResponse) {
	var qr responses.QueueResponse
	for d := range entries {
		err := json.Unmarshal(d.Body, &qr)
		if err != nil {
			b.log.Fatal(err.Error())
		}
		msgs <- qr
	}
}

func (b *BrokerService) processResponse(q <-chan responses.QueueResponse, pool *PoolService) {
	for r := range q {
		qr := responses.QueueResponse{
			RoomID:  r.RoomID,
			Message: r.Message,
		}
		msg := Message{
			Type: websocket.TextMessage,
			Body: Body{
				RoomID:  int(qr.RoomID),
				User:    "user",
				Message: qr.Message,
			},
		}
		pool.Broadcast <- msg
	}
}
