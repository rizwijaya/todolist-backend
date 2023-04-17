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

func (au *ActivityUsecase) CreateActivity(activity domain.Activities) (domain.Activities, error) {
	return au.repoActivity.Create(activity)
}

func (au *ActivityUsecase) UpdateActivity(id string, activity domain.Activities) (domain.Activities, error) {
	return au.repoActivity.Update(id, activity)
}

func (au *ActivityUsecase) DeleteActivity(id string) error {
	return au.repoActivity.Delete(id)
}
