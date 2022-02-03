package device

import (
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/device/v1alpha1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func NewDeviceClient() (pb.DevicesClient, error) {
	address := fmt.Sprintf("%s:%s", viper.GetString("microservices.device.host"), viper.GetString("microservices.device.port"))

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return pb.NewDevicesClient(conn), nil
}
