package rabbit

import (
	"delay-message-plugins/domain"

	"github.com/streadway/amqp"
)

type consumeServ struct {
	repo domain.ConsumeRepository
}

func NewConsumeService(repo domain.ConsumeRepository) domain.ConsumeService {
	return &consumeServ{
		repo: repo,
	}
}

func (obj *consumeServ) Receive(queueName string) (deliv <-chan amqp.Delivery, err error) {
	data := domain.Consume{
		QueueName: queueName,
	}
	return obj.repo.Receive(data)
}
