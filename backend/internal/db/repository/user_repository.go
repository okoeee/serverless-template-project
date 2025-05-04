package repository

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"os"
	"serverless-go-react-native/backend/internal/db/dao"
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

func (r *UserRepository) CreateUser(ctx context.Context, user *dao.UserDao) error {
	item, err := attributevalue.MarshalMap(user)
	if err != nil {
		return err
	}

	_, err = r.Client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &r.TableName,
		Item:      item,
	})
	return err
}

func (r *UserRepository) GetUserById(ctx context.Context, userId string) (*dao.UserDao, error) {
	result, err := r.Client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &r.TableName,
		Key: map[string]types.AttributeValue{
			"userId": &types.AttributeValueMemberS{Value: userId},
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

	return &userDao, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *dao.UserDao) error {
	item, err := attributevalue.MarshalMap(user)
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

func (r *UserRepository) DeleteUser(ctx context.Context, userId string) error {

	_, err := r.Client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &r.TableName,
		Key: map[string]types.AttributeValue{
			"userId": &types.AttributeValueMemberS{Value: userId},
		},
	})
	return err

}
