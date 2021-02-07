package social

import (
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"hvxahv/api/cli"
	pb "hvxahv/api/hvxahv/v1"
	"hvxahv/pkg/models"
	"log"
)

// OutboxAcceptClient ... 发件箱同意请求客户端
func OutboxAcceptClient(data *models.Accept) (*pb.AcceptReply, error) {
	p := viper.GetString("port.outbox")
	conn, err := cli.Conn(p, "outbox")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	cli := pb.NewOutboxClient(conn)

	d := &pb.AcceptData{
		Name: data.Name,
		Actor: data.Actor,
		RequestId: data.RequestId,
	}

	r, err := cli.Accept(context.Background(), d)
	if err != nil {
		log.Printf("接受消息失败，发送消息给 Outbox 服务端失败: %v", err)
	}
	return r, err
}

