package conversions

import (
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"
)

func ServiceFromDb(service dbmodels.Service) entity.Service {
	switch service {
	case dbmodels.PaymentService:
		return entity.PaymentService
	case dbmodels.BonusService:
		return entity.BonusService
	case dbmodels.SbpService:
		return entity.SbpService
	default:
		panic("Unknown db service: " + service)
	}
}

func ServiceToDb(operation entity.Service) dbmodels.Service {
	switch operation {
	case entity.PaymentService:
		return dbmodels.PaymentService
	case entity.BonusService:
		return dbmodels.BonusService
	case entity.SbpService:
		return dbmodels.SbpService
	default:
		panic("Unknown app service: " + operation)
	}
}

func ServicePriceFromDB(sp dbmodels.ServicePrice) entity.ServicePrice {
	return entity.ServicePrice{
		OrganizationID: sp.OrganizationID,
		Service:        ServiceFromDb(sp.Service),
		Price:          sp.Price,
	}
}

func ServicePriceToDB(sp entity.ServicePrice) dbmodels.ServicePrice {
	return dbmodels.ServicePrice{
		OrganizationID: sp.OrganizationID,
		Service:        ServiceToDb(sp.Service),
		Price:          sp.Price,
	}
}
