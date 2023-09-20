package services

import (
	"wash-payment/internal/app"
	"wash-payment/internal/services/rabbit"
	"wash-payment/internal/services/user"

	"go.uber.org/zap"
)

func NewServices(l *zap.SugaredLogger, dal *app.Dal) *app.Services {
	return &app.Services{
		UserService:   user.NewUserService(l, dal.UserRepo),
		RabbitService: rabbit.NewRabbitService(l),
	}
}
