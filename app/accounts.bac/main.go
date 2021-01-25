/**
	Accounts 作为整个程序的账户管理的微服务
 */
package main
//
//import (
//	"fmt"
//	"github.com/spf13/viper"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/reflection"
//	pb "hvxahv/api/kernel/v1"
//	"hvxahv/services/accounts.bac/services"
//	"hvxahv/pkg/bot"
//	"log"
//	"net"
//)
//
//func main()  {
//	services.InitDB()
//
//	viper.SetConfigFile("./configs/config.yaml")
//	err := viper.ReadInConfig()
//	if err != nil {
//		panic(fmt.Errorf("Fatal error config file: %s \n", err))
//	}
//
//	p := viper.GetString("port.services.bac")
//	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", p))
//	if err != nil {
//		fmt.Printf("Accounts gRPC Services Failed to Listen: %v", err)
//		return
//	} else {
//		// 在控制台中打印服务启动 log 通知 并 Bot ，Account 服务已经开启
//		log.Println("Accounts gRPC Services is running....", p)
//		go bot.ServicesRunningNotice("account", p)
//	}
//	s := grpc.NewServer()
//	pb.RegisterAccountsServer(s, &server{})
//	reflection.Register(s)
//
//	if err := s.Serve(lis); err != nil {
//		fmt.Printf("Accounts gRPC Services failed to start: %v", err)
//	}
//}
