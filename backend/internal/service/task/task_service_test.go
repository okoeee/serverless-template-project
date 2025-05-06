package task

import (
	"backend/internal/models"
	"context"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockTaskRepo struct{ mock.Mock }

func (m *mockTaskRepo) CreateTask(ctx context.Context, task *models.Task) error {
	args := m.Called(ctx, task)
	return args.Error(0)
}

func (m *mockTaskRepo) UpdateTask(ctx context.Context, task *models.Task) error {
	args := m.Called(ctx, task)
	return args.Error(0)
}

func TestTaskService_CreateTask(t *testing.T) {

	fixedUserId, _ := models.ParseUserId("123e4567-e89b-12d3-a456-426614174000")
	title := "Test Task"
	description := "Test Description"

	repo := new(mockTaskRepo)
	repo.On("CreateTask", mock.Anything, mock.Anything).Return(nil)

	service := NewTaskService(repo)

	ctx := context.Background()

	command := CreateTaskParam{
		UserId:      fixedUserId,
		Title:       title,
		Description: description,
		DueDate:     nil,
	}

	err := service.CreateTask(ctx, command)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	repo.AssertExpectations(t)

}

func TestTaskService_UpdateTask(t *testing.T) {

	fixedUserId, _ := models.ParseUserId("123e4567-e89b-12d3-a456-426614174000")
	title := "Test Task"
	description := "Test Description"

	repo := new(mockTaskRepo)
	repo.On("UpdateTask", mock.Anything, mock.Anything).Return(nil)

	service := NewTaskService(repo)

	ctx := context.Background()

	command := UpdateTaskParam{
		TaskId:      models.TaskId{},
		UserId:      fixedUserId,
		Title:       title,
		Description: description,
		DueDate:     nil,
		Status:      models.TaskStatusTodo,
	}

	err := service.UpdateTask(ctx, command)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	repo.AssertExpectations(t)
}
