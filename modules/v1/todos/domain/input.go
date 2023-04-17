package domain

type InsertTodos struct {
	Activity_group_id int    `json:"activity_group_id"`
	Title             string `json:"title"`
	Is_active         *bool  `json:"is_active"`
	Priority          string `json:"priority"`
}
type UpdateTodos struct {
	Activity_group_id int    `json:"activity_group_id"`
	Title             string `json:"title"`
	Is_active         *bool  `json:"is_active"`
	Priority          string `json:"priority"`
}
