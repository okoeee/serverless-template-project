package dao

import (
	"time"
)

type TaskDao struct {
	TaskId      string    `dynamodbav:"taskId"`
	UserId      string    `dynamodbav:"userId"`
	Title       string    `dynamodbav:"title"`
	Description string    `dynamodbav:"description"`
	Status      string    `dynamodbav:"status"`
	DueDate     *time.Time `dynamodbav:"dueDate"`
	CreatedAt   time.Time `dynamodbav:"createdAt"`
	UpdatedAt   time.Time `dynamodbav:"updatedAt"`
}
