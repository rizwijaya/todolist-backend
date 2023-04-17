package domain

type UpdateTodos struct {
	Todo_id           int    `json:"id"`
	Activity_group_id int    `json:"activity_group_id"`
	Title             string `json:"title"`
	Is_active         *bool  `json:"is_active"`
	Priority          string `json:"priority"`
	GormModel
}
