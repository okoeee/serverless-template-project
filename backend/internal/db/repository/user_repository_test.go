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

	user := models.NewUser(
		"Test",
		"test@gmail.com",
	)

	userRepository := initAndGetSUT(ctx)

	err := userRepository.CreateUser(ctx, user)

	if err != nil {
		log.Fatal("Error creating user:", err)
	}
}

func TestUserRepository_GetUserById(t *testing.T) {

	ctx := context.Background()

	userRepository := initAndGetSUT(ctx)

	user, err := userRepository.GetUserByEmail(ctx, "test@gmail.com")
	if err != nil {
		log.Fatal("Error getting user:", err)
	}

	userDao, err := userRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		log.Fatal("Error getting user:", err)
	}

	log.Println("User is:", userDao)

}

func TestUserRepository_UpdateUser(t *testing.T) {

	ctx := context.Background()

	userRepository := initAndGetSUT(ctx)

	user, err := userRepository.GetUserByEmail(ctx, "test@gmail.com")
	if err != nil {
		log.Fatal("Error getting user:", err)
	}

	user.Name = "Updated"

	err = userRepository.UpdateUser(ctx, user)

	if err != nil {
		log.Fatal("Error creating user:", err)
	}

}

func TestUserRepository_DeleteUser(t *testing.T) {
	ctx := context.Background()

	userRepository := initAndGetSUT(ctx)

	user, err := userRepository.GetUserByEmail(ctx, "test@gmail.com")
	if err != nil {
		log.Fatal("Error getting user:", err)
	}

	err = userRepository.DeleteUser(ctx, user.UserId)
	if err != nil {
		log.Fatal("Error deleting user:", err)
	}

	log.Println("User deleted successfully")
}
