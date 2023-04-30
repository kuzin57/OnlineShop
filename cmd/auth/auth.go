package auth

import (
	"fmt"

	"github.com/kuzin57/OnlineShop/cmd/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uint32 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

const (
	checkUserExistenceQuery = "SELECT * FROM bshop.Users WHERE email"
)

func Login(reqUser *User) error {
	var (
		id       uint32
		name     string
		email    string
		password string
	)

	checkingQuery := fmt.Sprintf(
		"%s = %s",
		checkUserExistenceQuery,
		reqUser.Email)
	row := db.Database.QueryRow(checkingQuery)

	if err := row.Scan(&id, &name, &email, &password); err != nil {
		return errIncorrectEmail
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(password),
		[]byte(reqUser.Password)); err != nil {
		return errInvalidPassword
	}
	return nil
}
