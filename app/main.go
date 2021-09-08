package main

import (
	config "delay-message-plugins/config"
	"fmt"
	"log"
	"os"

	_rabbitRepo "delay-message-plugins/rabbit/repository"
	_rabbitServ "delay-message-plugins/rabbit/service"
)

var (
	dir        = os.Getenv("ENV")
	cfg        = config.LoadConfig(dir)
	conn       = config.ConnectAmqp(cfg)
	repoBroker = _rabbitRepo.NewBrokerRepository(*conn)
	servBroker = _rabbitServ.NewBrokerService(repoBroker)
)

func main() {
	broker, err := servBroker.CreateQueue()
	if err != nil {
		log.Fatalln(err)
		return
	}

	repoPublish := _rabbitRepo.NewPublishRepository(broker.Ch)
	servPublish := _rabbitServ.NewPublishService(repoPublish)

	err = servPublish.Send(broker.ExchangeName, broker.KeyRoute, "coba ya 2 heheh")
	if err != nil {
		log.Fatalln(err)
		return
	}

	repoConsume := _rabbitRepo.NewConsumeRepository(broker.Ch)
	servConsume := _rabbitServ.NewConsumeService(repoConsume)

	msgs, err := servConsume.Receive(broker.QueueName)
	if err != nil {
		log.Fatalln(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Println("Receive Data : " + string(d.Body))
		}
	}()

	<-forever
}
