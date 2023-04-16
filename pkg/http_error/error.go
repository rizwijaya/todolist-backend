package http_error

import (
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
	switch fe.Tag() {
	case "required":
		return fe.Field() + " cannot be null"
	case "email":
		return fe.Field() + " must be a valid email address!"
	case "min":
		return fe.Field() + " minimum " + fe.Param() + " characters!"
	case "max":
		return fe.Field() + " maximum " + fe.Param() + " characters!"
	case "alphanum":
		return fe.Field() + " must be alphanumeric!"
	case "numeric":
		return fe.Field() + " must be numeric!"
	case "eqfield":
		return fe.Field() + " must be equal to " + fe.Param() + "!"
	case "alphanumunicode":
		return fe.Field() + " must be alphanumeric unicode!"
	default:
		return fe.Field() + " is invalid!"
	}
}
