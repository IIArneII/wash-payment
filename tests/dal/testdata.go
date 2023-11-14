package dal

import (
	"time"
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"

	"github.com/Pallinder/go-randomdata"
	uuid "github.com/satori/go.uuid"
)

func generateUser(role dbmodels.Role, organizationID uuid.NullUUID, version int) dbmodels.User {
	return dbmodels.User{
		ID:             uuid.NewV4().String(),
		Email:          randomdata.Email(),
		Name:           randomdata.FullName(randomdata.RandomGender),
		Role:           role,
		OrganizationID: organizationID,
		Version:        int64(version),
	}
}

func generateOrganization(balance int64, version int) dbmodels.Organization {
	return dbmodels.Organization{
		ID:          uuid.NewV4(),
		Name:        randomdata.FirstName(randomdata.Male),
		DisplayName: uuid.NewV4().String(),
		Description: randomdata.RandStringRunes(50),
		Deleted:     false,
		Version:     int64(version),
		Balance:     balance,
	}
}

func generateGroup(organizationID uuid.UUID, version int) dbmodels.Group {
	return dbmodels.Group{
		ID:             uuid.NewV4(),
		OrganizationID: organizationID,
		Name:           randomdata.FirstName(randomdata.Male),
		Description:    randomdata.RandStringRunes(50),
		Deleted:        false,
		Version:        int64(version),
	}
}

func generateTransaction(operation dbmodels.Operation, amount int64, organizationID uuid.UUID) dbmodels.Transaction {
	return dbmodels.Transaction{
		ID:             uuid.NewV4(),
		OrganizationID: organizationID,
		Amount:         amount,
		Operation:      operation,
		CreatedAt:      time.Now().UTC().Truncate(time.Millisecond),
	}
}

func generateGroupForService(organizationID uuid.UUID, version int) entity.Group {
	return entity.Group{
		ID:             uuid.NewV4(),
		OrganizationID: organizationID,
		Name:           randomdata.FirstName(randomdata.Male),
		Description:    randomdata.RandStringRunes(50),
		Version:        int64(version),
		Deleted:        false,
	}
}

func generateGroupUpdateForService(version int64, name string, description string) entity.GroupUpdate {
	return entity.GroupUpdate{
		Version:     &version,
		Name:        &name,
		Description: &description,
	}
}

func generateUserForService()
