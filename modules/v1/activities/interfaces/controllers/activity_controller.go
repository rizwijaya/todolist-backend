package controllers

import (
	"log"
	"net/http"
	"todolist-backend/modules/v1/activities/domain"
	api "todolist-backend/pkg/api_response"
	"todolist-backend/pkg/http_error"

	"github.com/go-playground/validator/v10"
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
	var (
		InsertActivity domain.InsertActivity
		Validator      = validator.New()
	)
	//Parse and validate request body
	if err := c.BodyParser(&InsertActivity); err != nil {
		log.Println(err)
		resp := api.NewErrorResponse("Bad Request", err.Error())
		return c.Status(http.StatusBadRequest).JSON(resp)
	}
	//Validate request body input
	err := Validator.Struct(InsertActivity)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			log.Println(err)
			resp := api.NewErrorResponse("Bad Request", http_error.FormValidationError(v))
			return c.Status(http.StatusBadRequest).JSON(resp)
		}
	}

	activity, err := uc.activityUsecase.CreateActivity(InsertActivity)
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
	var (
		id             = c.Params("id")
		UpdateActivity domain.UpdateActivity
	)
	//Parse and validate request body
	if err := c.BodyParser(&UpdateActivity); err != nil {
		log.Println(err)
		resp := api.NewErrorResponse("Bad Request", err.Error())
		return c.Status(http.StatusBadRequest).JSON(resp)
	}
	//Validate request body input
	if UpdateActivity.Title == "" {
		log.Println("Title is required")
		resp := api.NewErrorResponse("Bad Request", http_error.ErrTitleRequired)
		return c.Status(http.StatusBadRequest).JSON(resp)
	}

	activity, err := uc.activityUsecase.UpdateActivity(id, UpdateActivity)
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
