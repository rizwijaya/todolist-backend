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

// @Summary Get All Activity
// @Description Get All Activity
// @Tags Activity
// @Accept json
// @Produce json
// @Success 200 {object} api.ResponseSuccess
// @Failure 500 {object} api.ResponseError
// @Router /activity-groups [get]
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

// @Summary Get Activity By ID
// @Description Get Activity By ID
// @Tags Activity
// @Accept json
// @Produce json
// @Param id path string true "Activity ID"
// @Success 200 {object} api.ResponseSuccess
// @Failure 404 {object} api.ResponseError
// @Failure 500 {object} api.ResponseError
// @Router /activity-groups/{id} [get]
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

// @Summary Create Activity
// @Description Create Activity
// @Tags Activity
// @Accept json
// @Produce json
// @Param body body domain.Activities true "Activities Request"
// @Success 201 {object} api.ResponseSuccess
// @Failure 400 {object} api.ResponseError
// @Failure 500 {object} api.ResponseError
// @Router /activity-groups [post]
func (uc *ActivityController) CreateActivity(c *fiber.Ctx) error {
	var (
		activity  domain.Activities
		Validator = validator.New()
	)
	//Parse and validate request body
	if err := c.BodyParser(&activity); err != nil {
		log.Println(err)
		resp := api.NewErrorResponse("Bad Request", err.Error())
		return c.Status(http.StatusBadRequest).JSON(resp)
	}
	//Validate request body input
	err := Validator.Struct(activity)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			log.Println(err)
			resp := api.NewErrorResponse("Bad Request", http_error.FormValidationError(v))
			return c.Status(http.StatusBadRequest).JSON(resp)
		}
	}

	activity, err = uc.activityUsecase.CreateActivity(activity)
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

// @Summary Update Activity
// @Description Update Activity
// @Tags Activity
// @Accept json
// @Produce json
// @Param id path string true "Activity ID"
// @Param body body domain.Activities true "Activities Request"
// @Success 200 {object} api.ResponseSuccess
// @Failure 400 {object} api.ResponseError
// @Failure 404 {object} api.ResponseError
// @Failure 500 {object} api.ResponseError
// @Router /activity-groups/{id} [patch]
func (uc *ActivityController) UpdateActivity(c *fiber.Ctx) error {
	var (
		id       = c.Params("id")
		activity domain.Activities
	)
	//Parse and validate request body
	if err := c.BodyParser(&activity); err != nil {
		log.Println(err)
		resp := api.NewErrorResponse("Bad Request", err.Error())
		return c.Status(http.StatusBadRequest).JSON(resp)
	}
	//Validate request body input
	if activity.Title == "" {
		log.Println("Title is required")
		resp := api.NewErrorResponse("Bad Request", http_error.ErrTitleRequired)
		return c.Status(http.StatusBadRequest).JSON(resp)
	}

	activity, err := uc.activityUsecase.UpdateActivity(id, activity)
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

// @Summary Delete Activity
// @Description Delete Activity
// @Tags Activity
// @Accept json
// @Produce json
// @Param id path string true "Activity ID"
// @Success 200 {object} api.ResponseSuccess
// @Failure 404 {object} api.ResponseError
// @Failure 500 {object} api.ResponseError
// @Router /activity-groups/{id} [delete]
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
