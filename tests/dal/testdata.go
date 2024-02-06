package dal

import (
	"time"
	"wash-payment/internal/app/entity"

	"github.com/Pallinder/go-randomdata"
	uuid "github.com/satori/go.uuid"
)

func generateUser(role entity.Role, organizationID *uuid.UUID, version int) entity.User {
	return entity.User{
		ID:             uuid.NewV4().String(),
		Email:          randomdata.Email(),
		Name:           randomdata.FullName(randomdata.RandomGender),
		Role:           role,
		OrganizationID: organizationID,
		Version:        int64(version),
	}
}

func generateOrganization(balance int64, version int) entity.Organization {
	return entity.Organization{
		ID:          uuid.NewV4(),
		Name:        randomdata.FirstName(randomdata.Male),
		DisplayName: uuid.NewV4().String(),
		Description: randomdata.RandStringRunes(50),
		Deleted:     false,
		Version:     int64(version),
		Balance:     balance,
	}
}

func generateGroup(organizationID uuid.UUID, version int) entity.Group {
	return entity.Group{
		ID:             uuid.NewV4(),
		OrganizationID: organizationID,
		Name:           randomdata.FirstName(randomdata.Male),
		Description:    randomdata.RandStringRunes(50),
		Deleted:        false,
		Version:        int64(version),
	}
}

func generateTransaction(operation entity.Operation, amount int64, organizationID uuid.UUID) entity.Transaction {
	service := randomdata.City()
	return entity.Transaction{
		ID:             uuid.NewV4(),
		OrganizationID: organizationID,
		Amount:         amount,
		Operation:      operation,
		CreatedAt:      time.Now().UTC().Truncate(time.Millisecond),
		Sevice:         &service,
	}
}
