package routes

import (
	todoControllerV1 "todolist-backend/modules/v1/todos/interfaces/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRouter(router *fiber.App, db *gorm.DB) *fiber.App {
	todoControllerV1 := todoControllerV1.NewTodoController(db)

	api := router.Group("/todo-items")
	api.Get("", todoControllerV1.GetAllTodos)
	api.Get("/:id", todoControllerV1.GetTodoById)
	api.Post("", todoControllerV1.CreateTodo)
	api.Patch("/:id", todoControllerV1.UpdateTodo)
	api.Delete("/:id", todoControllerV1.DeleteTodo)

	return router
}
