package auth

import "errors"

var (
	errInvalidSigningMethod = errors.New("invalid signing method")
	errInvalidClaims        = errors.New("claims are not of type *tokenClaims")
)
