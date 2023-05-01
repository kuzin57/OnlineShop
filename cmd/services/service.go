package services

import "github.com/kuzin57/OnlineShop/cmd/db"

type Authorization interface {
	CreateUser(user *db.User) (uint32, error)
	GenerateToken(username, password string) (string, error)
<<<<<<< HEAD
	ParseToken(token string) (uint32, error)
	RecoverPassword(string) error
=======
	ParseToken(token string) (int, error)
>>>>>>> 35fe851 (made some changes)
}
