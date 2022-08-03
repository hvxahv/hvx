/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package internal

import (
	"github.com/google/uuid"
	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/article"
	v "github.com/hvxahv/hvx/microsvc"
	"github.com/pkg/errors"
)

const (
	serviceName = "article"
)

type server struct {
	pb.ArticleServer
}

func Run() error {
	s := v.New(
		v.WithServiceName(serviceName),
		v.WithServiceVersion("v1alpha1"),
		v.WithServiceID(uuid.New().String()),
	).ListenerWithEndpoints()

	pb.RegisterArticleServer(s, &server{})

	if err := s.Run(); err != nil {
		return err
	}

	if err := pb.RegisterArticleHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Errorf("Failed to register %s services: %v", serviceName, err)
	}
	return nil
}
