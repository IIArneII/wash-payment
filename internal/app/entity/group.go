package entity

import (
	uuid "github.com/satori/go.uuid"
)

type (
	Group struct {
		ID             uuid.UUID
		OrganizationID uuid.UUID
		Name           string
		Description    string
		Version        int
		Deleted        bool
	}

	GroupUpdate struct {
		Version     *int
		Name        *string
		Description *string
		Deleted     *bool
	}
)
