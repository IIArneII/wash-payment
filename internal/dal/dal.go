package dal

import (
	"wash-payment/internal/app"
	"wash-payment/internal/dal/user"

	"github.com/gocraft/dbr/v2"
	"go.uber.org/zap"
)

func NewDal(l *zap.SugaredLogger, db *dbr.Connection) *app.Dal {
	return &app.Dal{
		UserRepo: user.NewRepo(l, db),
	}
}
