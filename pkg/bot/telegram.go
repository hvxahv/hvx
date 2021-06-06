package bot

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func TG(token string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot

}