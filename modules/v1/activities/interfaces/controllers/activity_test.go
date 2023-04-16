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
	"todolist-backend/modules/v1/activities/domain"
	m_usecaseActivity "todolist-backend/modules/v1/activities/usecases/mock"
	api "todolist-backend/pkg/api_response"
	"todolist-backend/pkg/http_error"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestActivityController_GetAllActivity(t *testing.T) {
	ctrl := gomock.NewController(t)
	now := time.Date(2023, time.April, 15, 10, 0, 0, 0, time.UTC)
	type tests struct {
		nameTest    string
		statusCode  int
		response    api.ResponseSuccess
		wantErr     bool
		err         api.ResponseError
		usecaseTest func(usecase *m_usecaseActivity.MockActivityAdapter)
	}
	//add test case
	test_cases := []tests{
		{
			nameTest:   "Test Case 1 Get All Activity: Success",
			statusCode: http.StatusOK,
			response: api.ResponseSuccess{
				Status:  "Success",
				Message: "Success",
				Data: []interface{}{
					map[string]interface{}{
						"id":        float64(1),
						"title":     "Activity 1",
						"email":     "activ@gmail.com",
						"createdAt": "2023-04-15T10:00:00Z",
						"updatedAt": "2023-04-15T10:00:00Z",
					},
					map[string]interface{}{
						"id":        float64(2),
						"title":     "Activity 2",
						"email":     "ac2@gmail.com",
						"createdAt": "2023-04-15T10:00:00Z",
						"updatedAt": "2023-04-15T10:00:00Z",
					},
				},
			},
			wantErr: false,
			usecaseTest: func(usecase *m_usecaseActivity.MockActivityAdapter) {
				usecase.EXPECT().GetAllActivity().Return([]domain.Activities{
					{
						ID:    1,
						Title: "Activity 1",
						Email: "activ@gmail.com",
						GormModel: domain.GormModel{
							CreatedAt: &now,
							UpdatedAt: &now,
						},
					},
					{
						ID:    2,
						Title: "Activity 2",
						Email: "ac2@gmail.com",
						GormModel: domain.GormModel{
							CreatedAt: &now,
							UpdatedAt: &now,
						},
					},
				}, nil)
			},
		},
		{
			nameTest:   "Test Case 2 Get All Activity: Success But Empty Data",
			statusCode: http.StatusOK,
			response: api.ResponseSuccess{
				Status:  "Success",
				Message: "Success",
				Data:    []interface{}{},
			},
			wantErr: false,
			usecaseTest: func(usecase *m_usecaseActivity.MockActivityAdapter) {
				usecase.EXPECT().GetAllActivity().Return([]domain.Activities{}, nil)
			},
		},
		{
			nameTest:   "Test Case 3 Get All Activity: Failed",
			statusCode: http.StatusInternalServerError,
			wantErr:    true,
			err: api.ResponseError{
				Status:  "Internal Server Error",
				Message: "Internal Server Error",
			},
			usecaseTest: func(usecase *m_usecaseActivity.MockActivityAdapter) {
				usecase.EXPECT().GetAllActivity().Return([]domain.Activities{}, errors.New("failed get data activity from database"))
			},
		},
	}

	for _, tt := range test_cases {
		t.Run(tt.nameTest, func(t *testing.T) {
			activityAdapter := m_usecaseActivity.NewMockActivityAdapter(ctrl)
			controller := &ActivityController{
				activityUsecase: activityAdapter,
			}

			if tt.usecaseTest != nil {
				tt.usecaseTest(activityAdapter)
			}

			router := fiber.New()
			router.Get("/activity-groups", controller.GetAllActivity)
			req := httptest.NewRequest("GET", "/activity-groups", nil)
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

func TestActivityController_GetActivityByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	now := time.Date(2023, time.April, 15, 10, 0, 0, 0, time.UTC)
	type tests struct {
		nameTest    string
		statusCode  int
		request     int
		response    api.ResponseSuccess
		wantErr     bool
		err         api.ResponseError
		usecaseTest func(usecase *m_usecaseActivity.MockActivityAdapter)
	}
	//add test case
	test_cases := []tests{
		{
			nameTest:   "Test Case 1 Get Activity By ID: Success",
			statusCode: http.StatusOK,
			request:    1,
			response: api.ResponseSuccess{
				Status:  "Success",
				Message: "Success",
				Data: map[string]interface{}{
					"id":        float64(1),
					"title":     "Activity 1",
					"email":     "activ@gmail.com",
					"createdAt": "2023-04-15T10:00:00Z",
					"updatedAt": "2023-04-15T10:00:00Z",
				},
			},
			wantErr: false,
			usecaseTest: func(usecase *m_usecaseActivity.MockActivityAdapter) {
				usecase.EXPECT().GetActivityByID("1").Return(domain.Activities{
					ID:    1,
					Title: "Activity 1",
					Email: "activ@gmail.com",
					GormModel: domain.GormModel{
						CreatedAt: &now,
						UpdatedAt: &now,
					},
				}, nil)
			},
		},
		{
			nameTest:   "Test Case 2 Get Activity By ID: ID Not Found",
			statusCode: http.StatusNotFound,
			request:    2222,
			wantErr:    true,
			err: api.ResponseError{
				Status:  "Not Found",
				Message: "Activity with ID 2222 Not Found",
			},
			usecaseTest: func(usecase *m_usecaseActivity.MockActivityAdapter) {
				usecase.EXPECT().GetActivityByID("2222").Return(domain.Activities{}, http_error.ErrRecordNotfound)
			},
		},
		{
			nameTest:   "Test Case 3 Get Activity By ID: Internal Server Error",
			statusCode: http.StatusInternalServerError,
			request:    22,
			wantErr:    true,
			err: api.ResponseError{
				Status:  "Internal Server Error",
				Message: "Internal Server Error",
			},
			usecaseTest: func(usecase *m_usecaseActivity.MockActivityAdapter) {
				usecase.EXPECT().GetActivityByID("22").Return(domain.Activities{}, errors.New("failed get data activity from database"))
			},
		},
	}

	for _, tt := range test_cases {
		t.Run(tt.nameTest, func(t *testing.T) {
			activityAdapter := m_usecaseActivity.NewMockActivityAdapter(ctrl)
			controller := &ActivityController{
				activityUsecase: activityAdapter,
			}

			if tt.usecaseTest != nil {
				tt.usecaseTest(activityAdapter)
			}

			router := fiber.New()
			router.Get("/activity-groups/:id", controller.GetActivityByID)
			req := httptest.NewRequest("GET", "/activity-groups/"+strconv.Itoa(tt.request), nil)
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

func TestActivityController_CreateActivity(t *testing.T) {
	ctrl := gomock.NewController(t)
	now := time.Date(2023, time.April, 15, 10, 0, 0, 0, time.UTC)
	type tests struct {
		nameTest           string
		statusCode         int
		request            domain.InsertActivity
		requestInvalidJSON string
		response           api.ResponseSuccess
		wantErr            bool
		err                api.ResponseError
		usecaseTest        func(usecase *m_usecaseActivity.MockActivityAdapter)
	}
	//add test case
	test_cases := []tests{
		{
			nameTest:   "Test Case 1 Create Activity: Success With Email",
			statusCode: http.StatusCreated,
			request: domain.InsertActivity{
				Title: "Activity 1",
				Email: "activity@mail.com",
			},
			response: api.ResponseSuccess{
				Status:  "Success",
				Message: "Success",
				Data: map[string]interface{}{
					"id":        float64(1),
					"title":     "Activity 1",
					"email":     "activity@mail.com",
					"createdAt": "2023-04-15T10:00:00Z",
					"updatedAt": "2023-04-15T10:00:00Z",
				},
			},
			wantErr: false,
			usecaseTest: func(usecase *m_usecaseActivity.MockActivityAdapter) {
				usecase.EXPECT().CreateActivity(domain.InsertActivity{
					Title: "Activity 1",
					Email: "activity@mail.com",
				}).Return(domain.Activities{
					ID:    1,
					Title: "Activity 1",
					Email: "activity@mail.com",
					GormModel: domain.GormModel{
						CreatedAt: &now,
						UpdatedAt: &now,
					},
				}, nil)
			},
		},
		{
			nameTest:   "Test Case 2 Create Activity: Success Without Email",
			statusCode: http.StatusCreated,
			request: domain.InsertActivity{
				Title: "Activity 1",
			},
			response: api.ResponseSuccess{
				Status:  "Success",
				Message: "Success",
				Data: map[string]interface{}{
					"id":        float64(1),
					"title":     "Activity 1",
					"createdAt": "2023-04-15T10:00:00Z",
					"updatedAt": "2023-04-15T10:00:00Z",
				},
			},
			wantErr: false,
			usecaseTest: func(usecase *m_usecaseActivity.MockActivityAdapter) {
				usecase.EXPECT().CreateActivity(domain.InsertActivity{
					Title: "Activity 1",
				}).Return(domain.Activities{
					ID:    1,
					Title: "Activity 1",
					GormModel: domain.GormModel{
						CreatedAt: &now,
						UpdatedAt: &now,
					},
				}, nil)
			},
		},
		{
			nameTest:   "Test Case 3 Create Activity: Failed With Invalid JSON",
			statusCode: http.StatusBadRequest,
			requestInvalidJSON: `{
				"title": "Activity 1",
			}`,
			wantErr: true,
			err: api.ResponseError{
				Status:  "Bad Request",
				Message: "invalid character '}' looking for beginning of object key string",
			},
		},
		{
			nameTest:   "Test Case 4 Create Activity: Failed With Empty Title",
			statusCode: http.StatusBadRequest,
			request: domain.InsertActivity{
				Email: "activ@mail.com",
			},
			wantErr: true,
			err: api.ResponseError{
				Status:  "Bad Request",
				Message: "Title cannot be null",
			},
		},
		{
			nameTest:   "Test Case 5 Create Activity: Failed Internal Server Error",
			statusCode: http.StatusInternalServerError,
			request: domain.InsertActivity{
				Email: "activ@mail.com",
				Title: "Activity 1",
			},
			wantErr: true,
			err: api.ResponseError{
				Status:  "Internal Server Error",
				Message: "Internal Server Error",
			},
			usecaseTest: func(usecase *m_usecaseActivity.MockActivityAdapter) {
				usecase.EXPECT().CreateActivity(domain.InsertActivity{
					Email: "activ@mail.com",
					Title: "Activity 1",
				}).Return(domain.Activities{}, errors.New("failed to create activity for database"))
			},
		},
	}

	for _, tt := range test_cases {
		t.Run(tt.nameTest, func(t *testing.T) {
			var req *http.Request
			activityAdapter := m_usecaseActivity.NewMockActivityAdapter(ctrl)
			controller := &ActivityController{
				activityUsecase: activityAdapter,
			}

			if tt.usecaseTest != nil {
				tt.usecaseTest(activityAdapter)
			}

			router := fiber.New()
			router.Post("/activity-groups", controller.CreateActivity)
			if tt.requestInvalidJSON != "" {
				req = httptest.NewRequest("POST", "/activity-groups", strings.NewReader(tt.requestInvalidJSON))
			} else {
				val, err := json.Marshal(tt.request)
				assert.NoError(t, err)
				req = httptest.NewRequest("POST", "/activity-groups", bytes.NewReader(val))
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
