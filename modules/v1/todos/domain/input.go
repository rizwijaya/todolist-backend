package domain

type InsertTodos struct {
	Todo_id           int    `json:"id"`
	Title             string `json:"title" validate:"required"`
	Activity_group_id int    `json:"activity_group_id" validate:"required"`
	Is_active         *bool  `json:"is_active" validate:"required"`
	Priority          string `json:"priority" validate:"omitempty"`
	GormModel
}

type UpdateTodos struct {
	Todo_id           int    `json:"id"`
	Activity_group_id int    `json:"activity_group_id"`
	Title             string `json:"title"`
	Is_active         *bool  `json:"is_active"`
	Priority          string `json:"priority"`
	GormModel
}
