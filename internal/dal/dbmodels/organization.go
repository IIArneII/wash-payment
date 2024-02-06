package dbmodels

import uuid "github.com/satori/go.uuid"

type (
	Organization struct {
		ID          uuid.UUID `db:"id"`
		Name        string    `db:"name"`
		DisplayName string    `db:"display_name"`
		Description string    `db:"description"`
		Version     int64     `db:"version"`
		Balance     int64     `db:"balance"`
		Deleted     bool      `db:"deleted"`
	}

	OrganizationUpdate struct {
		Name        *string `db:"name"`
		DisplayName *string `db:"display_name"`
		Description *string `db:"description"`
		Version     *int64  `db:"version"`
		Deleted     *bool   `db:"deleted"`
	}
)
