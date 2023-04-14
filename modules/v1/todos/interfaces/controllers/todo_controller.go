package controllers

import "github.com/gofiber/fiber/v2"

func (uc *TodoController) Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
