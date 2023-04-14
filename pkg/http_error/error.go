package http_error

import "github.com/gofiber/fiber/v2"

func PageNotFound() func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		c.JSON(fiber.Map{
			"status":  "error",
			"message": "Page not found",
		})
	}
}
