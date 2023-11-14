package app

import (
	"wash-payment/internal/app/entity"
)

type (
	Auth struct {
		User         entity.User
		Disabled     bool
		UserMetadata AuthUserMeta
	}

	AuthUserMeta struct {
		CreationTimestamp    int64
		LastLogInTimestamp   int64
		LastRefreshTimestamp int64
	}

	Repositories struct {
		UserRepo         UserRepo
		OrganizationRepo OrganizationRepo
		GroupRepo        GroupRepo
		//NEW
		TransactionRepo TransactionRepo
	}

	Services struct {
		UserService         UserService
		OrganizationService OrganizationService
		GroupService        GroupService
	}
)
