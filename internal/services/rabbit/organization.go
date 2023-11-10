package rabbit

import (
	"context"
	"wash-payment/internal/app/conversions"
	"wash-payment/internal/app/entity"
	et "wash-payment/internal/transport/rabbit/entity"

	uuid "github.com/satori/go.uuid"
)

func (s *rabbitService) UpsertOrganization(ctx context.Context, organization et.Organization) error {
	organizationCreate, err := conversions.OrganizationCreateFromRabbit(organization)
	if err != nil {
		return err
	}

	_, err = s.services.OrganizationService.Upsert(ctx, uuid.Nil, organizationCreate, entity.OrganizationUpdate{})
	if err != nil {
		return err
	}

	return nil
}

// NEW
func (s *rabbitService) ProcessWithdrawal(ctx context.Context, organization et.Organization, amount int64) error {

	organizationCreate, err := conversions.OrganizationCreateFromRabbit(organization)
	if err != nil {
		return err
	}

	err = s.services.OrganizationService.Withdrawal(ctx, organizationCreate.ID, amount)

	return err
}
