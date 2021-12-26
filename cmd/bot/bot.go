package bot

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/snowflake/v5"
	"github.com/sirupsen/logrus"
)

var log = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.InfoLevel,
}

func Start() {
	client := disgord.New(disgord.Config{
		BotToken:           os.Getenv("DISCORD_TOKEN"),
		Logger:             log,
		LoadMembersQuietly: true,
	})
	defer client.Gateway().StayConnectedUntilInterrupted()
	client.Gateway().MessageCreate(msgHandler)
}

func threadNeedsName(session disgord.Session, channelID disgord.Snowflake, usageText string) {
	_, err := session.Channel(channelID).CreateMessage(&disgord.CreateMessageParams{
		Content: fmt.Sprintf("Thread name is a required input. Usage: `%s`", usageText),
	})
	if err != nil {
		log.Error(err)
	}
}

func msgHandler(session disgord.Session, evt *disgord.MessageCreate) {
	strs := strings.Split(evt.Message.Content, " ")
	switch strs[0] {
	case "!chat":
		if len(strs[1:]) == 0 {
			threadNeedsName(session, evt.Message.ChannelID, "!chat: You need to specify the user you want to chat with.")
		} else {
			log.Info("Opening chat with user: ", strs[1])
			thread, err := session.Channel(evt.Message.ChannelID).CreateThreadNoMessage(&disgord.CreateThreadParamsNoMessage{
				Name:                "MP TRADE",
				AutoArchiveDuration: disgord.AutoArchiveThreadMinute,
				Type:                disgord.ChannelTypeGuildPublicThread,
				// Type:                disgord.ChannelTypeGuildPrivateThread,
				Invitable: true,
			})
			if err != nil {
				log.Error(err)
			}
			_ = session.Channel(thread.ID).AddThreadMember(evt.Message.Author.ID)

			userID := convertStringtoSnowflake(strs[1])

			_ = session.Channel(thread.ID).AddThreadMember(userID)
			_, err = session.Channel(thread.ID).CreateMessage(&disgord.CreateMessageParams{Content: "Lets trade!"})
			if err != nil {
				log.Error(err)
			}
		}
	case "!leave":
		log.Info(evt.Message.Author.Username, " is leaving thread")
		err := session.Channel(evt.Message.ChannelID).RemoveThreadMember(evt.Message.Author.ID)
		if err != nil {
			log.Error(errors.New("failed to leave thread"+" || "), err)
		}

	case "!help":
		_, err := session.Channel(evt.Message.ChannelID).CreateMessage(&disgord.CreateMessageParams{
			Content: "!chat: Start a new thread with a user.\n!leave: Leave the current thread.\n!help: Show this message.",
		})
		if err != nil {
			log.Error(err)
		}
	}

}

func convertStringtoSnowflake(userIDStr string) snowflake.Snowflake {
	rx := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	match := rx.FindAllString(userIDStr, -1)
	var element string
	for _, element := range match {
		log.Info(element)
	}

	number, _ := strconv.ParseUint(element, 10, 64)
	userID := snowflake.NewSnowflake(number)

	return userID
}
