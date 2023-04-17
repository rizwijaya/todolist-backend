package domain

// type InsertActivity struct {
// 	Title string `json:"title" validate:"required"`
// 	Email string `json:"email" validate:"omitempty"`
// }
type UpdateActivity struct {
	Title string `json:"title" validate:"required"`
	Email string `json:"email" validate:"omitempty"`
}
