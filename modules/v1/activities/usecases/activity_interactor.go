package usecase

import (
	"todolist-backend/modules/v1/activities/domain"
	"todolist-backend/pkg/http_error"
)

func (au *ActivityUsecase) GetAllActivity() ([]domain.Activities, error) {
	return au.repoActivity.FindAll()
}

func (au *ActivityUsecase) GetActivityByID(id string) (domain.Activities, error) {
	return au.repoActivity.FindByID(id)
}

func (au *ActivityUsecase) CreateActivity(activity domain.InsertActivity) (domain.Activities, error) {
	newActivity := domain.Activities{
		Title: activity.Title,
		Email: activity.Email,
	}
	return au.repoActivity.Create(newActivity)
}

func (au *ActivityUsecase) UpdateActivity(id string, activity domain.UpdateActivity) (domain.Activities, error) {
	newActivity := domain.Activities{
		Title: activity.Title,
		Email: activity.Email,
	}
	res, err := au.repoActivity.FindByID(id)
	if err != nil {
		if http_error.IsSame(err, http_error.ErrRecordNotfound) {
			return res, http_error.ErrRecordNotfound
		}
		return res, err
	}

	act, err := au.repoActivity.Update(id, newActivity)
	if err != nil {
		return act, err
	}
	return au.repoActivity.FindByID(id)
}

func (au *ActivityUsecase) DeleteActivity(id string) error {
	_, err := au.repoActivity.FindByID(id)
	if err != nil {
		if http_error.IsSame(err, http_error.ErrRecordNotfound) {
			return http_error.ErrRecordNotfound
		}
		return err
	}

	return au.repoActivity.Delete(id)
}
