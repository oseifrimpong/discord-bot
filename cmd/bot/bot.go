package bot

import (
	"discussion-bot/internal"
	"errors"
	"os"
	"strings"

	"github.com/andersfylling/disgord"
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

func msgHandler(session disgord.Session, evt *disgord.MessageCreate) {
	strs := strings.Split(evt.Message.Content, " ")
	switch strs[0] {
	case "!chat":
		if len(strs[1:]) == 0 {
			response(session, evt.Message.ChannelID, "!chat: Start a new thread with a user.\n!leave: Leave the current thread.\n")
		} else {
			err := createPrivateThread(session, evt, strs[1])
			if err != nil {
				log.Error("Creating private thread failed", err)
			}
		}
	case "!leave":
		log.Info(evt.Message.Author.Username, " is leaving thread")
		err := session.Channel(evt.Message.ChannelID).RemoveThreadMember(evt.Message.Author.ID)
		if err != nil {
			log.Error(errors.New("failed to leave thread"+" || "), err)
		}
		_, err = session.Channel(evt.Message.ChannelID).Delete()
		if err != nil {
			log.Error(errors.New("failed to delete channel"+" || "), err)
		}
	case "!WL":
		if len(strs[1:]) == 0 {
			response(session, evt.Message.ChannelID, "!WL: Check WL\n")
		} else {
		}
		log.Info("WL")
	}
}

func createPrivateThread(session disgord.Session, evt *disgord.MessageCreate, userB string) error {
	log.Info("Opening chat with user: ", userB)
	thread, err := session.Channel(evt.Message.ChannelID).CreateThreadNoMessage(&disgord.CreateThreadParamsNoMessage{
		Name:                "MP TRADE",
		AutoArchiveDuration: disgord.AutoArchiveThreadMinute,
		// Type:                disgord.ChannelTypeGuildPublicThread,
		Type:      disgord.ChannelTypeGuildPrivateThread,
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

	userID := internal.ConvertStringtoSnowflake(userB)

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

func response(session disgord.Session, channelID disgord.Snowflake, usageText string) {
	_, err := session.Channel(channelID).CreateMessage(&disgord.CreateMessageParams{
		Content: usageText,
	})
	if err != nil {
		log.Error(err)
	}
}
