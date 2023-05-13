package auth

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

const (
	tokenHeader = "Token"
)

var (
	errNotAuthorized = errors.New("user not authorized")
)

func CheckAuthorized(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	if r.Header[tokenHeader] == nil {
		return errNotAuthorized
	}

	token, err := jwt.ParseWithClaims(
		r.Header[tokenHeader][0],
		&tokenClaims{},
		func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errNotAuthorized
			}

			return []byte(signKey), nil
		})

	if err != nil {
		return errNotAuthorized
	}

	if !token.Valid {
		return errNotAuthorized
	}

	return nil
}
