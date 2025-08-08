package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	domain "task-manager/Domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocks for use cases

type MockTaskUseCase struct{ mock.Mock }

type MockUserUseCase struct{ mock.Mock }

func (m *MockTaskUseCase) GetAllTasks() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *MockTaskUseCase) GetTaskByID(userID int) (domain.Task, error) {
	args := m.Called(userID)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskUseCase) CreateTask(task domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskUseCase) UpdateTask(userID int, task domain.Task) (domain.Task, error) {
	args := m.Called(userID, task)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskUseCase) DeleteTask(userID int) error {
	args := m.Called(userID)
	return args.Error(0)
}

func (m *MockUserUseCase) RegisterUser(username, password string) error {
	args := m.Called(username, password)
	return args.Error(0)
}

func (m *MockUserUseCase) LoginUser(username, password string) (string, error) {
	args := m.Called(username, password)
	return args.String(0), args.Error(1)
}

func (m *MockUserUseCase) PromoteUser(username string) error {
	args := m.Called(username)
	return args.Error(0)
}

func setupGin() {
	gin.SetMode(gin.TestMode)
}

func TestGetTasks_Success(t *testing.T) {
	setupGin()
	rec := httptest.NewRecorder()
	c, r := gin.CreateTestContext(rec)

	mockUC := new(MockTaskUseCase)
	tasks := []domain.Task{{UserID: 1, Title: "A"}}
	mockUC.On("GetAllTasks").Return(tasks, nil).Once()
	ctrl := NewTaskController(mockUC)

	r.GET("/tasks", ctrl.GetTasks)
	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUC.AssertExpectations(t)

	_ = c
}

func TestGetTaskByID_InvalidID(t *testing.T) {
	setupGin()
	rec := httptest.NewRecorder()
	_, r := gin.CreateTestContext(rec)
	mockUC := new(MockTaskUseCase)
	ctrl := NewTaskController(mockUC)

	r.GET("/tasks/:id", ctrl.GetTaskByID)
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/tasks/abc", nil))

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	mockUC.AssertNotCalled(t, "GetTaskByID", mock.Anything)
}

func TestGetTaskByID_NotFound(t *testing.T) {
	setupGin()
	rec := httptest.NewRecorder()
	_, r := gin.CreateTestContext(rec)
	mockUC := new(MockTaskUseCase)
	ctrl := NewTaskController(mockUC)

	mockUC.On("GetTaskByID", 999).Return(domain.Task{}, assert.AnError).Once()
	r.GET("/tasks/:id", ctrl.GetTaskByID)
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/tasks/999", nil))

	assert.Equal(t, http.StatusNotFound, rec.Code)
	mockUC.AssertExpectations(t)
}

func TestCreateTask_ValidationError(t *testing.T) {
	setupGin()
	rec := httptest.NewRecorder()
	_, r := gin.CreateTestContext(rec)
	mockUC := new(MockTaskUseCase)
	ctrl := NewTaskController(mockUC)

	body, _ := json.Marshal(map[string]any{"title": "", "description": "d"})
	mockUC.On("CreateTask", mock.AnythingOfType("domain.Task")).Return(domain.ErrInvalidTaskTitle).Once()

	r.POST("/tasks", ctrl.CreateTask)
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(body)))

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	mockUC.AssertExpectations(t)
}

func TestCreateTask_Success(t *testing.T) {
	setupGin()
	rec := httptest.NewRecorder()
	_, r := gin.CreateTestContext(rec)
	mockUC := new(MockTaskUseCase)
	ctrl := NewTaskController(mockUC)

	payload := domain.Task{Title: "A", Description: "B"}
	b, _ := json.Marshal(payload)
	mockUC.On("CreateTask", mock.AnythingOfType("domain.Task")).Return(nil).Once()

	r.POST("/tasks", ctrl.CreateTask)
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(b)))

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUC.AssertExpectations(t)
}

