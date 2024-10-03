package dbmodels

import uuid "github.com/satori/go.uuid"

type (
	ServicePrice struct {
		OrganizationID uuid.UUID `db:"organization_id"`
		Service        Service   `db:"service"`
		Price          int64     `db:"price"`
	}
)
