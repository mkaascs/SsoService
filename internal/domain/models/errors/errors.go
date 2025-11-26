package errors

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)
