package rabbit

import (
	"delay-message-plugins/domain"
)

type service struct {
	repo domain.BrokerRepository
}

func NewBrokerService(repo domain.BrokerRepository) domain.BrokerService {
	return &service{
		repo: repo,
	}
}

func (obj *service) CreateQueue() (result domain.Broker, err error) {
	data := domain.Broker{
		ExchangeName: "delay-message-exchange-plugins",
		Kind:         "x-delayed-message",
		KeyRoute:     "delay-message.pasar",
		QueueName:    "delay-message-pasar-plugins",
	}

	return obj.repo.CreateQueue(data)
}
