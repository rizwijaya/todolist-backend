package routes

import (
	activityControllerV1 "todolist-backend/modules/v1/activities/interfaces/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRouter(router *fiber.App, db *gorm.DB) *fiber.App {
	activityControllerV1 := activityControllerV1.NewActivityController(db)

	api := router.Group("/activity-groups")
	api.Get("", activityControllerV1.GetAllActivity)
	api.Get("/:id", activityControllerV1.GetActivityByID)
	api.Post("", activityControllerV1.CreateActivity)
	api.Patch("/:id", activityControllerV1.UpdateActivity)
	api.Delete("/:id", activityControllerV1.DeleteActivity)

	return router
}
