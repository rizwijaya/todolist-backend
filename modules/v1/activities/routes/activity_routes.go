package routes

import (
	activityControllerV1 "todolist-backend/modules/v1/activities/interfaces/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRouter(router *fiber.App, db *gorm.DB) *fiber.App {
	activityControllerV1 := activityControllerV1.NewController(db)

	api := router.Group("/api/v1")
	api.Get("/", activityControllerV1.Hello)

	return router
}
