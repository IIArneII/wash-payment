package entity

import (
	uuid "github.com/satori/go.uuid"
)

type (
	User struct {
		ID             string
		Email          *string
		Name           *string
		Role           Role
		OrganizationID *uuid.UUID
		Deleted        bool
	}
)

type Role string

const (
	SystemManagerRole Role = "systemManager"
	AdminRole         Role = "admin"
)
