package transaction

import (
	"wash-payment/internal/app"

	"github.com/gocraft/dbr/v2"
	"go.uber.org/zap"
)

type transactionRepo struct {
	l  *zap.SugaredLogger
	db *dbr.Connection
}

func NewRepo(l *zap.SugaredLogger, db *dbr.Connection) app.TransactionRepo {
	return &transactionRepo{
		l:  l,
		db: db,
	}
}
