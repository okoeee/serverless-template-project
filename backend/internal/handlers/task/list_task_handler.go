package task

import (
	"backend/internal/models"
	"backend/internal/request/json/writes"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Repository interface {
	FilterTaskByUserId(ctx context.Context, userId models.UserId) ([]*models.Task, error)
}

type ListTaskHandler struct {
	TaskRepository Repository
}

func NewListTaskHandler(taskRepository Repository) *ListTaskHandler {
	return &ListTaskHandler{
		TaskRepository: taskRepository,
	}
}

func (h *ListTaskHandler) Handle(c *gin.Context) {

	// TODO Dummy UserId
	userId, err := models.ParseUserId("123e4567-e89b-12d3-a456-426614174000")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	ctx := c.Request.Context()
	tasks, err := h.TaskRepository.FilterTaskByUserId(ctx, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks: " + err.Error()})
		return
	}

	tasksResponse := make([]writes.TaskResponse, len(tasks))
	for i, task := range tasks {
		tasksResponse[i] = writes.TaskResponse{
			TaskId:      task.TaskId.String(),
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			DueDate:     task.DueDate,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasksResponse,
	})

}
