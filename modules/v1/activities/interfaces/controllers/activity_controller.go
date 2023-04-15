package controllers

import (
	"log"
	"net/http"
	api "todolist-backend/pkg/api_response"
	"todolist-backend/pkg/http_error"

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

func (uc *ActivityController) GetActivityByID(c *fiber.Ctx) error {
	id := c.Params("id")
	activity, err := uc.activityUsecase.GetActivityByID(id)
	if err != nil {
		log.Println(err)
		if http_error.IsSame(err, http_error.ErrRecordNotfound) {
			resp := api.NewErrorResponse("Not Found", "Activity with ID "+id+" Not Found")
			return c.Status(http.StatusNotFound).JSON(resp)
		}
		resp := api.NewErrorResponse("Internal Server Error", "Internal Server Error")
		return c.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp := api.NewSuccessResponse("Success", "Success", activity)
	return c.Status(http.StatusOK).JSON(resp)
}
