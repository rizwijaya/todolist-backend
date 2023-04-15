package usecase

import (
	"todolist-backend/modules/v1/activities/domain"
	activityRepository "todolist-backend/modules/v1/activities/interfaces/repositories"
)

type ActivityAdapter interface {
	GetAllActivity() ([]domain.Activities, error)
}

type ActivityUsecase struct {
	repoActivity *activityRepository.ActivityRepository
}

func NewActivityUsecase(repoActivity *activityRepository.ActivityRepository) *ActivityUsecase {
	return &ActivityUsecase{repoActivity}
}
