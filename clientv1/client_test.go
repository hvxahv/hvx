package clientv1

import (
	"fmt"
	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/account"
	"github.com/hvxahv/hvx/clientv1/cfg"
	"golang.org/x/net/context"
	"testing"
	"time"
)

func TestGRPC(t *testing.T) {
	ctx := context.Background()
	c, err := New(ctx,
		cfg.SetEndpoints("hvxahv.disism.internal:50010"),
		cfg.SetDialTimeout(10*time.Second),
	)
	if err != nil {
		t.Error(err)
	}
	defer c.Close()
	exist, err := c.IsExist(ctx, &pb.IsExistRequest{
		Username: "hvturingga",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(exist.IsExist)
}
