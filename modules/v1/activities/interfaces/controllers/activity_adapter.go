package controllers

import (
	activityRepository "todolist-backend/modules/v1/activities/interfaces/repositories"
	activityUsecase "todolist-backend/modules/v1/activities/usecases"

	"gorm.io/gorm"
)

type ActivityController struct {
	activityUsecase activityUsecase.ActivityAdapter
}

func NewActivityController(db *gorm.DB) *ActivityController {
	//Activity
	repo := activityRepository.NewActivityRepository(db)
	au := activityUsecase.NewActivityUsecase(repo)

	return &ActivityController{
		activityUsecase: au,
	}
}
