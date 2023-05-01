package auth

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kuzin57/OnlineShop/cmd/db"
)

const (
	salt     = "vndfkjnkvj938958*&^*&*"
	signKey  = "w87r8fyschcjdh*&^*&^*&hbj"
	tokenTTL = 12 * time.Hour
)

type AuthService struct {
	auth *db.AuthPostgres
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID uint32 `json:"user_id"`
}

func NewAuthService(postgres *db.AuthPostgres) *AuthService {
	return &AuthService{
		auth: postgres,
	}
}

func (s *AuthService) CreateUser(user *db.User) (uint32, error) {
	user.Password = generateHashedPassword(user.Password)
	return s.auth.CreateUser(user)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.auth.GetUser(email, generateHashedPassword(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: user.Id,
	})

	return token.SignedString([]byte(signKey))
}

func (s *AuthService) ParseToken(token string) (int, error) {
	return 0, nil
}

func generateHashedPassword(password string) string {
	hashFunc := sha1.New()
	hashFunc.Write([]byte(password))

	return fmt.Sprintf("%x", hashFunc.Sum([]byte(salt)))
}
