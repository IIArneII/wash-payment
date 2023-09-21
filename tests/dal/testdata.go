package dal

import (
	"wash-payment/internal/dal/dbmodels"

	"github.com/Pallinder/go-randomdata"
	uuid "github.com/satori/go.uuid"
)

var (
	user1 = dbmodels.User{
		ID:    uuid.NewV4().String(),
		Email: randomdata.Email(),
		Name:  randomdata.FullName(randomdata.RandomGender),
		Role:  dbmodels.SystemManagerRole,
	}

	user2 = dbmodels.User{
		ID:    uuid.NewV4().String(),
		Email: randomdata.Email(),
		Name:  randomdata.FullName(randomdata.RandomGender),
		Role:  dbmodels.AdminRole,
	}
)
