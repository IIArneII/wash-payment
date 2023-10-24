package group

import (
	"wash-payment/internal/app"

	"go.uber.org/zap"
)

type groupService struct {
	l         *zap.SugaredLogger
	groupRepo app.GroupRepo
}

func NewService(l *zap.SugaredLogger, groupRepo app.GroupRepo) app.GroupService {
	return &groupService{
		l:         l,
		groupRepo: groupRepo,
	}
}
