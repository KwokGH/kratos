package entity

import "errors"

var (
	ErrRecordNotFound  = errors.New("record not found")
	ErrDuplicateRecord = errors.New("duplicate record")
	ErrUnAuthorized    = errors.New("unauthorized")
	ErrInvalidArgs     = errors.New("invalid args")
)
