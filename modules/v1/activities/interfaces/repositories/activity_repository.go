package repository

import "todolist-backend/modules/v1/activities/domain"

func (ar *ActivityRepository) FindAll() ([]domain.Activities, error) {
	var activities []domain.Activities
	err := ar.db.Find(&activities).Error
	return activities, err
}
