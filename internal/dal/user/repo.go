package user

import (
	"wash-payment/internal/app"

	"github.com/gocraft/dbr/v2"
	"go.uber.org/zap"
)

type userRepo struct {
	l  *zap.SugaredLogger
	db *dbr.Connection
}

func NewRepo(l *zap.SugaredLogger, db *dbr.Connection) app.UserRepo {
	return &userRepo{
		l:  l,
		db: db,
	}
}
