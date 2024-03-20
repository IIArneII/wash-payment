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
		ServicePrices: entity.ServicePrices{
			Bonus: 0,
			Sbp:   0,
		},
	}
}

func generateServicePrice(organizationID uuid.UUID, service entity.Service, price int64) entity.ServicePrice {
	return entity.ServicePrice{
		OrganizationID: organizationID,
		Service:        service,
		Price:          price,
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

func generateTransactionDeposit(amount int64, organizationID uuid.UUID, userID string) entity.Transaction {
	return entity.Transaction{
		ID:             uuid.NewV4(),
		OrganizationID: organizationID,
		Amount:         amount,
		Operation:      entity.DepositOperation,
		CreatedAt:      time.Now().UTC().Truncate(time.Millisecond),
		UserID:         &userID,
		Service:        entity.PaymentService,
	}
}

func generateTransactionDebit(amount int64, organizationID uuid.UUID, groupID uuid.UUID) entity.Transaction {
	stationsСount := 5
	forDate := time.Now().UTC().Truncate(24 * time.Hour)
	return entity.Transaction{
		ID:             uuid.NewV4(),
		OrganizationID: organizationID,
		GroupID:        &groupID,
		Amount:         amount,
		Operation:      entity.DebitOperation,
		CreatedAt:      time.Now().UTC().Truncate(time.Millisecond),
		ForDate:        &forDate,
		Service:        entity.BonusService,
		StationsСount:  &stationsСount,
	}
}
