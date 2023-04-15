package controllers

import (
	todoRepository "todolist-backend/modules/v1/todos/interfaces/repositories"
	todoUsecase "todolist-backend/modules/v1/todos/usecases"

	"gorm.io/gorm"
)

type TodoController struct {
	todoUsecase todoUsecase.TodoAdapter
}

func NewTodoController(db *gorm.DB) *TodoController {
	//Todo
	repo := todoRepository.NewTodoRepository(db)
	tu := todoUsecase.NewTodoUsecase(repo)

	return &TodoController{
		todoUsecase: tu,
	}
}
