package bot

import (
	"fmt"
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

func threadNeedsName(session disgord.Session, channelID disgord.Snowflake, usageText string) {
	_, err := session.Channel(channelID).CreateMessage(&disgord.CreateMessageParams{
		Content: fmt.Sprintf("Thread name is a required input. Usage: `%s`", usageText),
	})
	if err != nil {
		log.Error(err)
	}
	// _, err := session.Channel(channelID).)
}

func Start() {
	client := disgord.New(disgord.Config{
		BotToken: os.Getenv("DISCORD_TOKEN"),
		Logger:   log,

		RejectEvents: []string{
			disgord.EvtThreadCreate,
			disgord.EvtThreadUpdate,
			disgord.EvtThreadDelete,
			disgord.EvtThreadMembersUpdate,
			disgord.EvtThreadListSync,
			disgord.EvtThreadListSync,
		},
	})
	defer client.Gateway().StayConnectedUntilInterrupted()
	client.Gateway().MessageCreate(msgHandler)
}

func msgHandler(session disgord.Session, evt *disgord.MessageCreate) {
	strs := strings.Split(evt.Message.Content, " ")
	switch strs[0] {
	case "!chat":
		if len(strs[1:]) == 0 {
			threadNeedsName(session, evt.Message.ChannelID, "!chat need a name")
		} else {
			threadName := strs[1]
			thread, err := session.Channel(evt.Message.ChannelID).CreateThread(evt.Message.ID, &disgord.CreateThreadParams{
				Name: threadName,
				// any auto archive thread duration greater than AutoArchiveThreadDay requires premium
				AutoArchiveDuration: disgord.AutoArchiveThreadDay,
			})
			if err != nil {
				log.Error(err)
			}
			// send a message in the newly created thread
			_, err = session.Channel(thread.ID).CreateMessage(&disgord.CreateMessageParams{Content: "HELLO WORLD"})
			if err != nil {
				log.Error(err)
			}
		}
	case "$makethreadnomessage":
		if len(strs[1:]) == 0 {
			threadNeedsName(session, evt.Message.ChannelID, "$makethreadnomessage my-awesome-thread-name")
		} else {
			threadName := strs[1]
			thread, err := session.Channel(evt.Message.ChannelID).CreateThreadNoMessage(&disgord.CreateThreadParamsNoMessage{
				Name: threadName,
				// any auto archive thread duration greater than AutoArchiveThreadDay requires premium
				AutoArchiveDuration: disgord.AutoArchiveThreadDay,
				Type:                disgord.ChannelTypeGuildPublicThread,
			})
			if err != nil {
				log.Error(err)
			}
			// send a message in the newly created thread
			_, err = session.Channel(thread.ID).CreateMessage(&disgord.CreateMessageParams{Content: "HELLO WORLD"})
			if err != nil {
				log.Error(err)
			}
		}
	}
}
