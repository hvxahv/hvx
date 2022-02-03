package saved

import (
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/saved/v1alpha1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func NewSavedClient() (pb.SavedClient, error) {
	address := fmt.Sprintf("%s:%s", viper.GetString("microservices.saved.host"), viper.GetString("microservices.saved.port"))

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return pb.NewSavedClient(conn), nil
}
