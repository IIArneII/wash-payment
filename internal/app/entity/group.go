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
		Version        int64
		Deleted        bool
	}

	GroupUpdate struct {
		Version     *int64
		Name        *string
		Description *string
		Deleted     *bool
	}
)
