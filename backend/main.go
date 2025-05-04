package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"serverless-go-react-native/backend/internal/db"
	"serverless-go-react-native/backend/internal/db/repository"
)

type Request struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type Response struct {
	Message string `json:"message"`
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	ctx := context.Background()

	fmt.Println("Hello from AWS Lambda!")
	fmt.Println("Request:", request)

	message := "Hello from AWS Lambda!"
	env := os.Getenv("APP_ENV")
	if env != "" {
		message = fmt.Sprintf("Hello from AWS Lambda! Environment: %s", env)
	}

	var requestData Request
	if err := json.Unmarshal([]byte(request.Body), &requestData); err != nil {
		errorMessage := fmt.Sprintf("Error parsing request body: %v", err)
		errorResponse := Response{
			Message: errorMessage,
		}
		jsonErrorResponse, _ := json.Marshal(errorResponse)

		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       string(jsonErrorResponse),
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
	}

	responseData := Response{
		Message: message,
	}

	// DB Client
	dbClient, err := db.GetDynamoDBClient(ctx)
	if err != nil {
		log.Fatal("Error creating DynamoDB client:", err)
	}

	userRepository := repository.NewUserRepository(dbClient)

	// Get User
	userId := "u-001"
	user, err := userRepository.GetUserById(ctx, userId)
	if err != nil {
		log.Fatal("Error getting user:", err)
	}

	log.Println("User is:", user)

	jsonResponse, err := json.Marshal(responseData)
	if err != nil {
		errorResponse := Response{
			Message: fmt.Sprintf("Error marshalling response: %v", err),
		}
		jsonErrorResponse, _ := json.Marshal(errorResponse)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       string(jsonErrorResponse),
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jsonResponse),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
