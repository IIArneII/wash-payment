package rabbit

import (
	"context"
	"errors"
	"strconv"
	"wash-payment/internal/app/conversions"

	globalEntity "wash-payment/internal/app/entity"
	"wash-payment/internal/transport/rabbit/entity"

	uuid "github.com/satori/go.uuid"
)

func (s *rabbitService) UpsertOrganization(ctx context.Context, organization entity.Organization) error {
	organizationCreate, err := conversions.OrganizationCreateFromRabbit(organization)
	if err != nil {
		return err
	}

	_, err = s.services.OrganizationService.Upsert(ctx, uuid.Nil, organizationCreate, globalEntity.OrganizationUpdate{})
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
	amount, err := strconv.ParseInt(payment.Amount, 10, 64)
	if err != nil {
		return err
	}

	err = s.services.OrganizationService.Withdrawal(ctx, organisationId, amount)
	if err != nil {
		return err
	}

	return errors.New("Успех")
}
