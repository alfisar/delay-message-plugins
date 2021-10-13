package email

import (
	"delay-message-plugins/domain"
	"strconv"
	"time"

	"gopkg.in/gomail.v2"
)

type repository struct {
	email domain.Email
}

func NewEmailRepository(email domain.Email) domain.EmailRepo {
	return &repository{
		email: email,
	}
}

func (repo *repository) EmailContent(timeNow time.Time) (content string) {
	content = "<b> Hai Sobat </b>, ini adalah hasil test scheduler email menggunakan rabbitmq, ini menggunakan delay 5, 7, dan 10 menit"
	return
}
func (repo *repository) SendEmail(expiredAt time.Time) (err error) {
	timeNow := expiredAt
	bodyEmail := repo.EmailContent(timeNow)
	port, _ := strconv.Atoi(repo.email.Port)
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", repo.email.Email)
	mailer.SetHeader("To", "alfisar417@gmail.com")
	mailer.SetHeader("subject", "Email Test Scheduler Rabbitmq")
	mailer.SetBody("text/html", bodyEmail)

	dial := gomail.NewDialer(
		repo.email.Host,
		port,
		repo.email.Email,
		repo.email.Password,
	)
	return dial.DialAndSend(mailer)
}
