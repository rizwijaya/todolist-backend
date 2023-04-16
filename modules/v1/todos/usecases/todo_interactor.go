package usecase

import (
	"todolist-backend/modules/v1/todos/domain"
	"todolist-backend/pkg/http_error"
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

func (tu *TodoUsecase) CreateTodo(todo domain.InsertTodos) (domain.Todos, error) {
	newTodos := domain.Todos{
		Activity_group_id: todo.Activity_group_id,
		Title:             todo.Title,
		Is_active:         todo.Is_active,
		Priority:          todo.Priority,
	}
	return tu.repoTodo.Create(newTodos)
}

func (tu *TodoUsecase) UpdateTodo(id string, todos domain.UpdateTodos) (domain.Todos, error) {
	newTodo := domain.Todos{
		Activity_group_id: todos.Activity_group_id,
		Title:             todos.Title,
		Is_active:         todos.Is_active,
		Priority:          todos.Priority,
	}

	res, err := tu.repoTodo.FindById(id)
	if err != nil {
		if http_error.IsSame(err, http_error.ErrRecordNotfound) {
			return res, http_error.ErrRecordNotfound
		}
		return res, err
	}

	tod, err := tu.repoTodo.Update(id, newTodo)
	if err != nil {
		return tod, err
	}
	return tu.repoTodo.FindById(id)
}

func (tu *TodoUsecase) DeleteTodo(id string) error {
	_, err := tu.repoTodo.FindById(id)
	if err != nil {
		if http_error.IsSame(err, http_error.ErrRecordNotfound) {
			return http_error.ErrRecordNotfound
		}
		return err
	}

	return tu.repoTodo.Delete(id)
}
