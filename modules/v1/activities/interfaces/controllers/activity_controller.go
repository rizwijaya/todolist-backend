package controllers

import (
	"log"
	"net/http"
	api "todolist-backend/pkg/api_response"

	"github.com/gofiber/fiber/v2"
)

func (uc *ActivityController) GetAllActivity(c *fiber.Ctx) error {
	activity, err := uc.activityUsecase.GetAllActivity()
	if err != nil {
		log.Println(err)
		resp := api.NewErrorResponse("Internal Server Error", "Internal Server Error")
		return c.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp := api.NewSuccessResponse("Success", "Success", activity)
	return c.Status(http.StatusOK).JSON(resp)
}
