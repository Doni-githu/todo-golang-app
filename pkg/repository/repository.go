package repository

import "github.com/jmoiron/sqlx"

type Authonrization interface {
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
	return &Repository{}
}
