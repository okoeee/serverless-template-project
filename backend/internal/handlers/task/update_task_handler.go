package task

import (
	"backend/internal/models"
	"backend/internal/request/json/reads"
	"backend/internal/service/task"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UpdateTaskHandler struct {
	taskService *task.TaskService
}

func NewUpdateTaskHandler(taskService *task.TaskService) *UpdateTaskHandler {
	return &UpdateTaskHandler{
		taskService: taskService,
	}
}

func (h *UpdateTaskHandler) Handle(c *gin.Context) {
	var req reads.TaskRequestForWrite
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}

	taskId := c.Param("taskId")
	parsedTaskId, err := models.ParseTaskId(taskId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	// TODO: Replace with authenticated user ID
	userId, _ := models.ParseUserId("123e4567-e89b-12d3-a456-426614174000")

	var dueDate *time.Time
	if !req.DueDate.IsZero() {
		dueDate = &req.DueDate
	}

	status, err := models.ParseTaskStatus(req.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task status"})
		return
	}

	cmd := task.UpdateTaskParam{
		TaskId:      parsedTaskId,
		UserId:      userId,
		Title:       req.Title,
		Description: req.Description,
		DueDate:     dueDate,
		Status:      status,
	}

	ctx := c.Request.Context()
	if err := h.taskService.UpdateTask(ctx, cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "Task updated successfully",
	})
}
