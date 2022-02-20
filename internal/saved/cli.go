package saved

import (
	pb "github.com/hvxahv/hvxahv/api/saved/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/microservices"
	"google.golang.org/grpc"
)

func NewSavedClient() (pb.SavedClient, error) {
	conn, err := grpc.Dial(microservices.GetSavedAddress(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return pb.NewSavedClient(conn), nil
}
