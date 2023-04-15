package repository

import (
	"todolist-backend/modules/v1/activities/domain"

	"gorm.io/gorm"
)

type RepositoryPresenter interface {
	FindAll() ([]domain.Activities, error)
	FindByID(id string) (domain.Activities, error)
	Create(activity domain.Activities) (domain.Activities, error)
	Update(id string, activity domain.Activities) (domain.Activities, error)
	Delete(id string) error
}

type ActivityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) *ActivityRepository {
	return &ActivityRepository{db}
}
