package main

import (
	"backend/internal/db"
	"backend/internal/db/repository"
	"backend/internal/handlers/task"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambdaV2

func init() {

	dynamoDBClient, err := db.GetDynamoDBClient(context.Background())
	if err != nil {
		panic("Failed to create DynamoDB client: " + err.Error())
	}

	listTaskHandler := task.ListTaskHandler{TaskRepository: repository.NewTaskRepository(dynamoDBClient)}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api/v1")
	{
		taskGroup := v1.Group("/task")
		{
			taskGroup.GET("", listTaskHandler.Handle)
		}
	}

	ginLambda = ginadapter.NewV2(r)
}

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	lambda.Start(Handler)
}
