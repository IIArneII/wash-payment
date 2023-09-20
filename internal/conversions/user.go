package conversions

import (
	"wash-payment/internal/dal/dbmodels"
	"wash-payment/internal/entity"
)

func UserFromDb(dbUser dbmodels.User) entity.User {
	return entity.User{
		ID:             dbUser.ID,
		Name:           dbUser.Name,
		Email:          dbUser.Email,
		Role:           RoleSelectionApp(dbUser.Role),
		OrganizationID: dbUser.OrganizationID,
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
