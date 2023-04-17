package domain

import "time"

type GormModel struct {
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

type Activities struct {
	ID    int    `json:"id" gorm:"column:activity_id;primaryKey"`
	Title string `json:"title" gorm:"column:title" validate:"required"`
	Email string `json:"email" gorm:"column:email"`
	GormModel
}

type ActivityWithoutEmail struct {
	ID    int    `json:"id" gorm:"column:activity_id"`
	Title string `json:"title" gorm:"column:title"`
	GormModel
}
