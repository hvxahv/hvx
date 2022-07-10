package internal

import (
	"github.com/google/uuid"
	gw "github.com/hvxahv/hvx/APIs/grpc-gateway/v1alpha1/public"
	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/public"
	v "github.com/hvxahv/hvx/microsvc"
	"github.com/pkg/errors"
)

const serviceName = "public"

type server struct {
	pb.PublicServer
}

func Run() error {
	s := v.New(
		v.WithServiceName(serviceName),
		v.WithServiceVersion("v1alpha"),
		v.WithServiceID(uuid.New().String()),
	).ListenerWithEndpoints()

	pb.RegisterPublicServer(s, &server{})
	if err := gw.RegisterPublicHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Errorf("Failed to register public services: %v", err)
	}
	if err := s.Run(); err != nil {
		return err
	}
	return nil
}
