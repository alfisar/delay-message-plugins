package domain

import "github.com/streadway/amqp"

type Broker struct {
	ExchangeName string
	Kind         string
	KeyRoute     string
	QueueName    string
	Ch           *amqp.Channel
}

type BrokerRepository interface {
	CreateQueue(data Broker) (result Broker, err error)
}

type BrokerService interface {
	CreateQueue() (result Broker, err error)
}

type Publish struct {
	ExchangeName string
	Key          string
}

type PublishRepository interface {
	Send(data Publish, message string, delay int) (err error)
}

type PublishService interface {
	Send(ExchangeName string, Key string, message string, delay int) (err error)
}

type Consume struct {
	QueueName string
}

type ConsumeRepository interface {
	Receive(data Consume) (deliv <-chan amqp.Delivery, err error)
}

type ConsumeService interface {
	Receive(queueName string) (deliv <-chan amqp.Delivery, err error)
}
