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

func (tc *TodoController) CreateTodo(c *fiber.Ctx) error {
	var (
		Validator  = validator.New()
		insertTodo domain.InsertTodos
	)
	//Parse and validate request body
	err := c.BodyParser(&insertTodo)
	if err != nil {
		log.Println(err)
		resp := api.NewErrorResponse("Bad Request", err.Error())
		return c.Status(http.StatusBadRequest).JSON(resp)
	}
	//Validate request body input
	err = Validator.Struct(insertTodo)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			if http_error.FormValidationError(v) == http_error.ErrIsActiveNull {
				sts := true
				insertTodo.Is_active = &sts
			}

			if http_error.FormValidationError(v) == http_error.ErrPriorityNull {
				insertTodo.Priority = "very-high"
			} else {
				log.Println(err)
				resp := api.NewErrorResponse("Bad Request", http_error.FormValidationError(v))
				return c.Status(http.StatusBadRequest).JSON(resp)
			}
		}
	}

	todo, err := tc.todoUsecase.CreateTodo(insertTodo)
	if err != nil {
		log.Println(err)
		resp := api.NewErrorResponse("Internal Server Error", "Internal Server Error")
		return c.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp := api.NewSuccessResponse("Success", "Success", todo)
	return c.Status(http.StatusCreated).JSON(resp)
}

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
