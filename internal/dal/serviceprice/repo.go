package serviceprice

import (
	"wash-payment/internal/app"

	"github.com/gocraft/dbr/v2"
	"go.uber.org/zap"
)

type servicePriceRepo struct {
	l  *zap.SugaredLogger
	db *dbr.Connection
}

func NewRepo(l *zap.SugaredLogger, db *dbr.Connection) app.ServicePriceRepo {
	return &servicePriceRepo{
		l:  l,
		db: db,
	}
}
