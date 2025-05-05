package repository

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"time"

	"backend/internal/db/dao"
	"backend/internal/models"
	"os"
)

type UserRepository struct {
	Client    *dynamodb.Client
	TableName string
}

func NewUserRepository(client *dynamodb.Client) *UserRepository {
	table := os.Getenv("USERS_TABLE_NAME")
	if table == "" {
		panic("USERS_TABLE_NAME is not set")
	}

	return &UserRepository{
		Client:    client,
		TableName: table,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {

	userDao := &dao.UserDao{
		UserId:    user.UserId.String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	item, err := attributevalue.MarshalMap(userDao)
	if err != nil {
		return err
	}

	_, err = r.Client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &r.TableName,
		Item:      item,
	})
	return err
}

func (r *UserRepository) GetUserById(ctx context.Context, userId models.UserId) (*models.User, error) {
	result, err := r.Client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &r.TableName,
		Key: map[string]types.AttributeValue{
			"userId": &types.AttributeValueMemberS{Value: userId.String()},
		},
	})
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	var userDao dao.UserDao
	if err := attributevalue.UnmarshalMap(result.Item, &userDao); err != nil {
		return nil, err
	}

	userId, err = models.ParseUserId(userDao.UserId)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		UserId: userId,
		Name:   userDao.Name,
		Email:  userDao.Email,
	}

	return user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	result, err := r.Client.Query(ctx, &dynamodb.QueryInput{
		TableName:              &r.TableName,
		IndexName:              aws.String("EmailIndex"),
		KeyConditionExpression: aws.String("email = :email"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":email": &types.AttributeValueMemberS{Value: email},
		},
	})
	if err != nil {
		return nil, err
	}

	if len(result.Items) == 0 {
		return nil, nil
	}

	var userDao dao.UserDao
	if err := attributevalue.UnmarshalMap(result.Items[0], &userDao); err != nil {
		return nil, err
	}

	userId, err := models.ParseUserId(userDao.UserId)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		UserId: userId,
		Name:   userDao.Name,
		Email:  userDao.Email,
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *models.User) error {

	userDao := &dao.UserDao{
		UserId:    user.UserId.String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	item, err := attributevalue.MarshalMap(userDao)
	if err != nil {
		return err
	}

	_, err = r.Client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:           &r.TableName,
		Item:                item,
		ConditionExpression: aws.String("attribute_exists(userId)"),
	})
	return err
}

func (r *UserRepository) DeleteUser(ctx context.Context, userId models.UserId) error {

	_, err := r.Client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &r.TableName,
		Key: map[string]types.AttributeValue{
			"userId": &types.AttributeValueMemberS{Value: userId.String()},
		},
	})
	return err

}
