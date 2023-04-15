package domain

type InsertActivity struct {
	Title string `json:"title" binding:"required"`
	Email string `json:"email" binding:"omitempty,email"`
}
type UpdateActivity struct {
	Title string `json:"title" binding:"omitempty"`
	Email string `json:"email" binding:"omitempty,email"`
}
