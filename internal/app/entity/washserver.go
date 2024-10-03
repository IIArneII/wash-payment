package entity

import (
	uuid "github.com/satori/go.uuid"
)

type (
	WashServer struct {
		ID          uuid.UUID
		Name        string
		Description string
		OwnerID     string
		GroupID     uuid.UUID
		Version     int
		Deleted     bool
	}

	WashServerUpdate struct {
		Version     *int
		Name        *string
		Description *string
		GroupID     *uuid.UUID
		OwnerID     *string
		Deleted     *bool
	}
)
