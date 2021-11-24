package client

import (
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/channel/v1alpha1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
)

// Channel Microservice client for Channel.
func Channel() (pb.ChannelClient, *grpc.ClientConn, error) {
	name := "channels"

	host := viper.GetString(fmt.Sprintf("microservices.%s.host", name))
	port := viper.GetString(fmt.Sprintf("microservices.%s.port", name))
	addr := fmt.Sprintf("%s:%s", host, port)
	fmt.Println(addr)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("failed to connect to %s service: %v", name , err)
		return nil, nil, nil
	}

	cli := pb.NewChannelClient(conn)

	return cli, conn, nil
}
