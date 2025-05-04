//go:build integration
// +build integration

package repository

import (
	"context"
	"log"
	"os"
	"testing"

	"backend/internal/db"
	"backend/internal/models"
)

func initAndGetSUT(ctx context.Context) *UserRepository {
	os.Setenv("USERS_TABLE_NAME", "dev-users")
	os.Setenv("AWS_ACCESS_KEY_ID", "dummy")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "dummy")
	os.Setenv("AWS_REGION", "ap-northeast-1")
	os.Setenv("ENDPOINT_URL_DYNAMODB", "http://localhost:8000")

	dbClient, err := db.GetDynamoDBClient(ctx)
	if err != nil {
		log.Fatal("Error creating DynamoDB client:", err)
	}

	userRepository := NewUserRepository(dbClient)

	return userRepository
}

func TestUserRepository_CreateUser(t *testing.T) {
	ctx := context.Background()

	user := &models.User{
		UserId: "u-001",
		Name:   "testuser",
		Email:  "test@gmail.com",
	}

	userRepository := initAndGetSUT(ctx)

	err := userRepository.CreateUser(ctx, user)

	if err != nil {
		log.Fatal("Error creating user:", err)
	}
}

func TestUserRepository_GetUserById(t *testing.T) {

	ctx := context.Background()

	userId := "u-001"

	userRepository := initAndGetSUT(ctx)

	userDao, err := userRepository.GetUserById(ctx, userId)
	if err != nil {
		log.Fatal("Error getting user:", err)
	}

	log.Println("User is:", userDao)

}

func TestUserRepository_UpdateUser(t *testing.T) {

	ctx := context.Background()

	user := &models.User{
		UserId: "u-001",
		Name:   "UpdatedTest",
		Email:  "test@gmail.com",
	}

	userRepository := initAndGetSUT(ctx)

	err := userRepository.UpdateUser(ctx, user)

	if err != nil {
		log.Fatal("Error creating user:", err)
	}

}

func TestUserRepository_DeleteUser(t *testing.T) {
	ctx := context.Background()

	userId := "u-001"

	userRepository := initAndGetSUT(ctx)

	err := userRepository.DeleteUser(ctx, userId)
	if err != nil {
		log.Fatal("Error deleting user:", err)
	}

	log.Println("User deleted successfully")
}
