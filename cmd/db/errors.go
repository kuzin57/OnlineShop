package db

import "errors"

var (
	errDBUnavailable     = errors.New("database id unavailable")
	errNoSuchUser        = errors.New("no such user")
	errIncorrectPassword = errors.New("incorrect password")
<<<<<<< HEAD
<<<<<<< HEAD
	errEmailExists       = errors.New("email already exists")
=======
>>>>>>> 35fe851 (made some changes)
=======
	errEmailExists       = errors.New("email already exists")
>>>>>>> 573a019 (finished with authorization, started with password recovery)
)
