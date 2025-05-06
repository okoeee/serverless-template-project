package repository

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"os"
	"time"

	"backend/internal/db/dao"
	"backend/internal/models"
)

type TaskRepository struct {
	Client    *dynamodb.Client
	TableName string
}

func NewTaskRepository(client *dynamodb.Client) *TaskRepository {
	table := os.Getenv("TASKS_TABLE_NAME")
	if table == "" {
		panic("TASKS_TABLE_NAME is not set")
	}

	return &TaskRepository{
		Client:    client,
		TableName: table,
	}
}

func (r *TaskRepository) CreateTask(ctx context.Context, task *models.Task) error {

	taskDao := &dao.TaskDao{
		TaskId:      task.TaskId.String(),
		UserId:      task.UserId.String(),
		Title:       task.Title,
		Description: task.Description,
		Status:      string(task.Status),
		DueDate:     task.DueDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	item, err := attributevalue.MarshalMap(taskDao)
	if err != nil {
		return err
	}

	_, err = r.Client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &r.TableName,
		Item:      item,
	})

	return err

}

func (r *TaskRepository) FilterTaskByUserId(ctx context.Context, userId models.UserId) ([]*models.Task, error) {

	result, err := r.Client.Query(ctx, &dynamodb.QueryInput{
		TableName:              &r.TableName,
		IndexName:              aws.String("UserTasksIndex"),
		KeyConditionExpression: aws.String("userId = :userId"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":userId": &types.AttributeValueMemberS{Value: userId.String()},
		},
	})
	if err != nil {
		return nil, err
	}

	var tasksDao []*dao.TaskDao
	err = attributevalue.UnmarshalListOfMaps(result.Items, &tasksDao)
	if err != nil {
		return nil, err
	}

	tasks := make([]*models.Task, len(tasksDao))
	for i, taskDao := range tasksDao {
		taskId, err := models.ParseTaskId(taskDao.TaskId)
		if err != nil {
			return nil, err
		}

		userId, err := models.ParseUserId(taskDao.UserId)
		if err != nil {
			return nil, err
		}

		tasks[i] = &models.Task{
			TaskId:      taskId,
			UserId:      userId,
			Title:       taskDao.Title,
			Description: taskDao.Description,
			Status:      models.TaskStatus(taskDao.Status),
			DueDate:     taskDao.DueDate,
		}
	}

	return tasks, nil
}

func (r *TaskRepository) UpdateTask(ctx context.Context, task *models.Task) error {

	taskDao := &dao.TaskDao{
		TaskId:      task.TaskId.String(),
		UserId:      task.UserId.String(),
		Title:       task.Title,
		Description: task.Description,
		Status:      string(task.Status),
		DueDate:     task.DueDate,
		UpdatedAt:   time.Now(),
	}

	item, err := attributevalue.MarshalMap(taskDao)
	if err != nil {
		return err
	}

	_, err = r.Client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:           &r.TableName,
		Item:                item,
		ConditionExpression: aws.String("attribute_exists(taskId)"),
	})

	return err

}

func (r *TaskRepository) DeleteTask(ctx context.Context, taskId models.TaskId) error {

	_, err := r.Client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &r.TableName,
		Key: map[string]types.AttributeValue{
			"taskId": &types.AttributeValueMemberS{Value: taskId.String()},
		},
	})

	return err

}
