package dbmodels

import (
	uuid "github.com/satori/go.uuid"
)

type (
	Group struct {
		ID             uuid.UUID `db:"id"`
		OrganizationID uuid.UUID `db:"organization_id"`
		Name           string    `db:"name"`
		Description    string    `db:"description"`
		Version        int       `db:"version"`
		Deleted        bool      `db:"deleted"`
	}

	GroupUpdate struct {
		Version     *int    `db:"version"`
		Name        *string `db:"name"`
		Description *string `db:"description"`
		Deleted     *bool   `db:"deleted"`
	}
)
