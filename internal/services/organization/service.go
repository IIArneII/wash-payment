package organization

import (
	"wash-payment/internal/app"

	"go.uber.org/zap"
)

type organizationService struct {
	l                *zap.SugaredLogger
	organizationRepo app.OrganizationRepo
	transactionRepo  app.TransactionRepo
}

func NewService(l *zap.SugaredLogger, organizationRepo app.OrganizationRepo, transactionRepo app.TransactionRepo) app.OrganizationService {
	return &organizationService{
		l:                l,
		transactionRepo:  transactionRepo,
		organizationRepo: organizationRepo,
	}
}
