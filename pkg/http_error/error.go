package http_error

import "github.com/gofiber/fiber/v2"

func IsSame(err error, target error) bool {
	return err.Error() == target.Error()
}

func PageNotFound() func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		c.JSON(fiber.Map{
			"status":  "error",
			"message": "Page not found",
		})
	}
}
