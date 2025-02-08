package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"youricktube/package/ricktuber"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TG_BOT_TOKEN")
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	rtuber := ricktuber.Ricktuber

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			msg := ""
			for _, videoId := range rtuber.FindYoutubeVideos(update.Message.Text) {
				msg += rtuber.GetLink(videoId) + "\n"
			}

			if msg != "" {
				api.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg))
			}
		}
	}
}
