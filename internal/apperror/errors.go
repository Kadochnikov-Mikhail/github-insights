package apperror

import (
	"errors"
)

var (
	ErrUsernameRequired = errors.New("username is required")
	ErrUserNotFound     = errors.New("user not found")
)
