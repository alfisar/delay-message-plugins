package config

import (
	"log"

	"github.com/streadway/amqp"
)

func ConnectAmqp(cfg *Config) *amqp.Connection {
	uri := cfg.AMQP_Driver + "://" + cfg.AMQP_User + ":" + cfg.AMQP_Pass + "@" + cfg.AMQP_Host + ":" + cfg.AMQP_Port + "/"

	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Panic("Cannot connect to RabbitMQ server : " + err.Error())
	} else {
		log.Println("Successfuly connect to RabbitMQ Server")
	}

	return conn
}
