package dbmodels

import uuid "github.com/satori/go.uuid"

type (
	User struct {
		ID             string     `db:"id"`
		Email          *string    `db:"email"`
		Name           *string    `db:"name"`
		Role           Role       `db:"role"`
		OrganizationID *uuid.UUID `db:"organization_id"`
		Deleted        bool       `db:"deleted"`
	}

	UserCreation struct {
		ID             string     `db:"id"`
		Email          string     `db:"email"`
		Name           string     `db:"name"`
		OrganizationId *uuid.UUID `db:"organization_id"`
	}

	UserUpdate struct {
		ID    string  `db:"id"`
		Email *string `db:"email"`
		Name  *string `db:"name"`
	}
)

type Role string

const (
	SystemManagerRole Role = "system_manager"
	AdminRole         Role = "admin"
)
