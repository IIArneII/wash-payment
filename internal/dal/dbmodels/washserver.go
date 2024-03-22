package dbmodels

import (
	uuid "github.com/satori/go.uuid"
)

type (
	WashServer struct {
		ID          uuid.UUID `db:"id"`
		Title       string    `db:"title"`
		Description string    `db:"description"`
		GroupID     uuid.UUID `db:"group_id"`
		Version     int64     `db:"version"`
		Deleted     bool      `db:"deleted"`
	}

	WashServerUpdate struct {
		Version     *int64        `db:"version"`
		Title       *string       `db:"title"`
		Description *string       `db:"description"`
		GroupID     uuid.NullUUID `db:"group_id"`
		Deleted     *bool         `db:"deleted"`
	}
)
