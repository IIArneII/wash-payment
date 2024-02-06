package transaction

import (
	"context"
	"wash-payment/internal/app"
	"wash-payment/internal/app/entity"
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

	txs, err := s.transactionRepo.List(ctx, filter)
	if err != nil {
		return entity.Page[entity.Transaction]{}, err
	}

	return txs, nil
}
