package task

import (
	"backend/internal/models"
	"context"
	"time"
)

type Repository interface {
	CreateTask(ctx context.Context, task *models.Task) error
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
