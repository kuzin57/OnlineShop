package services

import "github.com/kuzin57/OnlineShop/pkg/db"

type Authorization interface {
	CreateUser(user *db.User) (uint32, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (uint32, error)
	UpdatePassword(email, newPassword string) error
}
