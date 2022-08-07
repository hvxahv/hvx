package clientv1

import (
	"fmt"
	"testing"
	"time"

	pb "github.com/hvxahv/hvx/APIs/v1alpha1/account"
	"golang.org/x/net/context"
)

func TestGRPC(t *testing.T) {
	ctx := context.Background()
	c, err := New(ctx,
		cfg.SetEndpoints("hvxahv.disism.internal:50010"),
		SetDialTimeout(10*time.Second),
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
