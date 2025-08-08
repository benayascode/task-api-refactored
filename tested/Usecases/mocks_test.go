package usecases

import (
	domain "task-manager/Domain"
	infrastructure "task-manager/Infrastructure"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
)

// MockTaskRepository mocks domain.TaskRepository
type MockTaskRepository struct{ mock.Mock }

func (m *MockTaskRepository) GetAllTasks() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *MockTaskRepository) GetTaskByID(userID int) (domain.Task, error) {
	args := m.Called(userID)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskRepository) CreateTask(task domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) UpdateTask(userID int, task domain.Task) (domain.Task, error) {
	args := m.Called(userID, task)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskRepository) DeleteTask(userID int) error {
	args := m.Called(userID)
	return args.Error(0)
}

// MockUserRepository mocks domain.UserRepository
type MockUserRepository struct{ mock.Mock }

func (m *MockUserRepository) RegisterUser(username, password string) error {
	args := m.Called(username, password)
	return args.Error(0)
}

func (m *MockUserRepository) AuthenticateUser(username, password string) (domain.User, error) {
	args := m.Called(username, password)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepository) PromoteUser(username string) error {
	args := m.Called(username)
	return args.Error(0)
}

// MockPasswordService mocks infrastructure.PasswordService
type MockPasswordService struct{ mock.Mock }

func (m *MockPasswordService) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockPasswordService) ComparePassword(hashedPassword, password string) error {
	args := m.Called(hashedPassword, password)
	return args.Error(0)
}

// MockJWTService mocks infrastructure.JWTService
type MockJWTService struct{ mock.Mock }

func (m *MockJWTService) GenerateToken(user domain.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func (m *MockJWTService) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	args := m.Called(tokenString)
	if claims, ok := args.Get(0).(jwt.MapClaims); ok {
		return claims, args.Error(1)
	}
	return nil, args.Error(1)
}

var _ domain.TaskRepository = (*MockTaskRepository)(nil)
var _ domain.UserRepository = (*MockUserRepository)(nil)
var _ infrastructure.PasswordService = (*MockPasswordService)(nil)
var _ infrastructure.JWTService = (*MockJWTService)(nil)
