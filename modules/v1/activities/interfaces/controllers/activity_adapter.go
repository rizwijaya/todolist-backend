package controllers

import (
	repositoryActivity "todolist-backend/modules/v1/activities/interfaces/repositories"
	activityUsecase "todolist-backend/modules/v1/activities/usecases"

	"gorm.io/gorm"
)

type ActivityController struct {
	activityUsecase *activityUsecase.ActivityUsecase
}

func NewActivityController(db *gorm.DB) *ActivityController {
	//Activity
	repo := repositoryActivity.NewActivityRepository(db)
	au := activityUsecase.NewActivityUsecase(repo)

	return &ActivityController{
		activityUsecase: au,
	}
}
