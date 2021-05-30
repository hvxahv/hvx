package main

import (
	"golang.org/x/net/context"
	bot2 "hvxahv/api/bot"
	"hvxahv/pkg/bot"
)

// Notice // 接到消息后将通过 bot pkg 的方法来发送给 TG Bot
func (s *server) Notice(ctx context.Context, in *bot2.BotNoticeSend) (*bot2.BotNoticeReply, error) {
	bot.NewAccountNotice(in.Message)
	return &bot2.BotNoticeReply{Reply: "Send bot notification successful"}, nil
}
