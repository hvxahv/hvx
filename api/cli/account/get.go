package account

import (
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"hvxahv/api/cli"
	pb "hvxahv/api/hvxahv/v1"
	"log"
	"strings"
)

// GetAccountsHandler 获取用户的个人资料
func GetAccountsClient(name string) (*pb.AccountData, error) {
	// 连接到 Accounts 服务端，并返回用户的个人数据
	addr := fmt.Sprintf("localhost:%s", viper.GetString("port.accounts"))
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Faild to connect to Accounts services.bac: %v", err)
	}
	defer conn.Close()

	client := pb.NewAccountsClient(conn)
	if err != nil {
		log.Println(err)
	}
	r, err := client.GetAccount(context.Background(), &pb.AccountName{
		Username: name,
	})
	if err != nil {
		log.Printf("获取 - %s - 账户时发送消息给 Accounts 服务端失败: %v", name, err)
	}
	return r, err

}

// GetActorClient Activitypub 协议的 Actor
func GetActorClient(name string) (*pb.AccountData, error) {
	p := viper.GetString("port.accounts")
	conn, err := cli.Conn(p, "Accounts")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	client := pb.NewAccountsClient(conn)
	r, err := client.GetActor(context.Background(), &pb.AccountName{
		Username: name,
	})
	if err != nil {
		log.Printf("获取 - %s - actor 时发送消息给 Accounts 服务端失败: %v", name, err)
	}
	return r, err

}

// GetWebFinger 获取  Activitypub 协议的 GetWebFinger
func GetWebFingerClient(name string) (*pb.AccountData, error) {
	// 将 url 传过来的数据进行过滤，得到真正的用户名
	if strings.HasPrefix(name, "acct:") {
		name = name[5:]
	}
	ali := strings.IndexByte(name, '@')
	if ali != -1 {
		name = name[:ali]
	}

	r, err := GetActorClient(name)
	if err != nil {
		log.Println(err)
	} else {
		return r, err
	}
	return nil, nil
}

// VerifyAccountsClient 获取用户的个人资料
func VerifyAccountsClient(name string) (*pb.AccountData, error) {
	p := viper.GetString("port.accounts")
	conn, err := cli.Conn(p, "Accounts")
	if err != nil {
		log.Println(err)
	}

	client := pb.NewAccountsClient(conn)
	r, err := client.VerifyAccount(context.Background(), &pb.AccountName{
		Username: name,
	})
	if err != nil {
		log.Printf("查询不到 %v, %s", name, err)
	}
	return r, err

}