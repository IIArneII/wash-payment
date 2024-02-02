package transaction

import (
	"wash-payment/internal/app"

	"go.uber.org/zap"
)

type transactionService struct {
	l                *zap.SugaredLogger
	organizationRepo app.OrganizationRepo
	transactionRepo  app.TransactionRepo
}

func NewService(l *zap.SugaredLogger, organizationRepo app.OrganizationRepo, transactionRepo app.TransactionRepo) app.TransactionService {
	return &transactionService{
		l:                l,
		transactionRepo:  transactionRepo,
		organizationRepo: organizationRepo,
	}
}
