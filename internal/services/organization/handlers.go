package organization

import (
	"context"
	"errors"
	"time"
	"wash-payment/internal/app"
	"wash-payment/internal/app/conversions"
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"

	uuid "github.com/satori/go.uuid"
)

func (s *organizationService) Get(ctx context.Context, auth app.Auth, organizationID uuid.UUID) (entity.Organization, error) {
	if auth.User.Role != entity.SystemManagerRole {
		if auth.User.Role == entity.NoAccessRole {
			return entity.Organization{}, app.ErrForbidden
		}

		if auth.User.OrganizationID == nil {
			return entity.Organization{}, app.ErrForbidden
		}

		if auth.User.OrganizationID != &organizationID {
			return entity.Organization{}, app.ErrForbidden
		}
	}

	organizationFromDB, err := s.organizationRepo.Get(ctx, organizationID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = app.ErrNotFound
		}

		return entity.Organization{}, err
	}

	return conversions.OrganizationFromDB(organizationFromDB), nil
}

func (s *organizationService) Create(ctx context.Context, organizationCreate entity.OrganizationCreate) (entity.Organization, error) {
	dbOrganization := conversions.OrganizationCreateToDB(organizationCreate)

	newOrganization, err := s.organizationRepo.Create(ctx, dbOrganization)
	if err != nil {
		if errors.Is(err, dbmodels.ErrAlreadyExists) {
			err = app.ErrAlreadyExists
		}

		return entity.Organization{}, err
	}

	return conversions.OrganizationFromDB(newOrganization), nil
}

func (s *organizationService) Update(ctx context.Context, organizationID uuid.UUID, organizationUpdate entity.OrganizationUpdate) error {
	dbOrganizationUpdate := conversions.OrganizationUpdateToDB(organizationUpdate)

	err := s.organizationRepo.Update(ctx, organizationID, dbOrganizationUpdate)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = app.ErrNotFound
		} else if errors.Is(err, dbmodels.ErrEmptyUpdate) {
			err = app.ErrBadRequest
		}

		return err
	}

	return nil
}

func (s *organizationService) Delete(ctx context.Context, organizationID uuid.UUID) error {
	err := s.organizationRepo.Delete(ctx, organizationID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = app.ErrNotFound
		}

		return err
	}

	return nil
}

func (s *organizationService) Deposit(ctx context.Context, auth app.Auth, organizationID uuid.UUID, amount int64) error {
	if auth.User.Role != entity.SystemManagerRole {
		return app.ErrForbidden
	}

	if amount <= 0 {
		return app.ErrBadValue
	}

	_, err := s.organizationRepo.Get(ctx, organizationID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = app.ErrNotFound
		}

		return err
	}

	transaction := dbmodels.Transaction{
		ID:             uuid.NewV4(),
		OrganizationID: organizationID,
		Amount:         amount,
		Operation:      dbmodels.DepositOperation,
		CreatedAt:      time.Now().UTC(),
	}

	_, err = s.transactionRepo.Create(ctx, transaction)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return app.ErrNotFound
		}
		if errors.Is(err, dbmodels.ErrAlreadyExists) {
			return app.ErrAlreadyExists
		}
		if errors.Is(err, dbmodels.ErrInsufficientFunds) {
			return app.ErrInsufficientFunds
		}

		return err
	}

	return nil
}

func (s *organizationService) Withdrawal(ctx context.Context, organizationID uuid.UUID, amount int64) error {
	if amount <= 0 {
		return app.ErrBadValue
	}

	organizationDB, err := s.organizationRepo.Get(ctx, organizationID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = app.ErrNotFound
		}

		return err
	}

	if organizationDB.Balance-amount < 0 {
		return app.ErrInsufficientFunds
	}

	transaction := dbmodels.Transaction{
		ID:             uuid.NewV4(),
		OrganizationID: organizationID,
		Amount:         amount,
		Operation:      dbmodels.DebitOperation,
		CreatedAt:      time.Now().UTC(),
	}

	_, err = s.transactionRepo.Create(ctx, transaction)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return app.ErrNotFound
		}
		if errors.Is(err, dbmodels.ErrAlreadyExists) {
			return app.ErrAlreadyExists
		}
		if errors.Is(err, dbmodels.ErrInsufficientFunds) {
			return app.ErrInsufficientFunds
		}

		return err
	}

	return nil
}
