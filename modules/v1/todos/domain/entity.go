package domain

import "time"

type GormModel struct {
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

type Todos struct {
	ID                int    `json:"id" gorm:"column:todo_id;primaryKey"`
	Activity_group_id int    `json:"activity_group_id" gorm:"column:activity_group_id"  validate:"required"`
	Title             string `json:"title" gorm:"column:title" validate:"required"`
	Is_active         *bool  `json:"is_active" gorm:"column:is_active" validate:"required"`
	Priority          string `json:"priority" gorm:"column:priority;default:'very-high'" validate:"omitempty"`
	GormModel
}
