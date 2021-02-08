package account

import (
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"hvxahv/api/client"
	pb "hvxahv/api/hvxahv/v1"
	"log"
)

// NewAccountClient 新建用户方法, 访问 accounts 微服务的客户端
func NewAccountClient(u, p string) (*pb.NewAccountReply, error) {
	port := viper.GetString("port.accounts")
	conn, err := client.Conn(port, "Accounts")
	if err != nil {
		log.Println(err)
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
