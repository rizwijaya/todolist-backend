package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
	"todolist-backend/modules/v1/todos/domain"
	m_usecaseTodos "todolist-backend/modules/v1/todos/usecases/mock"
	api "todolist-backend/pkg/api_response"
	"todolist-backend/pkg/http_error"

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

func TestActivityController_GetTodoById(t *testing.T) {
	ctrl := gomock.NewController(t)
	now := time.Date(2023, time.April, 15, 10, 0, 0, 0, time.UTC)
	type tests struct {
		nameTest    string
		statusCode  int
		id          int
		response    api.ResponseSuccess
		wantErr     bool
		err         api.ResponseError
		usecaseTest func(usecase *m_usecaseTodos.MockTodoAdapter)
	}
	//add test case
	test_cases := []tests{
		{
			nameTest:   "Test Case 1 Get Todos By Id: Success",
			statusCode: http.StatusOK,
			id:         1,
			response: api.ResponseSuccess{
				Status:  "Success",
				Message: "Success",
				Data: map[string]interface{}{
					"id":                float64(1),
					"activity_group_id": float64(1),
					"title":             "Todo 1",
					"is_active":         true,
					"priority":          "very-high",
					"createdAt":         "2023-04-15T10:00:00Z",
					"updatedAt":         "2023-04-15T10:00:00Z",
				},
			},
			wantErr: false,
			usecaseTest: func(usecase *m_usecaseTodos.MockTodoAdapter) {
				sts := true
				usecase.EXPECT().GetTodoById("1").Return(domain.Todos{
					ID:                1,
					Activity_group_id: 1,
					Title:             "Todo 1",
					Is_active:         &sts,
					Priority:          "very-high",
					GormModel: domain.GormModel{
						CreatedAt: &now,
						UpdatedAt: &now,
					},
				}, nil)
			},
		},
		{
			nameTest:   "Test Case 2 Get Todos By Id: Failed Id Not Found",
			statusCode: http.StatusNotFound,
			id:         99999,
			wantErr:    true,
			err: api.ResponseError{
				Status:  "Not Found",
				Message: "Todo with ID 99999 Not Found",
			},
			usecaseTest: func(usecase *m_usecaseTodos.MockTodoAdapter) {
				usecase.EXPECT().GetTodoById("99999").Return(domain.Todos{}, http_error.ErrRecordNotfound)
			},
		},
		{
			nameTest:   "Test Case 3 Get Todos By Id: Failed Internal Server Error",
			statusCode: http.StatusInternalServerError,
			id:         999,
			wantErr:    true,
			err: api.ResponseError{
				Status:  "Internal Server Error",
				Message: "Internal Server Error",
			},
			usecaseTest: func(usecase *m_usecaseTodos.MockTodoAdapter) {
				usecase.EXPECT().GetTodoById("999").Return(domain.Todos{}, errors.New("failed get data todos from database"))
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
			router.Get("/todo-items/:id", controller.GetTodoById)
			req = httptest.NewRequest("GET", "/todo-items/"+strconv.Itoa(tt.id), nil)
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

func TestActivityController_CreateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	now := time.Date(2023, time.April, 15, 10, 0, 0, 0, time.UTC)
	type tests struct {
		nameTest           string
		statusCode         int
		request            domain.Todos
		requestInvalidJSON string
		response           api.ResponseSuccess
		wantErr            bool
		err                api.ResponseError
		usecaseTest        func(usecase *m_usecaseTodos.MockTodoAdapter)
	}
	sts_true := true
	//add test case
	test_cases := []tests{
		{
			nameTest:   "Test Case 1 Create Todo: Success",
			statusCode: http.StatusCreated,
			request: domain.Todos{
				Activity_group_id: 1,
				Title:             "Activity 1",
				Is_active:         &sts_true,
				Priority:          "very-high",
			},
			response: api.ResponseSuccess{
				Status:  "Success",
				Message: "Success",
				Data: map[string]interface{}{
					"id":                float64(1),
					"title":             "Activity 1",
					"activity_group_id": float64(1),
					"is_active":         true,
					"priority":          "very-high",
					"updatedAt":         "2023-04-15T10:00:00Z",
					"createdAt":         "2023-04-15T10:00:00Z",
				},
			},
			wantErr: false,
			usecaseTest: func(usecase *m_usecaseTodos.MockTodoAdapter) {
				usecase.EXPECT().CreateTodo(domain.Todos{
					Activity_group_id: 1,
					Title:             "Activity 1",
					Is_active:         &sts_true,
					Priority:          "very-high",
				}).Return(domain.Todos{
					ID:                1,
					Activity_group_id: 1,
					Title:             "Activity 1",
					Is_active:         &sts_true,
					Priority:          "very-high",
					GormModel: domain.GormModel{
						CreatedAt: &now,
						UpdatedAt: &now,
					},
				}, nil)
			},
		},
		{
			nameTest:   "Test Case 2 Create Todo: Failed With Invalid JSON",
			statusCode: http.StatusBadRequest,
			requestInvalidJSON: `{
				"title": "Todo 1",
			}`,
			wantErr: true,
			err: api.ResponseError{
				Status:  "Bad Request",
				Message: "invalid character '}' looking for beginning of object key string",
			},
		},
		{
			nameTest:   "Test Case 3 Create Todo: Without Activity Group Id",
			statusCode: http.StatusBadRequest,
			request: domain.Todos{
				Title:     "Activity 1",
				Is_active: &sts_true,
				Priority:  "very-high",
			},
			wantErr: true,
			err: api.ResponseError{
				Status:  "Bad Request",
				Message: "activity_group_id cannot be null",
			},
		},
		{
			nameTest:   "Test Case 4 Create Todo: Without Title",
			statusCode: http.StatusBadRequest,
			request: domain.Todos{
				Activity_group_id: 1,
				Is_active:         &sts_true,
				Priority:          "very-high",
			},
			wantErr: true,
			err: api.ResponseError{
				Status:  "Bad Request",
				Message: "title cannot be null",
			},
		},
		{
			nameTest:   "Test Case 5 Create Todo: Success Empty Is Active and Priority",
			statusCode: http.StatusCreated,
			request: domain.Todos{
				Activity_group_id: 1,
				Title:             "Activity 1",
			},
			response: api.ResponseSuccess{
				Status:  "Success",
				Message: "Success",
				Data: map[string]interface{}{
					"id":                float64(1),
					"title":             "Activity 1",
					"activity_group_id": float64(1),
					"is_active":         true,
					"priority":          "very-high",
					"updatedAt":         "2023-04-15T10:00:00Z",
					"createdAt":         "2023-04-15T10:00:00Z",
				},
			},
			wantErr: false,
			usecaseTest: func(usecase *m_usecaseTodos.MockTodoAdapter) {
				usecase.EXPECT().CreateTodo(domain.Todos{
					Activity_group_id: 1,
					Title:             "Activity 1",
					Is_active:         &sts_true,
					Priority:          "very-high",
				}).Return(domain.Todos{
					ID:                1,
					Activity_group_id: 1,
					Title:             "Activity 1",
					Is_active:         &sts_true,
					Priority:          "very-high",
					GormModel: domain.GormModel{
						CreatedAt: &now,
						UpdatedAt: &now,
					},
				}, nil)
			},
		},
		{
			nameTest:   "Test Case 6 Create Todo: Failed Internal Server Error",
			statusCode: http.StatusInternalServerError,
			request: domain.Todos{
				Activity_group_id: 1,
				Title:             "Activity 1",
			},
			wantErr: true,
			err: api.ResponseError{
				Status:  "Internal Server Error",
				Message: "Internal Server Error",
			},
			usecaseTest: func(usecase *m_usecaseTodos.MockTodoAdapter) {
				usecase.EXPECT().CreateTodo(domain.Todos{
					Activity_group_id: 1,
					Title:             "Activity 1",
					Is_active:         &sts_true,
					Priority:          "very-high",
				}).Return(domain.Todos{}, errors.New("failed insert data todo to database"))
			},
		},
	}

	for _, tt := range test_cases {
		t.Run(tt.nameTest, func(t *testing.T) {
			var req *http.Request
			todoAdapter := m_usecaseTodos.NewMockTodoAdapter(ctrl)
			controller := &TodoController{
				todoUsecase: todoAdapter,
			}

			if tt.usecaseTest != nil {
				tt.usecaseTest(todoAdapter)
			}

			router := fiber.New()
			router.Post("/todo-items", controller.CreateTodo)
			if tt.requestInvalidJSON != "" {
				req = httptest.NewRequest("POST", "/todo-items", strings.NewReader(tt.requestInvalidJSON))
			} else {
				val, err := json.Marshal(tt.request)
				assert.NoError(t, err)
				req = httptest.NewRequest("POST", "/todo-items", bytes.NewReader(val))
			}
			req.Header.Set("Content-Type", "application/json")
			response, err := router.Test(req, -1)
			assert.NoError(t, err)

			responseData, err := ioutil.ReadAll(response.Body)
			assert.NoError(t, err)

			//Testing Response and StatusCode
			assert.Equal(t, tt.statusCode, response.StatusCode)
			if !tt.wantErr {
				activityResult := api.ResponseSuccess{}
				err = json.Unmarshal(responseData, &activityResult)
				assert.NoError(t, err)
				assert.Equal(t, activityResult, tt.response)
			} else {
				activityResult := api.ResponseError{}
				err = json.Unmarshal(responseData, &activityResult)
				assert.NoError(t, err)
				assert.Equal(t, activityResult, tt.err)
			}
		})
	}
}

func TestActivityController_UpdateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	now := time.Date(2023, time.April, 15, 10, 0, 0, 0, time.UTC)
	type tests struct {
		nameTest           string
		statusCode         int
		id                 int
		request            domain.Todos
		requestInvalidJSON string
		response           api.ResponseSuccess
		wantErr            bool
		err                api.ResponseError
		usecaseTest        func(usecase *m_usecaseTodos.MockTodoAdapter)
	}
	sts_true := true
	//add test case
	test_cases := []tests{
		{
			nameTest:   "Test Case 1 Update Todo: Success",
			statusCode: http.StatusOK,
			id:         1,
			request: domain.Todos{
				Activity_group_id: 1,
				Title:             "Activity 1",
				Is_active:         &sts_true,
				Priority:          "very-high",
			},
			response: api.ResponseSuccess{
				Status:  "Success",
				Message: "Success",
				Data: map[string]interface{}{
					"id":                float64(1),
					"title":             "Activity 1",
					"activity_group_id": float64(1),
					"is_active":         true,
					"priority":          "very-high",
					"updatedAt":         "2023-04-15T10:00:00Z",
					"createdAt":         "2023-04-15T10:00:00Z",
				},
			},
			wantErr: false,
			usecaseTest: func(usecase *m_usecaseTodos.MockTodoAdapter) {
				usecase.EXPECT().UpdateTodo("1", domain.UpdateTodos{
					Activity_group_id: 1,
					Title:             "Activity 1",
					Is_active:         &sts_true,
					Priority:          "very-high",
				}).Return(domain.Todos{
					ID:                1,
					Activity_group_id: 1,
					Title:             "Activity 1",
					Is_active:         &sts_true,
					Priority:          "very-high",
					GormModel: domain.GormModel{
						CreatedAt: &now,
						UpdatedAt: &now,
					},
				}, nil)
			},
		},
		{
			nameTest:   "Test Case 2 Update Todo: Failed With ID Not Found",
			statusCode: http.StatusNotFound,
			id:         99999,
			request: domain.Todos{
				Activity_group_id: 99999,
				Title:             "Activity 99999",
				Is_active:         &sts_true,
				Priority:          "very-high",
			},
			wantErr: true,
			err: api.ResponseError{
				Status:  "Not Found",
				Message: "Todo with ID 99999 Not Found",
			},
			usecaseTest: func(usecase *m_usecaseTodos.MockTodoAdapter) {
				usecase.EXPECT().UpdateTodo("99999", domain.UpdateTodos{
					Activity_group_id: 99999,
					Title:             "Activity 99999",
					Is_active:         &sts_true,
					Priority:          "very-high",
				}).Return(domain.Todos{}, http_error.ErrRecordNotfound)
			},
		},
		{
			nameTest:   "Test Case 3 Update Todo: Failed With Invalid JSON",
			statusCode: http.StatusBadRequest,
			id:         1,
			requestInvalidJSON: `{
				"title": "Todo 1",
			}`,
			wantErr: true,
			err: api.ResponseError{
				Status:  "Bad Request",
				Message: "invalid character '}' looking for beginning of object key string",
			},
		},
		{
			nameTest:   "Test Case 4 Update Todo: Failed Internal Server Error",
			statusCode: http.StatusInternalServerError,
			id:         1,
			request: domain.Todos{
				Activity_group_id: 1,
				Title:             "Activity 1",
			},
			wantErr: true,
			err: api.ResponseError{
				Status:  "Internal Server Error",
				Message: "Internal Server Error",
			},
			usecaseTest: func(usecase *m_usecaseTodos.MockTodoAdapter) {
				usecase.EXPECT().UpdateTodo("1", domain.UpdateTodos{
					Activity_group_id: 1,
					Title:             "Activity 1",
				}).Return(domain.Todos{}, errors.New("failed insert data todo to database"))
			},
		},
	}

	for _, tt := range test_cases {
		t.Run(tt.nameTest, func(t *testing.T) {
			var req *http.Request
			todoAdapter := m_usecaseTodos.NewMockTodoAdapter(ctrl)
			controller := &TodoController{
				todoUsecase: todoAdapter,
			}

			if tt.usecaseTest != nil {
				tt.usecaseTest(todoAdapter)
			}

			router := fiber.New()
			router.Patch("/todo-items/:id", controller.UpdateTodo)
			if tt.requestInvalidJSON != "" {
				req = httptest.NewRequest("PATCH", "/todo-items/"+strconv.Itoa(tt.id), strings.NewReader(tt.requestInvalidJSON))
			} else {
				val, err := json.Marshal(tt.request)
				assert.NoError(t, err)
				req = httptest.NewRequest("PATCH", "/todo-items/"+strconv.Itoa(tt.id), bytes.NewReader(val))
			}
			req.Header.Set("Content-Type", "application/json")
			response, err := router.Test(req, -1)
			assert.NoError(t, err)

			responseData, err := ioutil.ReadAll(response.Body)
			assert.NoError(t, err)

			//Testing Response and StatusCode
			assert.Equal(t, tt.statusCode, response.StatusCode)
			if !tt.wantErr {
				activityResult := api.ResponseSuccess{}
				err = json.Unmarshal(responseData, &activityResult)
				assert.NoError(t, err)
				assert.Equal(t, activityResult, tt.response)
			} else {
				activityResult := api.ResponseError{}
				err = json.Unmarshal(responseData, &activityResult)
				assert.NoError(t, err)
				assert.Equal(t, activityResult, tt.err)
			}
		})
	}
}

func TestActivityController_DeleteTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	type tests struct {
		nameTest    string
		statusCode  int
		id          int
		response    api.ResponseSuccess
		wantErr     bool
		err         api.ResponseError
		usecaseTest func(usecase *m_usecaseTodos.MockTodoAdapter)
	}
	//add test case
	test_cases := []tests{
		{
			nameTest:   "Test Case 1 Delete Todos: Success",
			statusCode: http.StatusOK,
			id:         1,
			response: api.ResponseSuccess{
				Status:  "Success",
				Message: "Success",
				Data:    map[string]interface{}{},
			},
			wantErr: false,
			usecaseTest: func(usecase *m_usecaseTodos.MockTodoAdapter) {
				usecase.EXPECT().DeleteTodo("1").Return(nil)
			},
		},
		{
			nameTest:   "Test Case 2 Delete Todos: Failed Id Not Found",
			statusCode: http.StatusNotFound,
			id:         9999,
			wantErr:    true,
			err: api.ResponseError{
				Status:  "Not Found",
				Message: "Todo with ID 9999 Not Found",
			},
			usecaseTest: func(usecase *m_usecaseTodos.MockTodoAdapter) {
				usecase.EXPECT().DeleteTodo("9999").Return(http_error.ErrRecordNotfound)
			},
		},
		{
			nameTest:   "Test Case 3 Delete Todos: Failed Internal Server Error",
			statusCode: http.StatusInternalServerError,
			id:         122,
			wantErr:    true,
			err: api.ResponseError{
				Status:  "Internal Server Error",
				Message: "Internal Server Error",
			},
			usecaseTest: func(usecase *m_usecaseTodos.MockTodoAdapter) {
				usecase.EXPECT().DeleteTodo("122").Return(errors.New("failed to delete todos for database"))
			},
		},
	}

	for _, tt := range test_cases {
		t.Run(tt.nameTest, func(t *testing.T) {
			var req *http.Request
			todoAdapter := m_usecaseTodos.NewMockTodoAdapter(ctrl)
			controller := &TodoController{
				todoUsecase: todoAdapter,
			}

			if tt.usecaseTest != nil {
				tt.usecaseTest(todoAdapter)
			}

			router := fiber.New()
			router.Delete("/todo-items/:id", controller.DeleteTodo)
			req = httptest.NewRequest("DELETE", "/todo-items/"+strconv.Itoa(tt.id), nil)
			req.Header.Set("Content-Type", "application/json")
			response, err := router.Test(req, -1)
			assert.NoError(t, err)

			responseData, err := ioutil.ReadAll(response.Body)
			assert.NoError(t, err)

			//Testing Response and StatusCode
			assert.Equal(t, tt.statusCode, response.StatusCode)
			if !tt.wantErr {
				activityResult := api.ResponseSuccess{}
				err = json.Unmarshal(responseData, &activityResult)
				assert.NoError(t, err)
				assert.Equal(t, activityResult, tt.response)
			} else {
				activityResult := api.ResponseError{}
				err = json.Unmarshal(responseData, &activityResult)
				assert.NoError(t, err)
				assert.Equal(t, activityResult, tt.err)
			}
		})
	}
}
