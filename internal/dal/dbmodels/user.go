package dbmodels

import uuid "github.com/satori/go.uuid"

type (
	User struct {
		ID             string        `db:"id"`
		Email          string        `db:"email"`
		Name           string        `db:"name"`
		Role           Role          `db:"role"`
		OrganizationID uuid.NullUUID `db:"organization_id"`
		Version        int64         `db:"version"`
	}

	UserUpdate struct {
		Role    *Role   `db:"role"`
		Name    *string `db:"name"`
		Email   *string `db:"email"`
		Version *int64  `db:"version"`
	}

	Role string
)

const (
	SystemManagerRole Role = "system_manager"
	AdminRole         Role = "admin"
	NoAccessRole      Role = "no_access"
)
