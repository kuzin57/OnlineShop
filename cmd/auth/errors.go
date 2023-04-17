package auth

import "errors"

var (
	errInvalidPassword = errors.New("Invalid password")
	errIncorrectEmail  = errors.New("No user with this email exists")
)
