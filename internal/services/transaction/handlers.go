package transaction

import (
	"context"
	"time"
	"wash-payment/internal/app"
	"wash-payment/internal/app/entity"

	uuid "github.com/satori/go.uuid"
)

func (s *transactionService) List(ctx context.Context, auth entity.Auth, filter entity.TransactionFilter) (entity.Page[entity.Transaction], error) {
	if auth.User.Role != entity.SystemManagerRole {
		if auth.User.Role == entity.NoAccessRole {
			return entity.Page[entity.Transaction]{}, app.ErrForbidden
		}

		if auth.User.OrganizationID == nil {
			return entity.Page[entity.Transaction]{}, app.ErrForbidden
		}

		if auth.User.OrganizationID != &filter.OrganizationID {
			return entity.Page[entity.Transaction]{}, app.ErrForbidden
		}
	}

	org, err := s.organizationRepo.Get(ctx, filter.OrganizationID)
	if err != nil {
		return entity.Page[entity.Transaction]{}, err
	}
	if org.Deleted {
		return entity.Page[entity.Transaction]{}, app.ErrNotFound
	}

	txs, err := s.transactionRepo.List(ctx, filter)
	if err != nil {
		return entity.Page[entity.Transaction]{}, err
	}

	return txs, nil
}

func (s *transactionService) Deposit(ctx context.Context, auth entity.Auth, organizationID uuid.UUID, amount int64) error {
	if auth.User.Role != entity.SystemManagerRole {
		return app.ErrForbidden
	}

	if amount <= 0 {
		return app.ErrBadValue
	}

	org, err := s.organizationRepo.Get(ctx, organizationID)
	if err != nil {
		return err
	}
	if org.Deleted {
		return app.ErrNotFound
	}

	_, err = s.transactionRepo.Create(ctx, entity.TransactionCreate{
		ID:             uuid.NewV4(),
		OrganizationID: organizationID,
		UserID:         &auth.User.ID,
		Amount:         amount,
		Operation:      entity.DepositOperation,
		CreatedAt:      time.Now().UTC(),
		Service:        entity.PaymentService,
	})
	return err
}

func (s *transactionService) Withdrawal(ctx context.Context, withdrawal entity.Withdrawal) error {
	if withdrawal.StationsCount <= 0 {
		return app.ErrBadValue
	}

	washServerDB, err := s.washserverRepo.Get(ctx, withdrawal.WashServerID)
	if err != nil {
		return err
	}
	if washServerDB.Deleted {
		return app.ErrNotFound
	}

	groupDB, err := s.groupRepo.Get(ctx, washServerDB.GroupID)
	if err != nil {
		return err
	}
	if groupDB.Deleted {
		return app.ErrNotFound
	}

	organizationDB, err := s.organizationRepo.Get(ctx, groupDB.OrganizationID)
	if err != nil {
		return err
	}
	if organizationDB.Deleted {
		return app.ErrNotFound
	}

	amount := int64(withdrawal.StationsCount) * getPrice(organizationDB.ServicePrices, withdrawal.Service)
	if amount > organizationDB.Balance {
		return app.ErrInsufficientFunds
	}

	forDate := withdrawal.ForDate.Truncate(24 * time.Hour)

	_, err = s.transactionRepo.Create(ctx, entity.TransactionCreate{
		ID:             uuid.NewV4(),
		OrganizationID: organizationDB.ID,
		GroupID:        &groupDB.ID,
		Amount:         amount,
		Operation:      entity.DebitOperation,
		CreatedAt:      time.Now().UTC(),
		Service:        withdrawal.Service,
		ForDate:        &forDate,
		StationsCount:  &withdrawal.StationsCount,
		WashServerID:   &washServerDB.ID,
	})
	if err != nil {
		return err
	}

	return nil
}

func getPrice(prices entity.ServicePrices, service entity.Service) int64 {
	switch service {
	case entity.BonusService:
		return prices.Bonus
	case entity.SbpService:
		return prices.Sbp
	default:
		panic("Unknown service: " + service)
	}
}
