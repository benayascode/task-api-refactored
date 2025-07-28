package Usecases

import (
	"context"
	"task_manager/Domain"
)

type TaskUseCase struct {
	TaskRepo Domain.TaskRepository
}

func (tu *TaskUseCase) GetAllTasks(ctx context.Context) ([]Domain.Task, error) {
	return tu.TaskRepo.GetAllTasks(ctx)
}

func (tu *TaskUseCase) GetTaskByID(ctx context.Context, id int) (Domain.Task, error) {
	return tu.TaskRepo.GetTaskByID(ctx, id)
}

func (tu *TaskUseCase) CreateTask(ctx context.Context, task Domain.Task) error {
	return tu.TaskRepo.CreateTask(ctx, task)
}

func (tu *TaskUseCase) UpdateTask(ctx context.Context, id int, task Domain.Task) error {
	return tu.TaskRepo.UpdateTask(ctx, id, task)
}

func (tu *TaskUseCase) DeleteTask(ctx context.Context, id int) error {
	return tu.TaskRepo.DeleteTask(ctx, id)
}
