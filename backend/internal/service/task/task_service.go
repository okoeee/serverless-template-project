package task

import (
	"backend/internal/models"
	"context"
	"time"
)

type Repository interface {
	CreateTask(ctx context.Context, task *models.Task) error
	UpdateTask(ctx context.Context, task *models.Task) error
	DeleteTask(ctx context.Context, taskId models.TaskId) error
}

type TaskService struct {
	repository Repository
}

func NewTaskService(repository Repository) *TaskService {
	return &TaskService{
		repository: repository,
	}
}

type CreateTaskParam struct {
	UserId      models.UserId
	Title       string
	Description string
	DueDate     *time.Time
}

func (s *TaskService) CreateTask(ctx context.Context, param CreateTaskParam) error {

	task, err := models.NewTask(
		param.UserId,
		param.Title,
		param.Description,
		param.DueDate,
	)
	if err != nil {
		return err
	}

	err = s.repository.CreateTask(ctx, task)
	if err != nil {
		return err
	}

	return nil

}

type UpdateTaskParam struct {
	TaskId      models.TaskId
	UserId      models.UserId
	Title       string
	Description string
	DueDate     *time.Time
	Status      models.TaskStatus
}

func (s *TaskService) UpdateTask(ctx context.Context, param UpdateTaskParam) error {
	task := &models.Task{
		TaskId:      param.TaskId,
		UserId:      param.UserId,
		Title:       param.Title,
		Description: param.Description,
		DueDate:     param.DueDate,
		Status:      param.Status,
	}

	err := s.repository.UpdateTask(ctx, task)
	if err != nil {
		return err
	}

	return nil
}

func (s *TaskService) DeleteTask(ctx context.Context, taskId models.TaskId) error {
	err := s.repository.DeleteTask(ctx, taskId)
	if err != nil {
		return err
	}

	return nil
}
