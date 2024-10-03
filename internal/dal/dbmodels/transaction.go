package dbmodels

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type (
	Transaction struct {
		ID             uuid.UUID  `db:"t_id"`
		OrganizationID uuid.UUID  `db:"t_organization_id"`
		Amount         int64      `db:"t_amount"`
		Operation      Operation  `db:"t_operation"`
		CreatedAt      time.Time  `db:"t_created_at"`
		ForDate        *time.Time `db:"t_for_date"`
		Service        Service    `db:"t_service"`
		StationsCount  *int       `db:"t_stations_count"`

		GroupID             uuid.NullUUID `db:"g_id"`
		GroupOrganizationID uuid.NullUUID `db:"g_organization_id"`
		GroupName           *string       `db:"g_name"`
		GroupDescription    *string       `db:"g_description"`
		GroupVersion        *int          `db:"g_version"`
		GroupDeleted        *bool         `db:"g_deleted"`

		WashServerID          uuid.NullUUID `db:"ws_id"`
		WashServerName        *string       `db:"ws_name"`
		WashServerDescription *string       `db:"ws_description"`
		WashServerOwnerID     *string       `db:"ws_owner_id"`
		WashServerGroupID     uuid.NullUUID `db:"ws_group_id"`
		WashServerVersion     *int          `db:"ws_version"`
		WashServerDeleted     *bool         `db:"ws_deleted"`

		UserID             *string       `db:"u_id"`
		UserEmail          *string       `db:"u_email"`
		UserName           *string       `db:"u_name"`
		UserRole           *Role         `db:"u_role"`
		UserOrganizationID uuid.NullUUID `db:"u_organization_id"`
		UserVersion        *int64        `db:"u_version"`
	}

	TransactionCreate struct {
		ID             uuid.UUID     `db:"id"`
		OrganizationID uuid.UUID     `db:"organization_id"`
		GroupID        uuid.NullUUID `db:"group_id"`
		Amount         int64         `db:"amount"`
		Operation      Operation     `db:"operation"`
		CreatedAt      time.Time     `db:"created_at"`
		ForDate        *time.Time    `db:"for_date"`
		Service        Service       `db:"service"`
		StationsCount  *int          `db:"stations_count"`
		UserID         *string       `db:"user_id"`
		WashServerID   uuid.NullUUID `db:"wash_server_id"`
	}

	Operation string
	Service   string
)

const (
	DepositOperation Operation = "deposit"
	DebitOperation   Operation = "debit"

	PaymentService Service = "payment"
	BonusService   Service = "bonus"
	SbpService     Service = "sbp"
)
