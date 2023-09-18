package entity

import (
	uuid "github.com/satori/go.uuid"
)

type (
	Organization struct {
		ID                            uuid.UUID
		Name                          string
		DisplayName                   string
		Description                   string
		IsDefault                     bool
		ReportsProcessingDelayMinutes int64
		BonusPercentage               int64
		Deleted                       bool
	}

	OrganizationCreation struct {
		Name                          string
		DisplayName                   string
		Description                   string
		ReportsProcessingDelayMinutes *int64
		BonusPercentage               *int64
	}

	OrganizationUpdate struct {
		Name                          *string
		DisplayName                   *string
		Description                   *string
		ReportsProcessingDelayMinutes *int64
		BonusPercentage               *int64
	}

	OrganizationFilter struct {
		BaseFilter
		OrganizationIDs []uuid.UUID
	}
)
