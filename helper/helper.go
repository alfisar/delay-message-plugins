package helper

import (
	"delay-message-plugins/config"
	"delay-message-plugins/domain"
)

func CofigEmail(cfg config.Config) domain.Email {
	return domain.Email{
		Host:     cfg.EMAIL_Host,
		Port:     cfg.EMAIL_Port,
		Email:    cfg.EMAIL_Email,
		Password: cfg.EMAIL_Password,
	}
}
