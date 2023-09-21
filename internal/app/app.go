package app

import (
	"wash-payment/internal/entity"
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
		UserRepo UserRepo
	}

	Services struct {
		UserService   UserService
		RabbitService RabbitService
	}
)
