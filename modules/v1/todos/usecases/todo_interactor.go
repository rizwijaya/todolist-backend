package usecase

import (
	"todolist-backend/modules/v1/todos/domain"
)

func (tu *TodoUsecase) GetAllTodos() ([]domain.Todos, error) {
	return tu.repoTodo.FindAll()
}

func (tu *TodoUsecase) GetTodosByGroupId(group_id string) ([]domain.Todos, error) {
	return tu.repoTodo.FindByGroupId(group_id)
}

func (tu *TodoUsecase) GetTodoById(id string) (domain.Todos, error) {
	return tu.repoTodo.FindById(id)
}

func (tu *TodoUsecase) CreateTodo(todo domain.Todos) (domain.Todos, error) {
	return tu.repoTodo.Create(todo)
}

func (tu *TodoUsecase) UpdateTodo(id string, todos domain.UpdateTodos) (domain.Todos, error) {
	newTodo := domain.Todos{
		Activity_group_id: todos.Activity_group_id,
		Title:             todos.Title,
		Is_active:         todos.Is_active,
		Priority:          todos.Priority,
	}

	return tu.repoTodo.Update(id, newTodo)
}

func (tu *TodoUsecase) DeleteTodo(id string) error {
	return tu.repoTodo.Delete(id)
}
