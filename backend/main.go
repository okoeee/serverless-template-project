package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"backend/internal/handlers"
)

func main() {
	lambda.Start(handlers.SampleRequestHandler)
}
