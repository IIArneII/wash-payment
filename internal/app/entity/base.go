package entity

type (
	Auth struct {
		User         User
		Disabled     bool
		UserMetadata AuthUserMeta
	}

	AuthUserMeta struct {
		CreationTimestamp    int64
		LastLogInTimestamp   int64
		LastRefreshTimestamp int64
	}

	Filter struct {
		Page     int
		PageSize int
	}

	Page[T any] struct {
		Items      []T
		Page       int
		PageSize   int
		TotalPages int
		TotalItems int
	}
)
