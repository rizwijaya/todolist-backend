package usecase

import (
	"todolist-backend/modules/v1/todos/domain"
	todoRepository "todolist-backend/modules/v1/todos/interfaces/repositories"
)

type TodoAdapter interface {
	GetAllTodos() ([]domain.Todos, error)
	GetTodosByGroupId(group_id string) ([]domain.Todos, error)
	GetTodoById(id string) (domain.Todos, error)
}

type TodoUsecase struct {
	repoTodo *todoRepository.TodoRepository
}

func NewTodoUsecase(repoTodo *todoRepository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{repoTodo}
}
