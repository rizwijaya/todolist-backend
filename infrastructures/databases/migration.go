package database

import "time"

type GormModel struct {
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;not null"`
}
type Activities struct {
	Activity_id int    `gorm:"column:activity_id;type:bigint(20);primaryKey;autoIncrement;not null"`
	Title       string `gorm:"column:title;type:varchar(255);not null"`
	Email       string `gorm:"column:email;type:varchar(255);not null"`
	GormModel
	Todos []Todos `gorm:"foreignKey:activity_group_id;references:activity_id"`
}

type Todos struct {
	Todo_id           int    `gorm:"column:todo_id;type:bigint(20);primaryKey;autoIncrement;not null"`
	Activity_group_id int    `gorm:"column:activity_group_id;type:bigint(20);not null"`
	Title             string `gorm:"column:title;type:varchar(255);not null"`
	Priority          string `gorm:"column:priority;type:varchar(255);not null"`
	Is_active         bool   `gorm:"column:is_active;type:tinyint(1);default:true;not null"`
	GormModel
}
