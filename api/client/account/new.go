package account

import (
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "hvxahv/api/hvxahv/v1"
	"log"
)

// NewAccountClient 新建用户方法, 访问 accounts 微服务的客户端
func NewAccountClient(u, p string) (*pb.NewAccountReply, error) {
	addr := fmt.Sprintf("localhost:%s", viper.GetString("port.accounts"))
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("Faild to connect to Accounts services.bac: %v", err)
	}
	defer conn.Close()

	client := pb.NewAccountsClient(conn)
	r, err := client.NewAccount(context.Background(), &pb.AccountData{
		Username: u, Password: p, //Avatar: "",
		//Bio: "", Name: "", Email: "",
		//Phone: "", Telegram: "", Social: "",
	})
	if err != nil {
		log.Printf("发送消息给 Accounts 服务端失败: %v", err)
	}
	return r, nil
}
