package lambda

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Hello from AWS Lambda!")
	fmt.Println("Request:", request)

	message := "Hello from AWS Lambda!"
	env := os.Getenv("APP_ENV")
	if env != "" {
		message = fmt.Sprintf("Hello from AWS Lambda! Environment: %s", env)
	}

	return events.APIGatewayProxyResponse{Body: message, StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
