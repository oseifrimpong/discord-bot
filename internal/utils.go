package internal

import (
	"context"
	"discussion-bot/config"
	"os"
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/andersfylling/snowflake/v5"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

func ConvertStringToSnowflake(userIDStr string) snowflake.Snowflake {
	rx := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	match := rx.FindAllString(userIDStr, -1)
	var userID snowflake.Snowflake
	for _, element := range match {
		log.Info(element)
		number, _ := strconv.ParseUint(element, 10, 64)
		userID = snowflake.NewSnowflake(number)
	}
	return userID
}

func AddressChecker(walletAddress string) bool {
	dbClient := config.DynamoClient()

	out, err := dbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("DYNAMODB_TABLE")),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: walletAddress},
		},
	})
	if err != nil {
		panic(err)
	}
	if out.Item == nil {
		return false // not in the database
	}
	return true
}
