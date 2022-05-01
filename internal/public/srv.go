package public

import (
	"github.com/google/uuid"
	pb "github.com/hvxahv/hvx/api/grpc/proto/public/v1alpha1"
	"github.com/hvxahv/hvx/pkg/v"
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
	).NewServer()

	pb.RegisterPublicServer(s, &server{})
	if err := pb.RegisterPublicHandler(s.GetCtx(), s.GetMux(), s.GetConn()); err != nil {
		return errors.Errorf("Failed to register public services: %v", err)
	}
	if err := s.Run(); err != nil {
		return err
	}
	return nil
}
