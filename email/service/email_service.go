package email

import (
	"delay-message-plugins/domain"
	"time"
)

type service struct {
	repo domain.EmailRepo
}

func NewEmilService(repo domain.EmailRepo) domain.EmailService {
	return &service{
		repo: repo,
	}
}

func (obj *service) SendEmail(expiredAt time.Time) (err error) {
	return obj.repo.SendEmail(expiredAt)
}
