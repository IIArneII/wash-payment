package organization

import (
	"wash-payment/internal/app"

	"github.com/gocraft/dbr/v2"
	"go.uber.org/zap"
)

type organizationRepo struct {
	l  *zap.SugaredLogger
	db *dbr.Connection
}

func NewRepo(l *zap.SugaredLogger, db *dbr.Connection) app.OrganizationRepo {
	return &organizationRepo{
		l:  l,
		db: db,
	}
}
