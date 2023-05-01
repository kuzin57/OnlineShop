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
<<<<<<< HEAD
	tokenTTL = time.Second * 30
)

type AuthService struct {
	auth         *db.AuthPostgres
	serviceEmail *ServiceEmail
=======
	tokenTTL = 12 * time.Hour
)

type AuthService struct {
	auth *db.AuthPostgres
>>>>>>> 35fe851 (made some changes)
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID uint32 `json:"user_id"`
}

func NewAuthService(postgres *db.AuthPostgres) *AuthService {
	return &AuthService{
<<<<<<< HEAD
		auth:         postgres,
		serviceEmail: InitServiceEmail(),
=======
		auth: postgres,
>>>>>>> 35fe851 (made some changes)
	}
}

func (s *AuthService) CreateUser(user *db.User) (uint32, error) {
<<<<<<< HEAD
	if err := s.auth.CheckEmailUnique(user.Email); err != nil {
		return 0, err
	}

=======
>>>>>>> 35fe851 (made some changes)
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

<<<<<<< HEAD
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

func (s *AuthService) RecoverPassword(email string) error {
	return nil
=======
func (s *AuthService) ParseToken(token string) (int, error) {
	return 0, nil
}

func generateHashedPassword(password string) string {
	hashFunc := sha1.New()
	hashFunc.Write([]byte(password))

	return fmt.Sprintf("%x", hashFunc.Sum([]byte(salt)))
>>>>>>> 35fe851 (made some changes)
}

func generateHashedPassword(password string) string {
	hashFunc := sha1.New()
	hashFunc.Write([]byte(password))

	return fmt.Sprintf("%x", hashFunc.Sum([]byte(salt)))
}