func TestUpdateTask_InvalidID(t *testing.T) {
	setupGin()
	rec := httptest.NewRecorder()
	_, r := gin.CreateTestContext(rec)
	mockUC := new(MockTaskUseCase)
	ctrl := NewTaskController(mockUC)

	r.PUT("/tasks/:id", ctrl.UpdateTask)
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodPut, "/tasks/abc", bytes.NewReader([]byte(`{"title":"x","description":"y"}`))))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	mockUC.AssertNotCalled(t, "UpdateTask", mock.Anything, mock.Anything)
}

func TestUpdateTask_NotFound(t *testing.T) {
	setupGin()
	rec := httptest.NewRecorder()
	_, r := gin.CreateTestContext(rec)
	mockUC := new(MockTaskUseCase)
	ctrl := NewTaskController(mockUC)

	mockUC.On("UpdateTask", 1, mock.AnythingOfType("domain.Task")).Return(domain.Task{}, assert.AnError).Once()
	r.PUT("/tasks/:id", ctrl.UpdateTask)
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodPut, "/tasks/1", bytes.NewReader([]byte(`{"title":"x","description":"y"}`))))
	assert.Equal(t, http.StatusNotFound, rec.Code)
	mockUC.AssertExpectations(t)
}

func TestDeleteTask_InvalidID(t *testing.T) {
	setupGin()
	rec := httptest.NewRecorder()
	_, r := gin.CreateTestContext(rec)
	mockUC := new(MockTaskUseCase)
	ctrl := NewTaskController(mockUC)

	r.DELETE("/tasks/:id", ctrl.DeleteTask)
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodDelete, "/tasks/abc", nil))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	mockUC.AssertNotCalled(t, "DeleteTask", mock.Anything)
}

func TestDeleteTask_Success(t *testing.T) {
	setupGin()
	rec := httptest.NewRecorder()
	_, r := gin.CreateTestContext(rec)
	mockUC := new(MockTaskUseCase)
	ctrl := NewTaskController(mockUC)

	mockUC.On("DeleteTask", 1).Return(nil).Once()
	r.DELETE("/tasks/:id", ctrl.DeleteTask)
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodDelete, "/tasks/1", nil))
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUC.AssertExpectations(t)
}

func TestRegister_Success(t *testing.T) {
	setupGin()
	rec := httptest.NewRecorder()
	_, r := gin.CreateTestContext(rec)
	mockUC := new(MockUserUseCase)
	ctrl := NewUserController(mockUC)

	mockUC.On("RegisterUser", "bob", "secret").Return(nil).Once()
	r.POST("/register", ctrl.Register)
	body := []byte(`{"username":"bob","password":"secret"}`)
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(body)))
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUC.AssertExpectations(t)
}

func TestLogin_Success(t *testing.T) {
	setupGin()
	rec := httptest.NewRecorder()
	_, r := gin.CreateTestContext(rec)
	mockUC := new(MockUserUseCase)
	ctrl := NewUserController(mockUC)

	mockUC.On("LoginUser", "bob", "secret").Return("tok", nil).Once()
	r.POST("/login", ctrl.Login)
	body := []byte(`{"username":"bob","password":"secret"}`)
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body)))
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUC.AssertExpectations(t)
}

func TestLogin_Invalid(t *testing.T) {
	setupGin()
	rec := httptest.NewRecorder()
	_, r := gin.CreateTestContext(rec)
	mockUC := new(MockUserUseCase)
	ctrl := NewUserController(mockUC)

	mockUC.On("LoginUser", "bob", "bad").Return("", assert.AnError).Once()
	r.POST("/login", ctrl.Login)
	body := []byte(`{"username":"bob","password":"bad"}`)
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body)))
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	mockUC.AssertExpectations(t)
}

func TestPromoteUser_Success(t *testing.T) {
	setupGin()
	rec := httptest.NewRecorder()
	_, r := gin.CreateTestContext(rec)
	mockUC := new(MockUserUseCase)
	ctrl := NewUserController(mockUC)

	mockUC.On("PromoteUser", "bob").Return(nil).Once()
	r.POST("/promote/:username", ctrl.PromoteUser)
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/promote/bob", nil))
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUC.AssertExpectations(t)
}
