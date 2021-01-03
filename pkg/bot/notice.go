package bot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func NewUserNoticeHandler(c *gin.Context)  {
	n := c.PostForm("username")
	SendNewUserNotice(n)
}


func ServicesRunningNotice(srvname string, port string) error {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
		return err
	}
	bot.Debug = true
	//log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	text := fmt.Sprintf("%s services is running..., port: %s", srvname, port)
	msg := tgbotapi.NewMessage(chatid, text)
	bot.Send(msg)

	return nil

}