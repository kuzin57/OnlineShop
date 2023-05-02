package auth

import (
	"errors"
	"fmt"
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

	fmt.Println("I am here, token:", r.Header[tokenHeader][0])

	token, err := jwt.ParseWithClaims(r.Header[tokenHeader][0], &tokenClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errNotAuthorized
		}

		return []byte(signKey), nil
	})

	fmt.Println("err", err)

	if err != nil {
		return errNotAuthorized
	}

	if !token.Valid {
		return errNotAuthorized
	}

	return nil
}
