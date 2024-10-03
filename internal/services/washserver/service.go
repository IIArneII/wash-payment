package washserver

import (
	"wash-payment/internal/app"

	"go.uber.org/zap"
)

type washServerService struct {
	l              *zap.SugaredLogger
	washServerRepo app.WashServerRepo
}

func NewService(l *zap.SugaredLogger, washServerRepo app.WashServerRepo) app.WashServerService {
	return &washServerService{
		l:              l,
		washServerRepo: washServerRepo,
	}
}
