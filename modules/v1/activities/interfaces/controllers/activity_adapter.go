package controllers

import (
	repositoryActivity "todolist-backend/modules/v1/activities/interfaces/repositories"
	activityUsecase "todolist-backend/modules/v1/activities/usecases"

	"gorm.io/gorm"
)

type ActivityController struct {
	activityUsecase activityUsecase.UsecasePresenter
}

func NewController(activityUsecase activityUsecase.UsecasePresenter) *ActivityController {
	return &ActivityController{activityUsecase}
}

func UserController(db *gorm.DB) *ActivityController {
	//Activity
	repository := repositoryActivity.NewRepository(db)
	usecase := activityUsecase.NewUsecase(repository)

	return NewController(usecase)
}
