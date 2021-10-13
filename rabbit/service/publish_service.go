package rabbit

import "delay-message-plugins/domain"

type publishServ struct {
	repo domain.PublishRepository
}

func NewPublishService(repo domain.PublishRepository) domain.PublishService {
	return &publishServ{
		repo: repo,
	}
}

func (obj *publishServ) Send(exchangeName string, key string, message string, delay int) (err error) {
	publishData := domain.Publish{
		ExchangeName: exchangeName,
		Key:          key,
	}
	return obj.repo.Send(publishData, message, delay)
}
