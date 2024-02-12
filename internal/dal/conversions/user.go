package conversions

import (
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"

	uuid "github.com/satori/go.uuid"
)

func UserFromDB(dbUser dbmodels.User) entity.User {
	var orgID *uuid.UUID
	if dbUser.OrganizationID.Valid {
		orgID = &dbUser.OrganizationID.UUID
	}

	return entity.User{
		ID:             dbUser.ID,
		Name:           dbUser.Name,
		Email:          dbUser.Email,
		Role:           RoleFromDB(dbUser.Role),
		OrganizationID: orgID,
		Version:        dbUser.Version,
	}
}

func UserToDB(appUser entity.User) dbmodels.User {
	var orgID uuid.NullUUID
	if appUser.OrganizationID != nil {
		orgID.UUID = *appUser.OrganizationID
		orgID.Valid = true
	}

	return dbmodels.User{
		ID:             appUser.ID,
		Name:           appUser.Name,
		Email:          appUser.Email,
		Role:           RoleToDB(appUser.Role),
		OrganizationID: orgID,
		Version:        appUser.Version,
	}
}

func UserUpdateToDB(appUserUpdate entity.UserUpdate) dbmodels.UserUpdate {
	userUpdate := dbmodels.UserUpdate{}

	if appUserUpdate.Email != nil {
		userUpdate.Email = appUserUpdate.Email
	}
	if appUserUpdate.Name != nil {
		userUpdate.Name = appUserUpdate.Name
	}
	if appUserUpdate.Version != nil {
		userUpdate.Version = appUserUpdate.Version
	}
	if appUserUpdate.Role != nil {
		newRole := RoleToDB(*appUserUpdate.Role)
		userUpdate.Role = &newRole
	}

	return userUpdate
}

func RoleFromDB(role dbmodels.Role) entity.Role {
	switch role {
	case dbmodels.AdminRole:
		return entity.AdminRole
	case dbmodels.SystemManagerRole:
		return entity.SystemManagerRole
	case dbmodels.NoAccessRole:
		return entity.NoAccessRole
	default:
		panic("Unknown db role: " + role)
	}
}

func RoleToDB(role entity.Role) dbmodels.Role {
	switch role {
	case entity.AdminRole:
		return dbmodels.AdminRole
	case entity.SystemManagerRole:
		return dbmodels.SystemManagerRole
	case entity.NoAccessRole:
		return dbmodels.NoAccessRole
	default:
		panic("Unknown app role: " + role)
	}
}

func RoleFromRabbit(role string) entity.Role {
	switch role {
	case "admin":
		return entity.AdminRole
	case "systemManager":
		return entity.SystemManagerRole
	case "system_manager":
		return entity.SystemManagerRole
	case "noAccess":
		return entity.NoAccessRole
	default:
		panic("Unknown rabbit role: " + role)
	}
}
