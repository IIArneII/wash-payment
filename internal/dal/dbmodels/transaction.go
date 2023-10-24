package dbmodels

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type (
	Transaction struct {
		ID             uuid.UUID `db:"id"`
		OrganizationID uuid.UUID `db:"organization_id"`
		Amount         int64     `db:"amount"`
		Operation      Operation `db:"operation"`
		CreatedAt      time.Time `db:"created_at"`
	}

	Operation string
)

const (
	DepositOperation Operation = "deposit"
	DebitOperation   Operation = "debit"
)
