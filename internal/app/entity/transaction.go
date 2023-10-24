package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type (
	Transaction struct {
		ID             uuid.UUID
		OrganizationID uuid.UUID
		Amount         int64
		Operation      Operation
		CreatedAt      time.Time
	}

	Operation string
)

const (
	DepositOperation Operation = "deposit"
	DebitOperation   Operation = "debit"
)
