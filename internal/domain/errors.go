package domain

import "errors"

var (
	ErrAccountNotFound     = errors.New("account not found")
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrDuplicateAPIKey     = errors.New("duplicate api key")
	ErrInvoiceNotFound     = errors.New("invoice not found")
	ErrUnauthorized        = errors.New("unauthorized access")
)
