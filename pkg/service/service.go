package service

import "github.com/Doni-githu/todo-golang-app/pkg/repository"

type Authonrization interface {
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
	return &Service{}
}
