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
	tokenTTL = time.Second * 30
)

type AuthService struct {
	auth         *db.AuthPostgres
	serviceEmail *ServiceEmail
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID uint32 `json:"user_id"`
}

func NewAuthService(postgres *db.AuthPostgres) *AuthService {
	return &AuthService{
		auth:         postgres,
		serviceEmail: InitServiceEmail(),
	}
}

func (s *AuthService) CreateUser(user *db.User) (uint32, error) {
	if err := s.auth.CheckEmailUnique(user.Email); err != nil {
		return 0, err
	}

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

func (s *AuthService) ParseToken(accessToken string) (uint32, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errInvalidSigningMethod
		}

		return []byte(signKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errInvalidClaims
	}

	return claims.UserID, nil
}

func (s *AuthService) UpdatePassword(email, newPassword string) error {
	return s.auth.UpdatePassword(email, generateHashedPassword(newPassword))
}

func generateHashedPassword(password string) string {
	hashFunc := sha1.New()
	hashFunc.Write([]byte(password))

	return fmt.Sprintf("%x", hashFunc.Sum([]byte(salt)))
}
