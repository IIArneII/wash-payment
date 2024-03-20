package rabbit

import (
	"context"
	"wash-payment/internal/app/entity"
	rabbitEntity "wash-payment/internal/transport/rabbit/entity"

	uuid "github.com/satori/go.uuid"
)

func (s *rabbitService) UpsertOrganization(ctx context.Context, rabbitOrganization rabbitEntity.Organization) error {
	organizationCreate, err := organizationCreateFromRabbit(rabbitOrganization)
	if err != nil {
		return err
	}

	_, err = s.services.OrganizationService.Upsert(ctx, organizationCreate)
	if err != nil {
		return err
	}

	return nil
}

func (s *rabbitService) Withdrawal(ctx context.Context, withdrawal rabbitEntity.Withdrawal) error {
	groupId, err := uuid.FromString(withdrawal.GroupId)
	if err != nil {
		return err
	}

	err = s.services.TransactionService.Withdrawal(ctx, entity.Withdrawal{
		GroupId:       groupId,
		StationsСount: withdrawal.StationsСount,
		Service:       serviceFromRabbit(withdrawal.Service),
		ForDate:       withdrawal.ForDate,
	})
	if err != nil {
		return err
	}

	return nil
}

func organizationCreateFromRabbit(org rabbitEntity.Organization) (entity.Organization, error) {
	id, err := uuid.FromString(org.ID)
	if err != nil {
		return entity.Organization{}, err
	}

	return entity.Organization{
		ID:          id,
		Name:        org.Name,
		DisplayName: org.DisplayName,
		Description: org.Description,
		Version:     org.Version,
		Deleted:     org.Deleted,
	}, nil
}

func serviceFromRabbit(service string) entity.Service {
	switch service {
	case "bonus":
		return entity.BonusService
	case "sbp":
		return entity.SbpService
	default:
		panic("Unknown rabbit service: " + service)
	}
}
