package config

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func DynamoClient() *dynamodb.Client {
	AWS_ACCESS_KEY_ID := os.Getenv("AWS_ACCESS_KEY_ID")
	AWS_SECRET_ACCESS_KEY := os.Getenv("AWS_SECRET_ACCESS_KEY")
	staticProvider := credentials.NewStaticCredentialsProvider(
		AWS_ACCESS_KEY_ID,
		AWS_SECRET_ACCESS_KEY,
		"",
	)

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(staticProvider))
	cfg.Region = os.Getenv("AWS_REGION")

	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)
	return svc
}
