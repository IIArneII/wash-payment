package entity

import (
	uuid "github.com/satori/go.uuid"
)

type (
	User struct {
		ID             string
		Email          string
		Name           string
		Role           Role
		OrganizationID *uuid.UUID
		Version        int64
	}

	UserUpdate struct {
		Role    *Role
		Name    *string
		Email   *string
		Version *int64
	}

	Role string
)

const (
	SystemManagerRole Role = "systemManager"
	AdminRole         Role = "admin"
	NoAccessRole      Role = "noAccess"
)
