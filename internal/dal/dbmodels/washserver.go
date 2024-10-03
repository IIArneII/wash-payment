package dbmodels

import (
	uuid "github.com/satori/go.uuid"
)

type (
	WashServer struct {
		ID          uuid.UUID `db:"id"`
		Name        string    `db:"name"`
		Description string    `db:"description"`
		GroupID     uuid.UUID `db:"group_id"`
		OwnerID     string    `db:"owner_id"`
		Version     int       `db:"version"`
		Deleted     bool      `db:"deleted"`
	}

	WashServerUpdate struct {
		Version     *int          `db:"version"`
		Name        *string       `db:"name"`
		Description *string       `db:"description"`
		GroupID     uuid.NullUUID `db:"group_id"`
		OwnerID     *string       `db:"owner_id"`
		Deleted     *bool         `db:"deleted"`
	}
)
