package app

import "errors"

var (
	ErrNotFound          = errors.New("entity not found")
	ErrForbidden         = errors.New("access denied")
	ErrBadRequest        = errors.New("no fields to update")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrBadValue          = errors.New("bad value")
	ErrAlreadyExists     = errors.New("record already exists")
	ErrEmptyUpdate       = errors.New("no fields to update")
	ErrOldVersion        = errors.New("old version")
)
