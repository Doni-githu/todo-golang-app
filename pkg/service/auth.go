package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/Doni-githu/todo-golang-app"
	"github.com/Doni-githu/todo-golang-app/pkg/repository"
)

const salt = "awdjkjnkhb'lalkmlklkmkbhbvjfhbjvdjlkl"

type AuthService struct {
	repo repository.Authonrization
}

func newAuthService(repo repository.Authonrization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}