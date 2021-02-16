package social

import (
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	pb "hvxahv/api/hvxahv/v1alpha1"
	"hvxahv/pkg/client"
	"hvxahv/pkg/models"
	"log"
)

// Inbox 功能的 gRPC 客户端, 它用来调用 inbox 的服务
func InboxClient(data *models.Inbox) (string, error) {
	p := viper.GetString("port.inbox")
	conn, err := client.Conn(p, "Inbox")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	cli := pb.NewInboxClient(conn)
	d := &pb.InboxData{
		Actor: data.Actor,
		RequestId: data.RequestId,
		EventType: data.EventType,
		Name: data.Name,
	}
	r, err := cli.NewInbox(context.Background(), d)
	if err != nil {
		log.Printf("接受消息失败，发送消息给 Inbox 服务端失败: %v", err)
	}
	return r.Reply, err
}

func GetInboxClient(name string) (*pb.GetInboxReply, error) {
	p := viper.GetString("port.inbox")
	conn, err := client.Conn(p, "Inbox")
	if err != nil {
		log.Println(err)
	}
	cli := pb.NewInboxClient(conn)

	r, err := cli.GetInbox(context.Background(), &pb.Name{Name: name})
	if err != nil {
		log.Printf("接受消息失败，发送消息给 Inbox 服务端失败: %v", err)
	}
	return r, err
}