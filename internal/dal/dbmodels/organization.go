package dbmodels

import uuid "github.com/satori/go.uuid"

type (
	Organization struct {
		ID                            uuid.UUID `db:"id"`
		Name                          string    `db:"name"`
		DisplayName                   string    `db:"display_name"`
		Description                   string    `db:"description"`
		IsDefault                     bool      `db:"is_default"`
		ReportsProcessingDelayMinutes int64     `db:"processing_delay"`
		BonusPercentage               int64     `db:"bonus_percentage"`
		Deleted                       bool      `db:"deleted"`
	}

	OrganizationCreation struct {
		Name                          string  `db:"name"`
		DisplayName                   string  `db:"display_name"`
		Description                   string  `db:"description"`
		ReportsProcessingDelayMinutes *string `db:"processing_delay"`
		BonusPercentage               *int64  `db:"bonus_percentage"`
	}

	OrganizationUpdate struct {
		Name                          *string `db:"name"`
		DisplayName                   *string `db:"display_name"`
		Description                   *string `db:"description"`
		ReportsProcessingDelayMinutes *int64  `db:"processing_delay"`
		BonusPercentage               *int64  `db:"bonus_percentage"`
	}

	OrganizationFilter struct {
		BaseFilter
		OrganizationIDs []uuid.UUID
	}
)
