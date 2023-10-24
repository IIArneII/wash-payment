package entity

import (
	uuid "github.com/satori/go.uuid"
)

type (
	Organization struct {
		ID          uuid.UUID
		Name        string
		DisplayName string
		Description string
		Version     int64
		Balance     int64
		Deleted     bool
	}

	OrganizationCreate struct {
		ID          uuid.UUID
		Name        string
		DisplayName string
		Description string
		Version     int64
		Deleted     bool
	}

	OrganizationUpdate struct {
		Name        *string
		DisplayName *string
		Description *string
		Version     *int64
	}
)
