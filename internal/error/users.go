package error

import "errors"

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrNameNotUnique = errors.New("name not unique")
	ErrPassword      = errors.New("password error")
	ErrInvalidEmail  = errors.New("invalid email error")
)
