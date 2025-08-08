package usecases

import (
	"errors"
	"testing"
	"time"

	domain "task-manager/Domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTaskUseCase_GetAllTasks(t *testing.T) {
	repo := new(MockTaskRepository)
	uc := NewTaskUseCase(repo)

	expected := []domain.Task{{UserID: 1, Title: "t1"}, {UserID: 2, Title: "t2"}}
	repo.On("GetAllTasks").Return(expected, nil).Once()

	got, err := uc.GetAllTasks()
	assert.NoError(t, err)
	assert.Equal(t, expected, got)
	repo.AssertExpectations(t)
}

func TestTaskUseCase_GetTaskByID(t *testing.T) {
	repo := new(MockTaskRepository)
	uc := NewTaskUseCase(repo)

	repo.On("GetTaskByID", 7).Return(domain.Task{UserID: 7, Title: "x"}, nil).Once()

	got, err := uc.GetTaskByID(7)
	assert.NoError(t, err)
	assert.Equal(t, 7, got.UserID)
	repo.AssertExpectations(t)
}

func TestTaskUseCase_GetTaskByID_NotFound(t *testing.T) {
	repo := new(MockTaskRepository)
	uc := NewTaskUseCase(repo)

	repo.On("GetTaskByID", 999).Return(domain.Task{}, errors.New("not found")).Once()

	_, err := uc.GetTaskByID(999)
	assert.Error(t, err)
	repo.AssertExpectations(t)
}

func TestTaskUseCase_CreateTask_Validation(t *testing.T) {
	repo := new(MockTaskRepository)
	uc := NewTaskUseCase(repo)

	err := uc.CreateTask(domain.Task{Title: "", Description: "d"})
	assert.ErrorIs(t, err, domain.ErrInvalidTaskTitle)
	repo.AssertNotCalled(t, "CreateTask", mock.Anything)

	err = uc.CreateTask(domain.Task{Title: "t", Description: ""})
	assert.ErrorIs(t, err, domain.ErrInvalidTaskDescription)
	repo.AssertNotCalled(t, "CreateTask", mock.Anything)
}

func TestTaskUseCase_CreateTask_Success(t *testing.T) {
	repo := new(MockTaskRepository)
	uc := NewTaskUseCase(repo)

	task := domain.Task{UserID: 1, Title: "title", Description: "desc", DueDate: time.Now(), Status: "open"}
	repo.On("CreateTask", mock.Anything).Return(nil).Once()

	err := uc.CreateTask(task)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestTaskUseCase_UpdateTask_Validation(t *testing.T) {
	repo := new(MockTaskRepository)
	uc := NewTaskUseCase(repo)

	_, err := uc.UpdateTask(1, domain.Task{Title: "", Description: "d"})
	assert.ErrorIs(t, err, domain.ErrInvalidTaskTitle)
	repo.AssertNotCalled(t, "UpdateTask", mock.Anything)

	_, err = uc.UpdateTask(1, domain.Task{Title: "t", Description: ""})
	assert.ErrorIs(t, err, domain.ErrInvalidTaskDescription)
	repo.AssertNotCalled(t, "UpdateTask", mock.Anything)
}

func TestTaskUseCase_UpdateTask_Success(t *testing.T) {
	repo := new(MockTaskRepository)
	uc := NewTaskUseCase(repo)

	upd := domain.Task{UserID: 1, Title: "new", Description: "d", Status: "done"}
	repo.On("UpdateTask", 1, upd).Return(upd, nil).Once()

	got, err := uc.UpdateTask(1, upd)
	assert.NoError(t, err)
	assert.Equal(t, upd, got)
	repo.AssertExpectations(t)
}

func TestTaskUseCase_DeleteTask(t *testing.T) {
	repo := new(MockTaskRepository)
	uc := NewTaskUseCase(repo)

	repo.On("DeleteTask", 2).Return(nil).Once()
	assert.NoError(t, uc.DeleteTask(2))
	repo.AssertExpectations(t)
}

func TestTaskUseCase_DeleteTask_Error(t *testing.T) {
	repo := new(MockTaskRepository)
	uc := NewTaskUseCase(repo)

	repo.On("DeleteTask", 3).Return(errors.New("not found")).Once()
	assert.Error(t, uc.DeleteTask(3))
	repo.AssertExpectations(t)
}
