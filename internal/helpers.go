package internal

import (
	"github.com/andersfylling/disgord"

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

func WhiteList(session disgord.Session, evt *disgord.MessageCreate, walletAddress string) error {
	log.Info("Checking whitelist for %s", evt.Message.Author.Username)
	if AddressChecker(walletAddress) {
		// userID := strconv.Itoa(int(evt.Message.Author.ID))
		_, err := session.Channel(evt.Message.ChannelID).CreateMessage(&disgord.CreateMessageParams{Content: evt.Message.Author.Username + ", you are on the whitelist!"})
		if err != nil {
			log.Error("Error creating message: ", err)
			return err
		}
	} else {
		_, err := session.Channel(evt.Message.ChannelID).CreateMessage(&disgord.CreateMessageParams{Content: evt.Message.Author.Username + "! Sorry, looks like you are not on the whitelist! If you think this is an error, please create a <#941680942995623966>"})
		if err != nil {
			log.Error("Error creating message: ", err)
			return err
		}
	}
	return nil
}
