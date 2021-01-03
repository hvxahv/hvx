package bot

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const (
	token  = "1428818014:AAEv4ZRvRR0wi4Rmu336vYO_UAtshOIjFxA"
	chatid = 441776537
)
// SendNewUserNotice 将通知发送给 Telegram 管理者
func SendNewUserNotice(name string) error {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
		return err
	}
	bot.Debug = true
	//log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	text := fmt.Sprintf("Notice... %s", name)
	msg := tgbotapi.NewMessage(chatid, text)
	bot.Send(msg)

	return nil
}
