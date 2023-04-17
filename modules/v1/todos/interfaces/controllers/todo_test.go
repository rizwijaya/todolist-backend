package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
	"todolist-backend/modules/v1/todos/domain"
	m_usecaseTodos "todolist-backend/modules/v1/todos/usecases/mock"
	api "todolist-backend/pkg/api_response"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestActivityController_GetAllTodos(t *testing.T) {
	ctrl := gomock.NewController(t)
	now := time.Date(2023, time.April, 15, 10, 0, 0, 0, time.UTC)
	type tests struct {
		nameTest          string
		statusCode        int
		activity_group_id int
		response          api.ResponseSuccess
		wantErr           bool
		err               api.ResponseError
		usecaseTest       func(usecase *m_usecaseTodos.MockTodoAdapter)
	}
	//add test case
	test_cases := []tests{
		{
			nameTest:   "Test Case 1 Get All Todos: Success",
			statusCode: http.StatusOK,
			response: api.ResponseSuccess{
				Status:  "Success",
				Message: "Success",
				Data: []interface{}{
					map[string]interface{}{
						"id":                float64(1),
						"activity_group_id": float64(1),
						"title":             "Todo 1",
						"is_active":         false,
						"priority":          "very-high",
						"createdAt":         "2023-04-15T10:00:00Z",
						"updatedAt":         "2023-04-15T10:00:00Z",
					},
					map[string]interface{}{
						"id":                float64(2),
						"activity_group_id": float64(1),
						"title":             "Todo 2",
						"is_active":         true,
						"priority":          "high",
						"createdAt":         "2023-04-15T10:00:00Z",
						"updatedAt":         "2023-04-15T10:00:00Z",
					},
					map[string]interface{}{
						"id":                float64(3),
						"activity_group_id": float64(2),
						"title":             "Todo 3",
						"is_active":         false,
						"priority":          "low",
						"createdAt":         "2023-04-15T10:00:00Z",
						"updatedAt":         "2023-04-15T10:00:00Z",
					},
				},
			},
			wantErr: false,
			usecaseTest: func(usecase *m_usecaseTodos.MockTodoAdapter) {
				sts_false := false
				sts_true := true
				usecase.EXPECT().GetAllTodos().Return([]domain.Todos{
					{
						ID:                1,
						Activity_group_id: 1,
						Title:             "Todo 1",
						Is_active:         &sts_false,
						Priority:          "very-high",
						GormModel: domain.GormModel{
							CreatedAt: &now,
							UpdatedAt: &now,
						},
					},
					{
						ID:                2,
						Activity_group_id: 1,
						Title:             "Todo 2",
						Is_active:         &sts_true,
						Priority:          "high",
						GormModel: domain.GormModel{
							CreatedAt: &now,
							UpdatedAt: &now,
						},
					},
					{
						ID:                3,
						Activity_group_id: 2,
						Title:             "Todo 3",
						Is_active:         &sts_false,
						Priority:          "low",
						GormModel: domain.GormModel{
							CreatedAt: &now,
							UpdatedAt: &now,
						},
					},
				}, nil)
			},
		},
		{
			nameTest:          "Test Case 2 Get All Todos: Success By Activity Group Id",
			statusCode:        http.StatusOK,
			activity_group_id: 1,
			response: api.ResponseSuccess{
				Status:  "Success",
				Message: "Success",
				Data: []interface{}{
					map[string]interface{}{
						"id":                float64(1),
						"activity_group_id": float64(1),
						"title":             "Todo 1",
						"is_active":         false,
						"priority":          "very-high",
						"createdAt":         "2023-04-15T10:00:00Z",
						"updatedAt":         "2023-04-15T10:00:00Z",
					},
					map[string]interface{}{
						"id":                float64(2),
						"activity_group_id": float64(1),
						"title":             "Todo 2",
						"is_active":         true,
						"priority":          "high",
						"createdAt":         "2023-04-15T10:00:00Z",
						"updatedAt":         "2023-04-15T10:00:00Z",
					},
				},
			},
			wantErr: false,
			usecaseTest: func(usecase *m_usecaseTodos.MockTodoAdapter) {
				sts_false := false
				sts_true := true
				usecase.EXPECT().GetTodosByGroupId("1").Return([]domain.Todos{
					{
						ID:                1,
						Activity_group_id: 1,
						Title:             "Todo 1",
						Is_active:         &sts_false,
						Priority:          "very-high",
						GormModel: domain.GormModel{
							CreatedAt: &now,
							UpdatedAt: &now,
						},
					},
					{
						ID:                2,
						Activity_group_id: 1,
						Title:             "Todo 2",
						Is_active:         &sts_true,
						Priority:          "high",
						GormModel: domain.GormModel{
							CreatedAt: &now,
							UpdatedAt: &now,
						},
					},
				}, nil)
			},
		},
		{
			nameTest:   "Test Case 3 Get All Todos: Failed",
			statusCode: http.StatusInternalServerError,
			wantErr:    true,
			err: api.ResponseError{
				Status:  "Internal Server Error",
				Message: "Internal Server Error",
			},
			usecaseTest: func(usecase *m_usecaseTodos.MockTodoAdapter) {
				usecase.EXPECT().GetAllTodos().Return([]domain.Todos{}, errors.New("failed get data todos from database"))
			},
		},
	}

	for _, tt := range test_cases {
		t.Run(tt.nameTest, func(t *testing.T) {
			todoAdapter := m_usecaseTodos.NewMockTodoAdapter(ctrl)
			controller := &TodoController{
				todoUsecase: todoAdapter,
			}

			if tt.usecaseTest != nil {
				tt.usecaseTest(todoAdapter)
			}
			var req *http.Request
			router := fiber.New()
			router.Get("/todo-items", controller.GetAllTodos)
			if tt.activity_group_id != 0 {
				req = httptest.NewRequest("GET", "/todo-items?activity_group_id="+strconv.Itoa(tt.activity_group_id), nil)
			} else {
				req = httptest.NewRequest("GET", "/todo-items", nil)
			}
			response, err := router.Test(req, -1)
			assert.NoError(t, err)

			responseData, err := ioutil.ReadAll(response.Body)
			assert.NoError(t, err)

			//Testing Response and StatusCode
			assert.Equal(t, tt.statusCode, response.StatusCode)
			if !tt.wantErr {
				todoResult := api.ResponseSuccess{}
				err = json.Unmarshal(responseData, &todoResult)
				assert.NoError(t, err)
				assert.Equal(t, todoResult, tt.response)
			} else {
				todoResult := api.ResponseError{}
				err = json.Unmarshal(responseData, &todoResult)
				assert.NoError(t, err)
				assert.Equal(t, todoResult, tt.err)
			}
		})
	}
}
