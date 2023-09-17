package user

import (
	"wash-payment/internal/app"

	"go.uber.org/zap"
)

type userService struct {
	l        *zap.SugaredLogger
	userRepo app.UserRepo
}

func NewUserService(l *zap.SugaredLogger, userRepo app.UserRepo) app.UserService {
	return &userService{
		l:        l,
		userRepo: userRepo,
	}
}
