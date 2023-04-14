package controllers

import (
	repositoryTodo "todolist-backend/modules/v1/todos/interfaces/repositories"
	todoUsecase "todolist-backend/modules/v1/todos/usecases"

	"gorm.io/gorm"
)

type TodoController struct {
	todoUsecase todoUsecase.UsecasePresenter
}

func NewController(todoUsecase todoUsecase.UsecasePresenter) *TodoController {
	return &TodoController{todoUsecase}
}

func UserController(db *gorm.DB) *TodoController {
	//Todo
	repository := repositoryTodo.NewRepository(db)
	usecase := todoUsecase.NewUsecase(repository)

	return NewController(usecase)
}
