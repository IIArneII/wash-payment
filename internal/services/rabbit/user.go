package rabbit

import (
	"context"
	"wash-payment/internal/app/entity"
	rabbitEntity "wash-payment/internal/transport/rabbit/entity"

	uuid "github.com/satori/go.uuid"
)

func (s *rabbitService) UpsertUser(ctx context.Context, rabbitUser rabbitEntity.User) error {
	user, err := userFromRabbit(rabbitUser)
	if err != nil {
		return err
	}

	_, err = s.services.UserService.Upsert(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func userFromRabbit(rabbitUser rabbitEntity.User) (entity.User, error) {
	var orgId *uuid.UUID
	if rabbitUser.OrganizationID != nil {
		orgIdfromStr, err := uuid.FromString(*rabbitUser.OrganizationID)
		if err != nil {
			return entity.User{}, err
		}
		orgId = &orgIdfromStr
	}

	return entity.User{
		ID:             rabbitUser.ID,
		Email:          rabbitUser.Email,
		Name:           rabbitUser.Name,
		OrganizationID: orgId,
		Version:        rabbitUser.Version,
		Role:           roleFromRabbit(rabbitUser.Role),
	}, nil
}

func roleFromRabbit(role string) entity.Role {
	switch role {
	case "admin":
		return entity.AdminRole
	case "systemManager":
		return entity.SystemManagerRole
	case "system_manager":
		return entity.SystemManagerRole //УБРАТЬ
	case "no_access":
		return entity.NoAccessRole
	case "noAccess":
		return entity.NoAccessRole
	default:
		panic("Unknown rabbit role: " + role)
	}
}
