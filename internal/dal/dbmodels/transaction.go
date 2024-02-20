package dbmodels

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type (
	Transaction struct {
		ID             uuid.UUID     `db:"id"`
		OrganizationID uuid.UUID     `db:"organization_id"`
		GroupID        uuid.NullUUID `db:"group_id"`
		Amount         int64         `db:"amount"`
		Operation      Operation     `db:"operation"`
		CreatedAt      time.Time     `db:"created_at"`
		Service        *Service      `db:"service"`
		StationsСount  *int          `db:"stations_count"`
		UserID         *string       `db:"user_id"`
	}

	Operation string
	Service   string
)

const (
	DepositOperation Operation = "deposit"
	DebitOperation   Operation = "debit"

	BonusService Service = "bonus"
	SbpService   Service = "sbp"
)
