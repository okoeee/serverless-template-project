package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Hello from AWS Lambda!")
	fmt.Println("Request:", request)

	return events.APIGatewayProxyResponse{Body: "Hello from AWS Lambda!", StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
