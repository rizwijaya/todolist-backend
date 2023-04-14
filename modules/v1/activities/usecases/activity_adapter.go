package usecase

import repositoriesActivity "todolist-backend/modules/v1/activities/interfaces/repositories"

type UsecasePresenter interface {
}

type Usecase struct {
	repository repositoriesActivity.RepositoryPresenter
}

func NewUsecase(repositories repositoriesActivity.RepositoryPresenter) *Usecase {
	return &Usecase{repositories}
}
