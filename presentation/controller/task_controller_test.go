package controller_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fuki01/onion-architecture/domain/task"
	"github.com/fuki01/onion-architecture/domain/user"
	"github.com/fuki01/onion-architecture/presentation/controller"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskUsecase struct {
	mock.Mock
}

func (m *MockTaskUsecase) CreateTask(name string, userId user.UserId, dueDate string) (task.TaskId, error) {
	args := m.Called(name, userId, dueDate)
	return args.Get(0).(task.TaskId), args.Error(1)
}

func (m *MockTaskUsecase) ExtendDueDate(id task.TaskId, dueDate string) error {
	args := m.Called(id, dueDate)
	return args.Error(0)
}

func (m *MockTaskUsecase) ChangeStatus(id task.TaskId, newStatus task.TaskStatus) error {
	args := m.Called(id, newStatus)
	return args.Error(0)
}

func TestNewTaskController(t *testing.T) {
	mockUsecase := new(MockTaskUsecase)
	tc := controller.NewTaskController(mockUsecase)
	assert.NotNil(t, tc)
}

func TestTaskControllerCreateTask(t *testing.T) {
	testCases := []struct {
		name           string
		mockSetup      func(m *MockTaskUsecase)
		reqBody        string
		expectedStatus int
	}{
		{
			name: "Success",
			mockSetup: func(m *MockTaskUsecase) {
				m.On("CreateTask", "タスク名", user.UserId(1), "2021-01-01").Return(task.TaskId(1), nil)
			},
			reqBody:        `{"name":"タスク名","user_id":1,"due_date":"2021-01-01"}`,
			expectedStatus: http.StatusCreated,
		},
		{
			name: "Usecase Error",
			mockSetup: func(m *MockTaskUsecase) {
				m.On("CreateTask", "タスク名", user.UserId(1), "2021-01-01").Return(task.TaskId(0), fmt.Errorf("error"))
			},
			reqBody:        `{"name":"タスク名","user_id":1,"due_date":"2021-01-01"}`,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockUsecase := new(MockTaskUsecase)
			tc.mockSetup(mockUsecase)

			controller := controller.NewTaskController(mockUsecase)

			req, _ := http.NewRequest("POST", "/tasks", bytes.NewBufferString(tc.reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()

			r := gin.Default()
			r.POST("/tasks", controller.CreateTask)
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
			mockUsecase.AssertExpectations(t)
		})
	}
}

func TestTaskControllerExtendDueDate(t *testing.T) {
	testCases := []struct {
		name           string
		mockSetup      func(m *MockTaskUsecase)
		reqBody        string
		expectedStatus int
	}{
		{
			name: "Success",
			mockSetup: func(m *MockTaskUsecase) {
				m.On("ExtendDueDate", task.TaskId(1), "2021-01-01").Return(nil)
			},
			reqBody:        `{"id":1,"due_date":"2021-01-01"}`,
			expectedStatus: http.StatusOK,
		},
		{
			name: "Usecase Error",
			mockSetup: func(m *MockTaskUsecase) {
				m.On("ExtendDueDate", task.TaskId(1), "2021-01-01").Return(fmt.Errorf("error"))
			},
			reqBody:        `{"id":1,"due_date":"2021-01-01"}`,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockUsecase := new(MockTaskUsecase)
			tc.mockSetup(mockUsecase)

			controller := controller.NewTaskController(mockUsecase)

			req, _ := http.NewRequest("PUT", "/tasks/1/extend_due_date", bytes.NewBufferString(tc.reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()

			r := gin.Default()
			r.PUT("/tasks/:id/extend_due_date",
				controller.ExtendDueDate)
			r.ServeHTTP(w, req)

			fmt.Println(w.Body.String())
			assert.Equal(t, tc.expectedStatus, w.Code)
			mockUsecase.AssertExpectations(t)
		})
	}
}

func TestTaskControllerChangeStatus(t *testing.T) {
	testCases := []struct {
		name           string
		mockSetup      func(m *MockTaskUsecase)
		reqBody        string
		expectedStatus int
	}{
		{
			name: "Success",
			mockSetup: func(m *MockTaskUsecase) {
				m.On("ChangeStatus", task.TaskId(1), task.StatusComplete).Return(nil)
			},
			reqBody:        `{"id":1,"new_status":"完了"}`,
			expectedStatus: http.StatusOK,
		},
		{
			name: "Usecase Error",
			mockSetup: func(m *MockTaskUsecase) {
				m.On("ChangeStatus", task.TaskId(1), task.StatusComplete).Return(fmt.Errorf("error"))
			},
			reqBody:        `{"id":1,"new_status":"完了"}`,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockUsecase := new(MockTaskUsecase)
			tc.mockSetup(mockUsecase)

			controller := controller.NewTaskController(mockUsecase)

			req, _ := http.NewRequest("PUT", "/tasks/1/change_status", bytes.NewBufferString(tc.reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()

			r := gin.Default()
			r.PUT("/tasks/:id/change_status", controller.ChangeStatus)
			r.ServeHTTP(w, req)

			fmt.Println("body", w.Body.String())
			assert.Equal(t, tc.expectedStatus, w.Code)
			mockUsecase.AssertExpectations(t)
		})
	}
}
