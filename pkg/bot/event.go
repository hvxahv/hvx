package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

// EventHandler ... Telegram Bot -> Event Handler .
func EventHandler() {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		i := update.Message.Text
		if i == "reg" {
			u := unr(bot, update, "开始注册你的 Godis ID. 输入你的用户名")
			log.Println(u)
		}
	}
}

func unr(bot *tgbotapi.BotAPI, update tgbotapi.Update, msg string) tgbotapi.MessageConfig {
	r := tgbotapi.NewMessage(update.Message.Chat.ID, msg)
	bot.Send(r)
	return r
}
