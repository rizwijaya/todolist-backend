package domain

import "time"

type GormModel struct {
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

type Todos struct {
	Todo_id           int    `json:"id" gorm:"column:todo_id"`
	Activity_group_id int    `json:"activity_group_id" gorm:"column:activity_group_id"`
	Title             string `json:"title" gorm:"column:title"`
	Is_active         bool   `json:"is_active" gorm:"column:is_active"`
	Priority          string `json:"priority" gorm:"column:priority"`
	GormModel
}
