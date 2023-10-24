package rabbit

import (
	"context"
	"wash-payment/internal/app/conversions"
	"wash-payment/internal/transport/rabbit/entity"
)

func (s *rabbitService) UpsertOrganization(ctx context.Context, organization entity.Organization) error {
	organizationCreate, err := conversions.OrganizationCreateFromRabbit(organization)
	if err != nil {
		return err
	}

	_, err = s.services.OrganizationService.Create(ctx, organizationCreate)
	if err != nil {
		return err
	}

	return nil
}
