package conversions

import (
	"wash-payment/internal/dal/dbmodels"
	"wash-payment/internal/entity"

	uuid "github.com/satori/go.uuid"
)

func UserFromDb(dbUser dbmodels.User) entity.User {
	var orgID *uuid.UUID
	if dbUser.OrganizationID.Valid {
		orgID = &dbUser.OrganizationID.UUID
	}

	return entity.User{
		ID:             dbUser.ID,
		Name:           dbUser.Name,
		Email:          dbUser.Email,
		Role:           RoleSelectionApp(dbUser.Role),
		OrganizationID: orgID,
	}
}

func RoleSelectionApp(role dbmodels.Role) entity.Role {
	switch role {
	case dbmodels.AdminRole:
		return entity.AdminRole
	case dbmodels.SystemManagerRole:
		return entity.SystemManagerRole
	default:
		panic("Unknown db role: " + role)
	}
}
