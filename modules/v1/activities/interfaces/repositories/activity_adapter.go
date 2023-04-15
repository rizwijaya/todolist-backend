package repository

import (
	"todolist-backend/modules/v1/activities/domain"

	"gorm.io/gorm"
)

type RepositoryPresenter interface {
	FindAll() ([]domain.Activities, error)
	FindByID(id string) (domain.Activities, error)
}

type ActivityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) *ActivityRepository {
	return &ActivityRepository{db}
}
