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
	client.Gateway().MessageCreate(handler)
}

func handler(session disgord.Session, evt *disgord.MessageCreate) {
	strs := strings.Split(evt.Message.Content, " ")
	switch strs[0] {
	case "!chat":
		if len(strs[1:]) == 0 {
			internal.Response(session, evt.Message.ChannelID, "!chat: Start a new thread with a user.\n!leave: Leave the current thread.\n")
		} else {
			err := internal.CreatePrivateThread(session, evt, strs[1])
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
			internal.Response(session, evt.Message.ChannelID, "!WL: add wallet address to check if you are on the white list\n")
		} else {
			err := internal.WhiteList(session, evt, "bhsytehekkksije6763483838923899237jjdg")
			if err != nil {
				log.Error(errors.New("failed to check whiteList "+" || "), err)
			}
		}
		log.Info("WL")
	}
}
