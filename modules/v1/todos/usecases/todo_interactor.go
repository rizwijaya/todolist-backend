package usecase

import "todolist-backend/modules/v1/todos/domain"

func (tu *TodoUsecase) GetAllTodos() ([]domain.Todos, error) {
	return tu.repoTodo.FindAll()
}

func (tu *TodoUsecase) GetTodosByGroupId(group_id string) ([]domain.Todos, error) {
	return tu.repoTodo.FindByGroupId(group_id)
}

func (tu *TodoUsecase) GetTodoById(id string) (domain.Todos, error) {
	return tu.repoTodo.FindById(id)
}
