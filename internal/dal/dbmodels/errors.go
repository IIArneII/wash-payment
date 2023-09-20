package dbmodels

import (
	"errors"

	"github.com/lib/pq"
)

var (
	ErrNotFound      = errors.New("entity not found")
	ErrAlreadyExists = errors.New("record already exists")
)

var (
	PQErrAlreadyExists pq.ErrorCode = "23505"
)
