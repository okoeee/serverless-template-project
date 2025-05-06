package task

import (
	"backend/internal/models"
	"backend/internal/service/task"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteTaskHandler struct {
	taskService *task.TaskService
}

func NewDeleteTaskHandler(taskService *task.TaskService) *DeleteTaskHandler {
	return &DeleteTaskHandler{
		taskService: taskService,
	}
}

func (h *DeleteTaskHandler) Handle(c *gin.Context) {
	taskId := c.Param("taskId")
	parsedTaskId, err := models.ParseTaskId(taskId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	ctx := c.Request.Context()
	if err := h.taskService.DeleteTask(ctx, parsedTaskId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "Task deleted successfully",
	})
}
