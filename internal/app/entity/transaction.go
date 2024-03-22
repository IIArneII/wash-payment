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
		ForDate        *time.Time
		Service        Service
		StationsСount  *int
		UserID         *string
		Group          *Group
		WashServer     *WashServer
	}

	TransactionCreate struct {
		ID             uuid.UUID
		OrganizationID uuid.UUID
		Amount         int64
		Operation      Operation
		CreatedAt      time.Time
		ForDate        *time.Time
		Service        Service
		StationsСount  *int
		UserID         *string
		GroupID        *uuid.UUID
		WashServerID   *uuid.UUID
	}

	Withdrawal struct {
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
