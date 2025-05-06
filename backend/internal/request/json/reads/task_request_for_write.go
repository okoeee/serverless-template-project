package reads

import "time"

type TaskRequestForWrite struct {
	Title       string    `json:"title" validate:"required,min=1,max=100"`
	Description string    `json:"description" validate:"max=500"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"dueDate"`
}
