package repository

import (
	"todolist-backend/modules/v1/todos/domain"

	"gorm.io/gorm"
)

type RepositoryPresenter interface {
	FindAll() ([]domain.Todos, error)
	FindByGroupId(group_id string) ([]domain.Todos, error)
	FindById(id string) (domain.Todos, error)
	Create(todos domain.Todos) (domain.Todos, error)
	Update(id string, todos domain.Todos) (domain.Todos, error)
	Delete(id string) error
}

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db}
}
