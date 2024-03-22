package entity

import (
	uuid "github.com/satori/go.uuid"
)

type (
	WashServer struct {
		ID          uuid.UUID
		Title       string
		Description string
		GroupID     uuid.UUID
		Version     int64
		Deleted     bool
	}

	WashServerUpdate struct {
		Version     *int64
		Title       *string
		Description *string
		GroupID     *uuid.UUID
		Deleted     *bool
	}
)
