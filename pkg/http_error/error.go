package http_error

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

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

func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func FormValidationError(fe validator.FieldError) string {
	field := strings.ToLower(fe.Field())
	switch fe.Tag() {
	case "required":
		return field + " cannot be null"
	case "email":
		return field + " must be a valid email address!"
	case "min":
		return field + " minimum " + fe.Param() + " characters!"
	case "max":
		return field + " maximum " + fe.Param() + " characters!"
	case "alphanum":
		return field + " must be alphanumeric!"
	case "numeric":
		return field + " must be numeric!"
	case "eqfield":
		return field + " must be equal to " + fe.Param() + "!"
	case "alphanumunicode":
		return field + " must be alphanumeric unicode!"
	default:
		return field + " is invalid!"
	}
}
