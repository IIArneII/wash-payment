package transaction

import (
	"wash-payment/internal/app"

	"go.uber.org/zap"
)

type transactionService struct {
	l                *zap.SugaredLogger
	organizationRepo app.OrganizationRepo
	groupRepo        app.GroupRepo
	washserverRepo   app.WashServerRepo
	transactionRepo  app.TransactionRepo
}

func NewService(l *zap.SugaredLogger, organizationRepo app.OrganizationRepo, transactionRepo app.TransactionRepo, groupRepo app.GroupRepo) app.TransactionService {
	return &transactionService{
		l:                l,
		transactionRepo:  transactionRepo,
		organizationRepo: organizationRepo,
		groupRepo:        groupRepo,
	}
}
