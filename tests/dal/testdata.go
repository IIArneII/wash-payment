package dal

import (
	"time"
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
		Version:        version,
	}
}

func generateOrganization(balance int64, version int) dbmodels.Organization {
	return dbmodels.Organization{
		ID:          uuid.NewV4(),
		Name:        randomdata.FirstName(randomdata.Male),
		DisplayName: uuid.NewV4().String(),
		Description: randomdata.RandStringRunes(50),
		Deleted:     false,
		Version:     version,
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
		Version:        version,
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
