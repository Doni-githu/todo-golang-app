package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/Doni-githu/todo-golang-app"
	"github.com/Doni-githu/todo-golang-app/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	signingKey = "awd35135161adljn1@#@1gfdnkjsadaw"
	tokenTTL   = 12 * time.Hour
	salt       = "awdjkjnkhb'lalkmlklkmkbhbvjfhbjvdjlkl"
)

type AuthService struct {
	repo repository.Authonrization
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func newAuthService(repo repository.Authonrization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}


