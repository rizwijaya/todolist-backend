package usecase

import (
	"todolist-backend/modules/v1/todos/domain"
	todoRepository "todolist-backend/modules/v1/todos/interfaces/repositories"
)

type TodoAdapter interface {
	GetAllTodos() ([]domain.Todos, error)
	GetTodosByGroupId(group_id string) ([]domain.Todos, error)
	GetTodoById(id string) (domain.Todos, error)
	CreateTodo(domain.Todos) (domain.Todos, error)
	UpdateTodo(id string, todos domain.UpdateTodos) (domain.Todos, error)
	DeleteTodo(id string) error
}

type TodoUsecase struct {
	repoTodo *todoRepository.TodoRepository
}

func NewTodoUsecase(repoTodo *todoRepository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{repoTodo}
}
