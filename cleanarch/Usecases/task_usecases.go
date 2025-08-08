package usecases

import (
	domain "task-manager/Domain"
)

type TaskUseCaseImpl struct {
	taskRepository domain.TaskRepository
}

func NewTaskUseCase(taskRepository domain.TaskRepository) domain.TaskUseCase {
	return &TaskUseCaseImpl{
		taskRepository: taskRepository,
	}
}

func (t *TaskUseCaseImpl) GetAllTasks() ([]domain.Task, error) {
	return t.taskRepository.GetAllTasks()
}

func (t *TaskUseCaseImpl) GetTaskByID(userID int) (domain.Task, error) {
	return t.taskRepository.GetTaskByID(userID)
}

func (t *TaskUseCaseImpl) CreateTask(task domain.Task) error {
	if err := task.Validate(); err != nil {
		return err
	}
	return t.taskRepository.CreateTask(task)
}

func (t *TaskUseCaseImpl) UpdateTask(userID int, task domain.Task) (domain.Task, error) {
	if err := task.Validate(); err != nil {
		return domain.Task{}, err
	}
	return t.taskRepository.UpdateTask(userID, task)
}

func (t *TaskUseCaseImpl) DeleteTask(userID int) error {
	return t.taskRepository.DeleteTask(userID)
}
