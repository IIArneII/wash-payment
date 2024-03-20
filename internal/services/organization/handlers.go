package organization

import (
	"context"
	"errors"
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

func (s *organizationService) SetServicePrices(ctx context.Context, auth entity.Auth, organizationID uuid.UUID, servicePrices entity.ServicePrices) error {
	if auth.User.Role != entity.SystemManagerRole {
		return app.ErrForbidden
	}

	if servicePrices.Bonus < 0 || servicePrices.Sbp < 0 {
		return app.ErrBadValue
	}

	_, err := s.organizationRepo.Get(ctx, organizationID)
	if err != nil {
		return err
	}

	_, err = s.servicePriceRepo.Update(ctx, organizationID, entity.BonusService, servicePrices.Bonus)
	if err != nil {
		return err
	}

	_, err = s.servicePriceRepo.Update(ctx, organizationID, entity.SbpService, servicePrices.Sbp)
	if err != nil {
		return err
	}

	return nil
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

			err = s.servicePricesForCreatedOrganization(ctx, newOrganization.ID)
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

func (s *organizationService) servicePricesForCreatedOrganization(ctx context.Context, organizationID uuid.UUID) error {
	_, err := s.servicePriceRepo.Create(ctx, entity.ServicePrice{
		OrganizationID: organizationID,
		Service:        entity.BonusService,
		Price:          0,
	})
	if err != nil {
		return err
	}

	_, err = s.servicePriceRepo.Create(ctx, entity.ServicePrice{
		OrganizationID: organizationID,
		Service:        entity.SbpService,
		Price:          0,
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
