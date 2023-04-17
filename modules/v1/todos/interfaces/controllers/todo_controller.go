package controllers

import (
	"log"
	"net/http"
	"todolist-backend/modules/v1/todos/domain"
	api "todolist-backend/pkg/api_response"
	"todolist-backend/pkg/http_error"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get All Todos
// @Description Get All Todos
// @Tags Todo
// @Accept json
// @Produce json
// @Param activity_group_id query string false "Activity Group ID"
// @Success 200 {object} api.ResponseSuccess
// @Failure 500 {object} api.ResponseError
// @Router /todo-items [get]
func (tc *TodoController) GetAllTodos(c *fiber.Ctx) error {
	var (
		group_id = c.Query("activity_group_id")
		err      error
		todos    []domain.Todos
	)
	if group_id == "" {
		todos, err = tc.todoUsecase.GetAllTodos()
	} else {
		todos, err = tc.todoUsecase.GetTodosByGroupId(group_id)
	}
	if err != nil {
		log.Println(err)
		resp := api.NewErrorResponse("Internal Server Error", "Internal Server Error")
		return c.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp := api.NewSuccessResponse("Success", "Success", todos)
	return c.Status(http.StatusOK).JSON(resp)
}

// @Summary Get Todo By ID
// @Description Get Todo By ID
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} api.ResponseSuccess
// @Failure 404 {object} api.ResponseError
// @Failure 500 {object} api.ResponseError
// @Router /todo-items/{id} [get]
func (tc *TodoController) GetTodoById(c *fiber.Ctx) error {
	id := c.Params("id")
	todo, err := tc.todoUsecase.GetTodoById(id)
	if err != nil {
		log.Println(err)
		if http_error.IsSame(err, http_error.ErrRecordNotfound) {
			resp := api.NewErrorResponse("Not Found", "Todo with ID "+id+" Not Found")
			return c.Status(http.StatusNotFound).JSON(resp)
		}

		resp := api.NewErrorResponse("Internal Server Error", "Internal Server Error")
		return c.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp := api.NewSuccessResponse("Success", "Success", todo)
	return c.Status(http.StatusOK).JSON(resp)
}

// @Summary Create Todo
// @Description Create Todo
// @Tags Todo
// @Accept json
// @Produce json
// @Param body body domain.Todos true "Todo Request"
// @Success 201 {object} api.ResponseSuccess
// @Failure 400 {object} api.ResponseError
// @Failure 500 {object} api.ResponseError
// @Router /todo-items [post]
func (tc *TodoController) CreateTodo(c *fiber.Ctx) error {
	var (
		Validator = validator.New()
		todo      domain.Todos
	)
	//Parse and validate request body
	err := c.BodyParser(&todo)
	if err != nil {
		log.Println(err)
		resp := api.NewErrorResponse("Bad Request", err.Error())
		return c.Status(http.StatusBadRequest).JSON(resp)
	}
	//Validate request body input
	err = Validator.Struct(todo)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			if http_error.FormValidationError(v) == http_error.ErrIsActiveNull {
				sts := true
				todo.Is_active = &sts
			} else if http_error.FormValidationError(v) == http_error.ErrPriorityNull {
				todo.Priority = "very-high"
			} else {
				log.Println(err)
				resp := api.NewErrorResponse("Bad Request", http_error.FormValidationError(v))
				return c.Status(http.StatusBadRequest).JSON(resp)
			}
		}
	}

	todo, err = tc.todoUsecase.CreateTodo(todo)
	if err != nil {
		log.Println(err)
		resp := api.NewErrorResponse("Internal Server Error", "Internal Server Error")
		return c.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp := api.NewSuccessResponse("Success", "Success", todo)
	return c.Status(http.StatusCreated).JSON(resp)
}

// @Summary Update Todo
// @Description Update Todo
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Param body body domain.UpdateTodos true "Todo Body"
// @Success 200 {object} api.ResponseSuccess
// @Failure 400 {object} api.ResponseError
// @Failure 404 {object} api.ResponseError
// @Failure 500 {object} api.ResponseError
// @Router /todo-items/{id} [patch]
func (tc *TodoController) UpdateTodo(c *fiber.Ctx) error {
	var Updatetodo domain.UpdateTodos
	id := c.Params("id")
	err := c.BodyParser(&Updatetodo)
	if err != nil {
		log.Println(err)
		resp := api.NewErrorResponse("Bad Request", err.Error())
		return c.Status(http.StatusBadRequest).JSON(resp)
	}

	todo, err := tc.todoUsecase.UpdateTodo(id, Updatetodo)
	if err != nil {
		log.Println(err)
		if http_error.IsSame(err, http_error.ErrRecordNotfound) {
			resp := api.NewErrorResponse("Not Found", "Todo with ID "+id+" Not Found")
			return c.Status(http.StatusNotFound).JSON(resp)
		}

		resp := api.NewErrorResponse("Internal Server Error", "Internal Server Error")
		return c.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp := api.NewSuccessResponse("Success", "Success", todo)
	return c.Status(http.StatusOK).JSON(resp)
}

// @Summary Delete Todo
// @Description Delete Todo
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} api.ResponseSuccess
// @Failure 404 {object} api.ResponseError
// @Failure 500 {object} api.ResponseError
// @Router /todo-items/{id} [delete]
func (tc *TodoController) DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	err := tc.todoUsecase.DeleteTodo(id)
	if err != nil {
		log.Println(err)
		if http_error.IsSame(err, http_error.ErrRecordNotfound) {
			resp := api.NewErrorResponse("Not Found", "Todo with ID "+id+" Not Found")
			return c.Status(http.StatusNotFound).JSON(resp)
		}

		resp := api.NewErrorResponse("Internal Server Error", "Internal Server Error")
		return c.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp := api.NewSuccessResponse("Success", "Success", nil)
	return c.Status(http.StatusOK).JSON(resp)
}
