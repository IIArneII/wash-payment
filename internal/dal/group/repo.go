package group

import (
	"wash-payment/internal/app"

	"github.com/gocraft/dbr/v2"
	"go.uber.org/zap"
)

type groupRepo struct {
	l  *zap.SugaredLogger
	db *dbr.Connection
}

func NewRepo(l *zap.SugaredLogger, db *dbr.Connection) app.GroupRepo {
	return &groupRepo{
		l:  l,
		db: db,
	}
}
