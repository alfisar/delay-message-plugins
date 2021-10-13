package domain

import "time"

type Email struct {
	Host     string
	Port     string
	Email    string
	Password string
}

type EmailRepo interface {
	SendEmail(expiredAt time.Time) (err error)
}

type EmailService interface {
	SendEmail(expiredAt time.Time) (err error)
}
