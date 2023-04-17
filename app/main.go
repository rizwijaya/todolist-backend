package main

import (
	"todolist-backend/infrastructures/config"
	database "todolist-backend/infrastructures/databases"
	routesActivityAPIV1 "todolist-backend/modules/v1/activities/routes"
	routesTodosAPIV1 "todolist-backend/modules/v1/todos/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func NewRouting() (*fiber.App, *gorm.DB) {
	database := database.NewDatabase()
	router := fiber.New()
	router.Use(cors.New())

	return router, database
}

func main() {
	config := config.New()
	router, db := NewRouting()

	if config.App.Mode == "development" {
		router.Use(logger.New())
	}

	//Routing
	router = routesActivityAPIV1.NewRouter(router, db)
	router = routesTodosAPIV1.NewRouter(router, db)
	if config.App.Port == "" {
		config.App.Port = "3030"
	}
	router.Listen(":" + config.App.Port)
}
