package repository

import (
	"github.com/Doni-githu/todo-golang-app"
	"github.com/jmoiron/sqlx"
)

type Authonrization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authonrization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authonrization: NewAuthPostgres(db),
	}
}
