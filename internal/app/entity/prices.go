package entity

import (
	uuid "github.com/satori/go.uuid"
)

type (
	ServicePrice struct {
		OrganizationID uuid.UUID
		Service        Service
		Price          int64
	}
)
