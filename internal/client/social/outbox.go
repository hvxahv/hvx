package social

import (
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	pb "hvxahv/api/hvxahv/v1alpha1"
	"hvxahv/internal/client"
	"log"
)

// OutboxAcceptClient ... 发件箱同意请求客户端
func OutboxAcceptClient(data *inbox.Accept) (*pb.ReplyCode, error) {
	p := viper.GetString("port.outbox")
	conn, err := client.Conn(p, "outbox")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	cli := pb.NewOutboxClient(conn)

	d := &pb.AcceptData{
		Name:      data.Name,
		Actor:     data.Actor,
		RequestId: data.RequestId,
	}

	r, err := cli.Accept(context.Background(), d)
	if err != nil {
		log.Printf("接受消息失败，发送消息给 Outbox 服务端失败: %v", err)
	}
	return r, err
}

// OutboxFollowClient ... 发送请求关注的客户端, 它接收两个参数 name 当前的用户名, actor 请求关注的人的 URL
func OutboxFollowClient(name, actor string) (*pb.ReplyCode, error) {
	p := viper.GetString("port.outbox")
	conn, err := client.Conn(p, "outbox")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	cli := pb.NewOutboxClient(conn)

	d := &pb.FollowData{
		Name:  name,
		Actor: actor,
	}

	r, err := cli.Follow(context.Background(), d)
	if err != nil {
		log.Printf("接受消息失败，发送消息给 Outbox 服务端失败: %v", err)
	}
	return r, err
}
