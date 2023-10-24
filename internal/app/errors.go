package app

import "errors"

var (
	ErrNotFound          = errors.New("entity not found")
	ErrForbidden         = errors.New("access denied")
	ErrBadRequest        = errors.New("no fields to update")
	ErrProfileInactive   = errors.New("profile inactive")
	ErrInsufficientFunds = errors.New("insufficient funds")

	ErrBadValue      = errors.New("bad value")
	ErrAlreadyExists = errors.New("record already exists")

	ErrMessageDuplicate = errors.New("duplicate message")
	ErrNotEnoughMoney   = errors.New("not enough money")
	ErrSessionNoUser    = errors.New("session without user")

	ErrWashServerConnectionInit = errors.New("failed to init wash server connection")
	ErrWashServerNotFound       = errors.New("wash server with specified key not exists")
	ErrEmptyWashServerKey       = errors.New("wash server key must be not empty")
	ErrWashServerKeyAlreadyUsed = errors.New("wash server key already in use")
)
