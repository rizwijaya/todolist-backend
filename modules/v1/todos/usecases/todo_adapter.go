package usecase

import repositoriesTodo "todolist-backend/modules/v1/todos/interfaces/repositories"

type UsecasePresenter interface {
}

type Usecase struct {
	repository repositoriesTodo.RepositoryPresenter
}

func NewUsecase(repositories repositoriesTodo.RepositoryPresenter) *Usecase {
	return &Usecase{repositories}
}
