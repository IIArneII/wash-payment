package entity

import (
	uuid "github.com/satori/go.uuid"
)

type (
	ServicePrices struct {
		Bonus int64
		Sbp   int64
	}

	Organization struct {
		ID            uuid.UUID
		Name          string
		DisplayName   string
		Description   string
		Version       int64
		Balance       int64
		Deleted       bool
		ServicePrices ServicePrices
	}

	OrganizationUpdate struct {
		Name        *string
		DisplayName *string
		Description *string
		Version     *int64
		Deleted     *bool
	}

	OrganizationFilter struct {
		Filter
	}
)
