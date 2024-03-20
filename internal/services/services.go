package services

import (
	"wash-payment/internal/app"
	"wash-payment/internal/services/group"
	"wash-payment/internal/services/organization"
	"wash-payment/internal/services/transaction"
	"wash-payment/internal/services/user"

	"go.uber.org/zap"
)

func NewServices(l *zap.SugaredLogger, dal *app.Repositories) *app.Services {
	return &app.Services{
		UserService:         user.NewService(l, dal.UserRepo),
		OrganizationService: organization.NewService(l, dal.OrganizationRepo, dal.TransactionRepo, dal.ServicePriceRepo),
		GroupService:        group.NewService(l, dal.GroupRepo),
		TransactionService:  transaction.NewService(l, dal.OrganizationRepo, dal.TransactionRepo, dal.GroupRepo),
	}
}
