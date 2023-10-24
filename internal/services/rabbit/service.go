package rabbit

import (
	"wash-payment/internal/app"

	"go.uber.org/zap"
)

type rabbitService struct {
	l        *zap.SugaredLogger
	services *app.Services
}

func NewService(l *zap.SugaredLogger, services *app.Services) app.RabbitService {
	return &rabbitService{
		l:        l,
		services: services,
	}
}
