package controllers

import (
	"fmt"
	"log"
	"net/http"
	"todolist-backend/modules/v1/activities/domain"
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

func (uc *ActivityController) CreateActivity(c *fiber.Ctx) error {
	var activityRequest domain.InsertActivity
	if err := c.BodyParser(&activityRequest); err != nil {
		log.Println(err)
		resp := api.NewErrorResponse("Bad Request", err.Error())
		return c.Status(http.StatusBadRequest).JSON(resp)
	}
	//Check title is empty
	if activityRequest.Title == "" {
		resp := api.NewErrorResponse("Bad Request", "title cannot be null")
		return c.Status(http.StatusBadRequest).JSON(resp)
	}

	activity, err := uc.activityUsecase.CreateActivity(activityRequest)
	if err != nil {
		log.Println(err)
		resp := api.NewErrorResponse("Internal Server Error", "Internal Server Error")
		return c.Status(http.StatusInternalServerError).JSON(resp)
	}
	if activity.Email == "" {
		activ := domain.ActivityWithoutEmail{}
		activ.ID = activity.ID
		activ.Title = activity.Title
		activ.CreatedAt = activity.CreatedAt
		activ.UpdatedAt = activity.UpdatedAt
		resp := api.NewSuccessResponse("Success", "Success", activ)
		return c.Status(http.StatusCreated).JSON(resp)
	}
	resp := api.NewSuccessResponse("Success", "Success", activity)
	return c.Status(http.StatusCreated).JSON(resp)
}

func (uc *ActivityController) UpdateActivity(c *fiber.Ctx) error {
	id := c.Params("id")
	var UpdateActivity domain.UpdateActivity
	if err := c.BodyParser(&UpdateActivity); err != nil {
		log.Println(err)
		resp := api.NewErrorResponse("Bad Request", err.Error())
		return c.Status(http.StatusBadRequest).JSON(resp)
	}
	//Check title is empty
	if UpdateActivity.Title == "" {
		resp := api.NewErrorResponse("Bad Request", "title cannot be null")
		return c.Status(http.StatusBadRequest).JSON(resp)
	}

	activity, err := uc.activityUsecase.UpdateActivity(id, UpdateActivity)
	fmt.Println(activity)
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

func (uc *ActivityController) DeleteActivity(c *fiber.Ctx) error {
	id := c.Params("id")
	err := uc.activityUsecase.DeleteActivity(id)
	if err != nil {
		log.Println(err)
		if http_error.IsSame(err, http_error.ErrRecordNotfound) {
			resp := api.NewErrorResponse("Not Found", "Activity with ID "+id+" Not Found")
			return c.Status(http.StatusNotFound).JSON(resp)
		}
		resp := api.NewErrorResponse("Internal Server Error", "Internal Server Error")
		return c.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp := api.NewSuccessResponse("Success", "Success", nil)
	return c.Status(http.StatusOK).JSON(resp)
}
