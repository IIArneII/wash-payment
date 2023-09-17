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

	UserCreation struct {
		ID             string
		Email          string
		Name           string
		OrganizationID *uuid.UUID
	}

	UserUpdate struct {
		ID    string
		Email *string
		Name  *string
	}

	Role string
)

const (
	SystemManagerRole Role = "systemManager"
	AdminRole         Role = "admin"
)
