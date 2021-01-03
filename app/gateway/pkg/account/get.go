package account

import (
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "hvxahv/api/kernel/v1"
	"log"
)

// GetAccountsHandler 获取用户的个人资料
func GetAccountsClient(author string)  {

	// 链接到 Accounts 服务端
	conn, err := grpc.Dial(viper.GetString("port.accounts"), grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Faild to connect to Accounts app: %v", err)
	}
	defer conn.Close()

	client := pb.NewAccountsClient(conn)
	if err != nil {
		log.Println(err)
	}
	r, err := client.GetAccount(context.Background(), &pb.AccountName{
		Username: author,
	})
	if err != nil {
		log.Printf("发送消息给 Accounts 服务端失败: %v", err)
	}
	log.Println(r)
}
