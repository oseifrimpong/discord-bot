package main

import (
	"discussion-bot/bot"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot.Start()

	<-make(chan struct{})
	return
}
