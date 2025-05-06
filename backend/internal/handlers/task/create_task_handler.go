package task

import (
	"backend/internal/models"
	"backend/internal/request/json/reads"
	"backend/internal/service/task"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateTaskHandler struct {
	taskService *task.TaskService
}

func NewCreateTaskHandler(taskService *task.TaskService) *CreateTaskHandler {
	return &CreateTaskHandler{
		taskService: taskService,
	}
}

func (h *CreateTaskHandler) Handle(c *gin.Context) {
	var req reads.TaskRequestForWrite
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}

	// TODO: Replace with authenticated user ID
	userId, _ := models.ParseUserId("123e4567-e89b-12d3-a456-426614174000")

	var dueDate *time.Time
	if !req.DueDate.IsZero() {
		dueDate = &req.DueDate
	}

	cmd := task.CreateTaskParam{
		UserId:      userId,
		Title:       req.Title,
		Description: req.Description,
		DueDate:     dueDate,
	}

	ctx := c.Request.Context()
	if err := h.taskService.CreateTask(ctx, cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "Task created successfully",
	})
}
