package domain

import "time"

type GormModel struct {
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

type Activities struct {
	Activity_id int    `json:"id" gorm:"column:activity_id"`
	Title       string `json:"title" gorm:"column:title"`
	Email       string `json:"email" gorm:"column:email"`
	GormModel
}
