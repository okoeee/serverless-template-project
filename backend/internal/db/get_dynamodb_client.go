package db

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"os"
)

func GetDynamoDBClient(ctx context.Context) (*dynamodb.Client, error) {

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	localEndpoint := os.Getenv("ENDPOINT_URL_DYNAMODB")

	svc := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		if localEndpoint != "" {
			o.BaseEndpoint = aws.String(localEndpoint)
		}
	})

	return svc, nil

}
