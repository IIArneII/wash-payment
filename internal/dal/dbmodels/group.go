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
		Version        int64     `db:"version"`
		Deleted        bool      `db:"deleted"`
	}

	GroupUpdate struct {
		Version     *int64  `db:"version"`
		Name        *string `db:"name"`
		Description *string `db:"description"`
	}
)
