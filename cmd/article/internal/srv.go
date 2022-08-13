package internal

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/article"
	svc "github.com/hvxahv/hvx/microsvc"
	"github.com/pkg/errors"
)

type server struct {
	pb.ArticleServer
}

const (
	serviceName = "article"
)

func Run() error {
	s := svc.New(
		svc.WithServiceName(serviceName),
		svc.WithServiceVersion("v1alpha"),
		svc.WithServiceID("serviceName"),
	).ListenerWithEndpoints()

	pb.RegisterArticleServer(s, &server{})

	if err := pb.RegisterArticleHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Errorf("Failed to register %s services: %v", serviceName, err)
	}

	if err := s.Run(); err != nil {
		return err
	}

	return nil
}
