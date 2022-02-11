package internal

import (
	"context"
	"discussion-bot/config"
	"fmt"
	"os"

	"github.com/andersfylling/disgord"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	log "github.com/sirupsen/logrus"
)

func Response(session disgord.Session, channelID disgord.Snowflake, usageText string) {
	_, err := session.Channel(channelID).CreateMessage(&disgord.CreateMessageParams{
		Content: usageText,
	})
	if err != nil {
		log.Error(err)
	}
}

func CreatePrivateThread(session disgord.Session, evt *disgord.MessageCreate, userB string) error {
	log.Info("Opening chat with user: ", userB)
	thread, err := session.Channel(evt.Message.ChannelID).CreateThreadNoMessage(&disgord.CreateThreadParamsNoMessage{
		Name:                "MP TRADE",
		AutoArchiveDuration: disgord.AutoArchiveThreadMinute,
		Type:                disgord.ChannelTypeGuildPublicThread,
		// Type:      disgord.ChannelTypeGuildPrivateThread,
		Invitable: true,
	})
	if err != nil {
		log.Error("Error creating thread: ", err)
		return err
	}

	err = session.Channel(thread.ID).AddThreadMember(evt.Message.Author.ID)
	if err != nil {
		log.Error("Error adding user A to thread: ", err)
		return err
	}

	userID := ConvertStringToSnowflake(userB)

	err = session.Channel(thread.ID).AddThreadMember(userID)
	if err != nil {
		log.Error("Error adding user B to thread: ", err)
		return err
	}

	_, err = session.Channel(thread.ID).CreateMessage(&disgord.CreateMessageParams{Content: "Lets trade!"})
	if err != nil {
		log.Error("Error creating message: ", err)
		return err
	}

	return nil
}

func WhiteList(session disgord.Session, evt *disgord.MessageCreate, userB string) error {
	log.Info("Checking white list")
	if addressChecker(userB) {
		_, err := session.Channel(evt.Message.ChannelID).CreateMessage(&disgord.CreateMessageParams{Content: "You are on the white list!"})
		if err != nil {
			log.Error("Error creating message: ", err)
			return err
		}
	} else {
		_, err := session.Channel(evt.Message.ChannelID).CreateMessage(&disgord.CreateMessageParams{Content: "You are not on the white list!"})
		if err != nil {
			log.Error("Error creating message: ", err)
			return err
		}
	}
	return nil
}

func addressChecker(walletAddress string) bool {
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
	fmt.Println(out.Item)
	return true
}
