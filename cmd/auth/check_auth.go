package auth

import (
	"errors"
<<<<<<< HEAD
=======
	"fmt"
>>>>>>> 573a019 (finished with authorization, started with password recovery)
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

<<<<<<< HEAD
=======
	fmt.Println("I am here, token:", r.Header[tokenHeader][0])

>>>>>>> 573a019 (finished with authorization, started with password recovery)
	token, err := jwt.ParseWithClaims(r.Header[tokenHeader][0], &tokenClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errNotAuthorized
		}

		return []byte(signKey), nil
	})

<<<<<<< HEAD
=======
	fmt.Println("err", err)

>>>>>>> 573a019 (finished with authorization, started with password recovery)
	if err != nil {
		return errNotAuthorized
	}

	if !token.Valid {
		return errNotAuthorized
	}

	return nil
}
