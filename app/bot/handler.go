package main

import (
	"golang.org/x/net/context"
	v1 "hvxahv/api/util/v1"
	"hvxahv/pkg/bot"
)

// Notice // 接到消息后将通过 bot pkg 的方法来发送给 TG Bot
func (s *server) Notice(ctx context.Context, in *v1.BotNoticeSend) (*v1.BotNoticeReply, error)  {
	bot.NewAccountNotice(in.Message)
	return &v1.BotNoticeReply{Reply: "Send bot notification successful"}, nil
}
