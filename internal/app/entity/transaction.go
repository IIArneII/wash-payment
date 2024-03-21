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
		ForDate        *time.Time
		Service        Service
		StationsСount  *int
		UserID         *string
		WashServerID   *uuid.UUID
	}

	Withdrawal struct {
		GroupId       uuid.UUID
		StationsСount int
		ForDate       time.Time
		Service       Service
		WashServerID  uuid.UUID
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

	PaymentService Service = "payment"
	BonusService   Service = "bonus"
	SbpService     Service = "sbp"
)
