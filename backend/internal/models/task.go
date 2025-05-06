package models

import (
	"errors"
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

func ParseTaskStatus(status string) (TaskStatus, error) {
	switch status {
	case string(TaskStatusTodo):
		return TaskStatusTodo, nil
	case string(TaskStatusInProgress):
		return TaskStatusInProgress, nil
	case string(TaskStatusDone):
		return TaskStatusDone, nil
	default:
		return "", errors.New("invalid task status")
	}
}

type Task struct {
	TaskId      TaskId
	UserId      UserId
	Title       string
	Description string
	Status      TaskStatus
	DueDate     *time.Time
}

func NewTask(userId UserId, title, description string, dueDate *time.Time) (*Task, error) {

	if err := validateTitle(title); err != nil {
		return nil, err
	}

	task := Task{
		TaskId:      newTaskId(),
		UserId:      userId,
		Title:       title,
		Description: description,
		Status:      TaskStatusTodo,
		DueDate:     dueDate,
	}
	return &task, nil

}

func (t *Task) Update(title, description string, status TaskStatus, dueDate *time.Time) error {

	if err := validateTitle(title); err != nil {
		return err
	}

	t.Title = title
	t.Description = description
	t.Status = status
	t.DueDate = dueDate

	return nil

}

func validateTitle(title string) error {
	if title == "" {
		return errors.New("title is required")
	}
	if len(title) > 100 {
		return errors.New("title must be less than 100 characters")
	}
	return nil
}
