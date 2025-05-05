package repository

import (
	"backend/internal/db"
	"backend/internal/models"
	"context"
	"log"
	"os"
	"testing"
)

func initAndGetSUT(ctx context.Context) *TaskRepository {
	os.Setenv("TASKS_TABLE_NAME", "dev-tasks")
	os.Setenv("AWS_ACCESS_KEY_ID", "dummy")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "dummy")
	os.Setenv("AWS_REGION", "ap-northeast-1")
	os.Setenv("ENDPOINT_URL_DYNAMODB", "http://localhost:8000")

	dbClient, err := db.GetDynamoDBClient(ctx)
	if err != nil {
		log.Fatal("Error creating DynamoDB client:", err)
	}

	sut := NewTaskRepository(dbClient)

	return sut
}

func TestTaskRepository_CreateTask(t *testing.T) {
	ctx := context.Background()

	user := models.NewUser("Test", "test@gmail.com")

	newTask := models.NewTask(
		user.UserId,
		"Clean the house",
		"Clean the house before the party",
		nil,
	)

	sut := initAndGetSUT(ctx)

	err := sut.CreateTask(ctx, newTask)
	if err != nil {
		t.Fatal("Error creating task:", err)
	}

	tasks, err := sut.filterTaskByUserId(ctx, user.UserId)
	if err != nil {
		t.Fatal("Error filtering tasks:", err)
	}

	found := false
	for _, task := range tasks {
		if task.TaskId == newTask.TaskId {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("Task not found in the list of tasks for the user")
	}

}

func TestTaskRepository_DeleteTask(t *testing.T) {
	ctx := context.Background()

	user := models.NewUser("Test", "test@gmail.com")

	newTask := models.NewTask(
		user.UserId,
		"Clean the house",
		"Clean the house before the party",
		nil,
	)

	sut := initAndGetSUT(ctx)

	err := sut.CreateTask(ctx, newTask)
	if err != nil {
		t.Fatal("Error creating task:", err)
	}

	err = sut.DeleteTask(ctx, newTask.TaskId)
	if err != nil {
		t.Fatal("Error deleting task:", err)
	}

}

func TestTaskRepository_UpdateTask(t *testing.T) {
	ctx := context.Background()

	user := models.NewUser("Test", "test@gmail.com")

	newTask := models.NewTask(
		user.UserId,
		"Clean the house",
		"Clean the house before the party",
		nil,
	)

	sut := initAndGetSUT(ctx)

	err := sut.CreateTask(ctx, newTask)
	if err != nil {
		t.Fatal("Error creating task:", err)
	}

	newTask.Status = models.TaskStatusDone
	err = sut.UpdateTask(ctx, newTask)
	if err != nil {
		t.Fatal("Error updating task:", err)
	}

	tasks, err := sut.filterTaskByUserId(ctx, user.UserId)
	if err != nil {
		t.Fatal("Error filtering tasks:", err)
	}

	isUpdated := false
	for _, task := range tasks {
		if task.TaskId == newTask.TaskId && task.Status == models.TaskStatusDone {
			isUpdated = true
			break
		}
	}

	if !isUpdated {
		t.Fatal("Task not found in the list of tasks for the user")
	}

}
