package routes

import (
	todoControllerV1 "todolist-backend/modules/v1/todos/interfaces/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRouter(router *fiber.App, db *gorm.DB) *fiber.App {
	todoControllerV1 := todoControllerV1.NewController(db)

	api := router.Group("/api/v1")
	api.Get("/", todoControllerV1.Hello)

	return router
}
