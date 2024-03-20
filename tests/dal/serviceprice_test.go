package dal

import (
	"testing"
	"wash-payment/internal/app"
	"wash-payment/internal/app/entity"

	"github.com/powerman/check"
)

func TestCreateServicePrice(tt *testing.T) {
	t := check.T(tt)
	err := truncate()
	t.Nil(err)

	var price int64 = 100

	organization1 := generateOrganization(10000, 1)
	servicePriceBonus := generateServicePrice(organization1.ID, entity.BonusService, price)
	servicePriceSbp := generateServicePrice(organization1.ID, entity.SbpService, price)

	org, err := repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)
	t.Equal(org.ServicePrices, entity.ServicePrices{})

	sp, err := repositories.ServicePriceRepo.Create(ctx, servicePriceBonus)
	t.Nil(err)
	t.Equal(sp, servicePriceBonus)

	org, err = repositories.OrganizationRepo.Get(ctx, org.ID)
	t.Nil(err)
	t.Equal(org.ServicePrices, entity.ServicePrices{Bonus: price})

	sp, err = repositories.ServicePriceRepo.Create(ctx, servicePriceSbp)
	t.Nil(err)
	t.Equal(sp, servicePriceSbp)

	org, err = repositories.OrganizationRepo.Get(ctx, org.ID)
	t.Nil(err)
	t.Equal(org.ServicePrices, entity.ServicePrices{Bonus: price, Sbp: price})

	_, err = repositories.ServicePriceRepo.Create(ctx, servicePriceSbp)
	t.Err(err, app.ErrAlreadyExists)
}

func TestGetServicePrice(tt *testing.T) {
	t := check.T(tt)
	err := truncate()
	t.Nil(err)

	organization1 := generateOrganization(10000, 1)
	servicePriceBonus := generateServicePrice(organization1.ID, entity.BonusService, 100)

	_, err = repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	_, err = repositories.ServicePriceRepo.Create(ctx, servicePriceBonus)
	t.Nil(err)

	sp, err := repositories.ServicePriceRepo.Get(ctx, servicePriceBonus.OrganizationID, servicePriceBonus.Service)
	t.Nil(err)
	t.Equal(sp, servicePriceBonus)

	_, err = repositories.ServicePriceRepo.Get(ctx, servicePriceBonus.OrganizationID, entity.SbpService)
	t.Err(err, app.ErrNotFound)
}

func TestUpdateServicePrice(tt *testing.T) {
	t := check.T(tt)
	err := truncate()
	t.Nil(err)

	organization1 := generateOrganization(10000, 1)
	servicePriceBonus := generateServicePrice(organization1.ID, entity.BonusService, 100)

	_, err = repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	_, err = repositories.ServicePriceRepo.Create(ctx, servicePriceBonus)
	t.Nil(err)

	var newPrice int64 = 150
	servicePriceBonus.Price = newPrice

	sp, err := repositories.ServicePriceRepo.Update(ctx, servicePriceBonus.OrganizationID, servicePriceBonus.Service, newPrice)
	t.Nil(err)
	t.Equal(sp, servicePriceBonus)

	org, err := repositories.OrganizationRepo.Get(ctx, organization1.ID)
	t.Nil(err)
	t.Equal(org.ServicePrices, entity.ServicePrices{Bonus: newPrice})

	_, err = repositories.ServicePriceRepo.Update(ctx, servicePriceBonus.OrganizationID, entity.SbpService, newPrice)
	t.Err(err, app.ErrNotFound)
}
