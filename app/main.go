package main

import (
	"log"
	"todolist-backend/infrastructures/config"
	database "todolist-backend/infrastructures/databases"
	routesActivityAPIV1 "todolist-backend/modules/v1/activities/routes"
	routesTodosAPIV1 "todolist-backend/modules/v1/todos/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func NewRouting() (*fiber.App, *gorm.DB) {
	database := database.NewDatabase()
	router := fiber.New()
	router.Use(cors.New())

	return router, database
}

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	router, db := NewRouting()

	if config.App.Mode == "development" {
		router.Use(func(c *fiber.Ctx) error {
			log.Println(c.Method(), c.OriginalURL())
			return c.Next()
		})
	}

	//Routing
	router = routesActivityAPIV1.NewRouter(router, db)
	router = routesTodosAPIV1.NewRouter(router, db)

	router.Listen(config.App.Url + ":" + config.App.Port)
}
