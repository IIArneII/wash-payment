package serviceprice

import (
	"context"
	"errors"
	"fmt"
	"wash-payment/internal/app"
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/conversions"
	"wash-payment/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

var columns = []string{"organization_id", "service", "price"}

func (r *servicePriceRepo) Get(ctx context.Context, organizationID uuid.UUID, service entity.Service) (entity.ServicePrice, error) {
	op := "failed to get service price by ID: %w"

	var dbServicePrice dbmodels.ServicePrice
	err := r.db.NewSession(nil).
		Select(columns...).
		From(dbmodels.ServicePricesTable).
		Where(dbmodels.ByOrgIDAndSvcCondition, organizationID, service).
		LoadOneContext(ctx, &dbServicePrice)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			err = app.ErrNotFound
		}
		return entity.ServicePrice{}, fmt.Errorf(op, err)
	}

	return conversions.ServicePriceFromDB(dbServicePrice), nil
}

func (r *servicePriceRepo) Create(ctx context.Context, servicePrice entity.ServicePrice) (entity.ServicePrice, error) {
	op := "failed to create service price: %w"

	dbServicePrice := conversions.ServicePriceToDB(servicePrice)
	var dbCreatedServicePrice dbmodels.ServicePrice
	err := r.db.NewSession(nil).
		InsertInto(dbmodels.ServicePricesTable).
		Columns(columns...).
		Record(dbServicePrice).
		Returning(columns...).
		LoadContext(ctx, &dbCreatedServicePrice)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == dbmodels.PQErrAlreadyExists {
			err = app.ErrAlreadyExists
		}
		return entity.ServicePrice{}, fmt.Errorf(op, err)
	}

	return conversions.ServicePriceFromDB(dbCreatedServicePrice), nil
}

func (r *servicePriceRepo) Update(ctx context.Context, organizationID uuid.UUID, service entity.Service, price int64) (entity.ServicePrice, error) {
	op := "failed to update service price: %w"

	result, err := r.db.NewSession(nil).
		Update(dbmodels.ServicePricesTable).
		Where(dbmodels.ByOrgIDAndSvcCondition, organizationID, service).
		Set("price", price).
		ExecContext(ctx)

	if err != nil {
		if errors.Is(err, dbr.ErrColumnNotSpecified) {
			err = app.ErrEmptyUpdate
		}
		return entity.ServicePrice{}, fmt.Errorf(op, err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return entity.ServicePrice{}, fmt.Errorf(op, err)
	}
	if count == 0 {
		return entity.ServicePrice{}, fmt.Errorf(op, app.ErrNotFound)
	}

	return r.Get(ctx, organizationID, service)
}
