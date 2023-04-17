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

func (tr *TodoRepository) Create(todos domain.Todos) (domain.Todos, error) {
	err := tr.db.Create(&todos).Error
	return todos, err
}

func (tr *TodoRepository) Update(id string, todos domain.Todos) (domain.Todos, error) {
	err := tr.db.Model(&todos).Where("todo_id = ?", id).Updates(&todos).First(&todos).Error
	return todos, err
}

func (tr *TodoRepository) Delete(id string) error {
	return tr.db.Where("todo_id = ?", id).First(&domain.Todos{}).Delete(&domain.Todos{}).Error
}
