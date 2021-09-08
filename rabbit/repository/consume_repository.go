package rabbit

import (
	"delay-message-plugins/domain"
	"fmt"

	"github.com/streadway/amqp"
)

type consumeRepo struct {
	ch *amqp.Channel
}

func NewConsumeRepository(ch *amqp.Channel) domain.ConsumeRepository {
	return &consumeRepo{
		ch: ch,
	}
}

func (obj *consumeRepo) Receive(data domain.Consume) (deliv <-chan amqp.Delivery, err error) {
	deliv, err = obj.ch.Consume(
		data.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		err = fmt.Errorf("Failed to get data : " + err.Error())
	}
	return
}
