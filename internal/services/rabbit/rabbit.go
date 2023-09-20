package rabbit

import (
	"wash-payment/internal/app"

	"go.uber.org/zap"
)

type rabbitService struct {
	l *zap.SugaredLogger
}

func NewRabbitService(l *zap.SugaredLogger) app.RabbitService {
	return &rabbitService{
		l: l,
	}
}
