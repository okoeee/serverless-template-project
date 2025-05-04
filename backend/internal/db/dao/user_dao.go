package dao

import (
	"time"
)

type UserDao struct {
	UserId    string    `dynamodbav:"userId"`
	Name      string    `dynamodbav:"name"`
	Email     string    `dynamodbav:"email"`
	CreatedAt time.Time `dynamodbav:"createdAt"`
	UpdatedAt time.Time `dynamodbav:"updatedAt"`
}
