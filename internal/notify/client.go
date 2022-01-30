package notify

import (
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/notify/v1alpha1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func NewNotifyClient() (pb.NotifyClient, error) {
	address := fmt.Sprintf("%s:%s", viper.GetString("microservices.notify.localhost"), viper.GetString("microservices.notify.port"))

	fmt.Println(address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return pb.NewNotifyClient(conn), nil
}
