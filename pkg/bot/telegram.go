package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

type bot struct {
	bot *tgbotapi.BotAPI
	level   string
	message string
}

// NewBot The constructor of the Bot package will accept the level and message string.
// Message level
// 1: Notice
// 2: Normal
// 3: Error
func NewBot(level int, message string) TG {
	b, err := initBot()
	if err != nil {
		log.Println(err)
	}
	var l string
	switch level {
	case 1:
		l = "Notice"
	case 2:
		l = "Normal"
	case 3:
		l = "Error"

	}
	return &bot{bot: b, level: l, message: message}
}

type TG interface {
	Send() error
	GetUpdateId() error
}

func (b *bot) Send() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	m := fmt.Sprintf("%s:%s", b.level, b.message)

	id, _ := strconv.ParseInt(viper.GetString("bot.tg_dev_id"), 10, 64)
	msg := tgbotapi.NewMessage(id, m)

	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (b *bot) GetUpdateId() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		m := fmt.Sprintf("%s:%s", b.level, b.message)
		log.Println(m)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, m)
		msg.ReplyToMessageID = update.Message.MessageID

		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}

	}
	return nil
}

// initBot Initialize Telegram Bot.
func initBot() (*tgbotapi.BotAPI, error) {
	// token
	t := viper.GetString("bot.tg_token")
	bot, err := tgbotapi.NewBotAPI(t)
	if err != nil {
		return nil, errors.Errorf("Failed to initialize Telegram Bot: %v", err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot, nil
}
