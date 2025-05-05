package models

import (
	"time"

	"github.com/google/uuid"
)

type TaskId uuid.UUID

func newTaskId() TaskId {
	return TaskId(uuid.New())
}

func (id TaskId) String() string {
	return uuid.UUID(id).String()
}

func ParseTaskId(id string) (TaskId, error) {
	taskId, err := uuid.Parse(id)
	if err != nil {
		return TaskId{}, err
	}
	return TaskId(taskId), nil
}

type TaskStatus string

const (
	TaskStatusTodo       TaskStatus = "TODO"
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	TaskStatusDone       TaskStatus = "DONE"
)

type Task struct {
	TaskId      TaskId
	UserId      UserId
	Title       string
	Description string
	Status      TaskStatus
	DueDate     *time.Time
}

func NewTask(userId UserId, title, description string, dueDate *time.Time) *Task {
	task := Task{
		TaskId:      newTaskId(),
		UserId:      userId,
		Title:       title,
		Description: description,
		Status:      TaskStatusTodo,
		DueDate:     dueDate,
	}
	return &task
}
