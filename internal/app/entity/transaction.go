package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type (
	Transaction struct {
		ID             uuid.UUID
		OrganizationID uuid.UUID
		GroupID        *uuid.UUID
		Amount         int64
		Operation      Operation
		CreatedAt      time.Time
		Sevice         *Service
		PostСount      *int
		UserID         *string
	}

	Withdrawal struct {
		GroupId   uuid.UUID
		PostСount int
		Amount    int64
		Service   Service
	}

	TransactionFilter struct {
		Filter
		OrganizationID uuid.UUID
	}

	Service   string
	Operation string
)

const (
	DepositOperation Operation = "deposit"
	DebitOperation   Operation = "debit"

	BonusService Service = "bonus"
	SbpService   Service = "sbp"
)
