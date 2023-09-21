package dbmodels

import uuid "github.com/satori/go.uuid"

type (
	User struct {
		ID             string        `db:"id"`
		Email          string        `db:"email"`
		Name           string        `db:"name"`
		Role           Role          `db:"role"`
		OrganizationID uuid.NullUUID `db:"organization_id"`
	}
)

type Role string

const (
	SystemManagerRole Role = "system_manager"
	AdminRole         Role = "admin"
)
