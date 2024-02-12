package organization

import (
	"context"
	"errors"
	"time"
	"wash-payment/internal/app"
	"wash-payment/internal/app/entity"

	uuid "github.com/satori/go.uuid"
)

func (s *organizationService) Get(ctx context.Context, auth entity.Auth, organizationID uuid.UUID) (entity.Organization, error) {
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

	organization, err := s.organizationRepo.Get(ctx, organizationID)
	if err != nil {
		return entity.Organization{}, err
	}
	if organization.Deleted {
		return entity.Organization{}, app.ErrNotFound
	}

	return organization, nil
}

func (s *organizationService) List(ctx context.Context, auth entity.Auth, filter entity.OrganizationFilter) (entity.Page[entity.Organization], error) {
	if auth.User.Role != entity.SystemManagerRole {
		return entity.Page[entity.Organization]{}, app.ErrForbidden
	}

	orgs, err := s.organizationRepo.List(ctx, filter)
	if err != nil {
		return entity.Page[entity.Organization]{}, err
	}

	return orgs, nil
}

func (s *organizationService) Upsert(ctx context.Context, organization entity.Organization) (entity.Organization, error) {
	if organization.ID == uuid.Nil {
		return entity.Organization{}, app.ErrNotFound
	}

	dbOrg, err := s.organizationRepo.Get(ctx, organization.ID)
	if err != nil {
		if errors.Is(err, app.ErrNotFound) {
			organization.Balance = 0
			newOrganization, err := s.organizationRepo.Create(ctx, organization)
			if err != nil {
				return entity.Organization{}, err
			}
			return newOrganization, nil
		}
		return entity.Organization{}, err
	} else {
		if dbOrg.Version >= organization.Version {
			return entity.Organization{}, app.ErrOldVersion
		}

		organizationUpdate := organizationToUpdate(organization)
		updatedOrg, err := s.organizationRepo.Update(ctx, organization.ID, organizationUpdate)
		if err != nil {
			return entity.Organization{}, err
		}

		return updatedOrg, nil
	}
}

func (s *organizationService) Deposit(ctx context.Context, auth entity.Auth, organizationID uuid.UUID, amount int64) error {
	if auth.User.Role != entity.SystemManagerRole {
		return app.ErrForbidden
	}

	if amount <= 0 {
		return app.ErrBadValue
	}

	_, err := s.organizationRepo.Get(ctx, organizationID)
	if err != nil {
		return err
	}

	_, err = s.transactionRepo.Create(ctx, entity.Transaction{
		ID:             uuid.NewV4(),
		OrganizationID: organizationID,
		Amount:         amount,
		Operation:      entity.DepositOperation,
		CreatedAt:      time.Now().UTC(),
	})
	return err
}

func (s *organizationService) Withdrawal(ctx context.Context, organizationID uuid.UUID, amount int64, service_name string) error {
	if amount <= 0 {
		return app.ErrBadValue
	}

	organizationDB, err := s.organizationRepo.Get(ctx, organizationID)
	if err != nil {
		return err
	}

	if organizationDB.Balance-amount < 0 {
		return app.ErrInsufficientFunds
	}

	_, err = s.transactionRepo.Create(ctx, entity.Transaction{
		ID:             uuid.NewV4(),
		OrganizationID: organizationID,
		Amount:         amount,
		Operation:      entity.DebitOperation,
		CreatedAt:      time.Now().UTC(),
		Sevice:         &service_name,
	})
	if err != nil {
		return err
	}

	return nil
}

func organizationToUpdate(org entity.Organization) entity.OrganizationUpdate {
	return entity.OrganizationUpdate{
		Name:        &org.Name,
		DisplayName: &org.DisplayName,
		Description: &org.Description,
		Version:     &org.Version,
		Deleted:     &org.Deleted,
	}
}
