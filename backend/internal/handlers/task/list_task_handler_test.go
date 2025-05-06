package task

import (
	"backend/internal/models"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockTaskRepo struct{ mock.Mock }

func (m *mockTaskRepo) FilterTaskByUserId(ctx context.Context, id models.UserId) ([]*models.Task, error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]*models.Task), args.Error(1)
}

func TestListTaskHandler_Handle(t *testing.T) {

	gin.SetMode(gin.TestMode)

	fixedUserId, _ := models.ParseUserId("123e4567-e89b-12d3-a456-426614174000")
	tasks := []*models.Task{
		{
			TaskId:      models.TaskId(uuid.New()),
			UserId:      fixedUserId,
			Title:       "Test Task 1",
			Description: "Description for test task 1",
			Status:      models.TaskStatusTodo,
		},
		{
			TaskId:      models.TaskId(uuid.New()),
			UserId:      fixedUserId,
			Title:       "Test Task 2",
			Description: "Description for test task 2",
			Status:      models.TaskStatusTodo,
		},
	}

	repo := new(mockTaskRepo)
	repo.On("FilterTaskByUserId", mock.Anything, fixedUserId).Return(tasks, nil)

	handler := NewListTaskHandler(repo)

	router := gin.New()
	router.GET("/task", handler.Handle)

	request := httptest.NewRequest(http.MethodGet, "/task", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", response.StatusCode)
	}

	repo.AssertExpectations(t)

}
