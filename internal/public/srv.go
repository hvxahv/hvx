package public

import (
	"github.com/google/uuid"
	pb "github.com/hvxahv/hvxahv/api/v1alpha1/proto/public/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/x"
	"github.com/pkg/errors"
)

const serviceName = "public"

type server struct {
	pb.PublicServiceServer
}

func Run() error {
	s := x.New(
		x.WithServiceName(serviceName),
		x.WithServiceVersion("v1alpha1"),
		x.WithServiceID(uuid.New().String()),
	).NewServer()

	pb.RegisterPublicServiceServer(s, &server{})
	if err := pb.RegisterPublicServiceHandler(s.GetCtx(), s.GetMux(), s.GetConn()); err != nil {
		return errors.Errorf("Failed to register gateway: %v", err)
	}
	if err := s.Run(); err != nil {
		return err
	}
	return nil
}
