package repository

import "todolist-backend/modules/v1/todos/domain"

func (tr *TodoRepository) FindAll() ([]domain.Todos, error) {
	var todos []domain.Todos
	err := tr.db.Find(&todos).Error
	return todos, err
}

func (tr *TodoRepository) FindByGroupId(group_id string) ([]domain.Todos, error) {
	var todos []domain.Todos
	err := tr.db.Where("activity_group_id = ?", group_id).Find(&todos).Error
	return todos, err
}

func (tr *TodoRepository) FindById(id string) (domain.Todos, error) {
	var todo domain.Todos
	err := tr.db.Where("todo_id = ?", id).First(&todo).Error
	return todo, err
}
