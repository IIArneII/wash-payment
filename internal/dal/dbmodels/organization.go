package dbmodels

import uuid "github.com/satori/go.uuid"

type (
	Organization struct {
		ID          uuid.UUID `db:"id"`
		Name        string    `db:"name"`
		Description string    `db:"description"`
		Deleted     bool      `db:"deleted"`
	}

	OrganizationUpdate struct {
		Name        string `db:"name"`
		Description string `db:"description"`
	}
)
