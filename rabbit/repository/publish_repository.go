package rabbit

import (
	"delay-message-plugins/domain"
	"fmt"

	"github.com/streadway/amqp"
)

type publishRepo struct {
	ch *amqp.Channel
}

func NewPublishRepository(ch *amqp.Channel) domain.PublishRepository {
	return &publishRepo{
		ch: ch,
	}
}

func (obj *publishRepo) Send(data domain.Publish, message string) (err error) {

	table := amqp.Table{
		"x-delay": 30000,
	}
	err = obj.ch.Publish(
		data.ExchangeName,
		data.Key,
		false,
		false,
		amqp.Publishing{
			Headers:      table,
			DeliveryMode: amqp.Persistent,
			ContentType:  "text",
			Body:         []byte(message),
		},
	)
	if err != nil {
		err = fmt.Errorf("Failed To Publish : " + err.Error())
	}

	return
}
