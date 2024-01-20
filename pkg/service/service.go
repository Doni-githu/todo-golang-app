package service

import (
	"github.com/Doni-githu/todo-golang-app"
	"github.com/Doni-githu/todo-golang-app/pkg/repository"
)

type Authonrization interface {
	CreateUser(user todo.User) (int, error) 
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authonrization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authonrization: newAuthService(repos.Authonrization),
	}
}
