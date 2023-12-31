package dal

import (
	"wash-payment/internal/app"
	"wash-payment/internal/dal/group"
	"wash-payment/internal/dal/organization"
	"wash-payment/internal/dal/transaction"
	"wash-payment/internal/dal/user"

	"github.com/gocraft/dbr/v2"
	"go.uber.org/zap"
)

func NewRepositories(l *zap.SugaredLogger, db *dbr.Connection) *app.Repositories {
	return &app.Repositories{
		UserRepo:         user.NewRepo(l, db),
		OrganizationRepo: organization.NewRepo(l, db),
		GroupRepo:        group.NewRepo(l, db),
		TransactionRepo:  transaction.NewRepo(l, db),
	}
}
