package db

import "errors"

var (
	errDBUnavailable     = errors.New("database id unavailable")
	errNoSuchUser        = errors.New("no such user")
	errIncorrectPassword = errors.New("incorrect password")
	errEmailExists       = errors.New("email already exists")
	errNoSuchEmail       = errors.New("no user with such email")
)
