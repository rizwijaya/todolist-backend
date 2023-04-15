package usecase

import (
	"todolist-backend/modules/v1/activities/domain"
)

func (au *ActivityUsecase) GetAllActivity() ([]domain.Activities, error) {
	return au.repoActivity.FindAll()
}

func (au *ActivityUsecase) GetActivityByID(id string) (domain.Activities, error) {
	return au.repoActivity.FindByID(id)
}
