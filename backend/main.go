package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"backend/internal/handlers"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	lambda.Start(handlers.SampleRequestHandler)
}
