package main

import (
	config "delay-message-plugins/config"
	"delay-message-plugins/helper"
	"fmt"
	"log"
	"os"
	"time"

	_rabbitRepo "delay-message-plugins/rabbit/repository"
	_rabbitServ "delay-message-plugins/rabbit/service"

	_emailRepo "delay-message-plugins/email/repository"
	_emailServ "delay-message-plugins/email/service"
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

	for i := 0; i < 3; i++ {
		// delay in milisecond
		var delay int
		if i == 0 {
			delay = 300000
		} else if i == 1 {
			delay = 420000
		} else {
			delay = 600000
		}
		err = servPublish.Send(broker.ExchangeName, broker.KeyRoute, "coba ya 2 heheh", delay)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}

	repoConsume := _rabbitRepo.NewConsumeRepository(broker.Ch)
	servConsume := _rabbitServ.NewConsumeService(repoConsume)

	msgs, err := servConsume.Receive(broker.QueueName)
	if err != nil {
		log.Fatalln(err)
	}

	forever := make(chan bool)

	emailData := helper.CofigEmail(*cfg)
	repoEmail := _emailRepo.NewEmailRepository(emailData)
	servEmail := _emailServ.NewEmilService(repoEmail)
	expiredAt := time.Now().Add(time.Minute * 15)
	go func() {
		for d := range msgs {
			fmt.Println(string(d.Body))
			err := servEmail.SendEmail(expiredAt)
			if err != nil {
				fmt.Println(err)
			}

		}
	}()

	<-forever
}
