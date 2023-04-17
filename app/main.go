package main

import (
	"todolist-backend/infrastructures/config"
	database "todolist-backend/infrastructures/databases"
	routesActivityAPIV1 "todolist-backend/modules/v1/activities/routes"
	routesTodosAPIV1 "todolist-backend/modules/v1/todos/routes"

	_ "todolist-backend/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"gorm.io/gorm"
)

// @title TodoList API Documentation
// @description This is a sample server TodoList server.
// @version 1.0.0
// @termsOfService http://swagger.io/terms/
// @contact.name Rizwijaya
// @contact.email admin@rizwijaya.com
// @license.name MIT
// @license.url http://opensource.org/licenses/MIT
// @host localhost:3030
// @BasePath /

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
	router.Get("/swagger/*", swagger.HandlerDefault)
	if config.App.Port == "" {
		config.App.Port = "3030"
	}
	router.Listen(":" + config.App.Port)
}
