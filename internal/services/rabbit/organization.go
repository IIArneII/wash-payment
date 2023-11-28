package rabbit

import (
	"context"
	"wash-payment/internal/app/conversions"

	"wash-payment/internal/transport/rabbit/entity"

	uuid "github.com/satori/go.uuid"
)

func (s *rabbitService) UpsertOrganization(ctx context.Context, rabbitOrganization entity.Organization) error {
	organizationCreate, err := conversions.OrganizationCreateFromRabbit(rabbitOrganization)
	if err != nil {
		return err
	}

	_, err = s.services.OrganizationService.Upsert(ctx, organizationCreate)
	if err != nil {
		return err
	}

	return nil
}

// NEW
func (s *rabbitService) ProcessWithdrawal(ctx context.Context, payment entity.Payment) error {

	organisationId, err := uuid.FromString(payment.OrganizationId)
	if err != nil {
		return err
	}

	err = s.services.OrganizationService.Withdrawal(ctx, organisationId, payment.Amount)
	if err != nil {
		return err
	}

	return nil
}
