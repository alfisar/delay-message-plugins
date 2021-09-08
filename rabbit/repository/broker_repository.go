package rabbit

import (
	"delay-message-plugins/domain"
	"fmt"

	"github.com/streadway/amqp"
)

type repository struct {
	conn amqp.Connection
}

func NewBrokerRepository(conn amqp.Connection) domain.BrokerRepository {
	return &repository{
		conn: conn,
	}
}

func (obj *repository) CreateQueue(data domain.Broker) (result domain.Broker, err error) {
	data.Ch, err = obj.conn.Channel()
	if err != nil {
		err = fmt.Errorf("Failed to open channel : " + err.Error())
		return
	}

	table := amqp.Table{
		"x-delayed-type": "direct",
	}
	err = data.Ch.ExchangeDeclare(
		data.ExchangeName,
		data.Kind,
		true,
		false,
		false,
		false,
		table,
	)
	if err != nil {
		err = fmt.Errorf("Failed Declare Exchange : " + err.Error())
		return
	}

	_, err = data.Ch.QueueDeclare(
		data.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		err = fmt.Errorf("Failed Declare Queue : " + err.Error())
		return
	}

	err = data.Ch.QueueBind(
		data.QueueName,
		data.KeyRoute,
		data.ExchangeName,
		false,
		nil,
	)

	if err != nil {
		err = fmt.Errorf("Failed Bind Queue : " + err.Error())
		return
	}

	result = data
	return
}
