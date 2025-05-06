package writes

import (
	"backend/internal/models"
	"time"
)

type TaskResponse struct {
	TaskId      string            `json:"taskId"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      models.TaskStatus `json:"status"`
	DueDate     *time.Time        `json:"dueDate"`
}
