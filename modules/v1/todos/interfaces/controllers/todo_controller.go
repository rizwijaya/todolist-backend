package controllers

import (
	"log"
	"net/http"
	"todolist-backend/modules/v1/todos/domain"
	api "todolist-backend/pkg/api_response"

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
