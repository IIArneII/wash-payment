package dbmodels

import (
	"errors"

	"github.com/lib/pq"
)

var (
	ErrNotFound          = errors.New("entity not found")
	ErrAlreadyExists     = errors.New("record already exists")
	ErrEmptyUpdate       = errors.New("no fields to update")
	ErrInsufficientFunds = errors.New("insufficient funds")
)

var (
	PQErrAlreadyExists pq.ErrorCode = "23505"
)
