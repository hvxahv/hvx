/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package article

import (
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/article/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/x"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

const serviceName = "article"

type article struct {
	pb.ArticleServiceServer
	*Articles
}

// Run starts the server. It will block until the server is shutdown. If the server fails to start, it will return an error.
func Run() error {
	log.Printf("App %s Started at %s\n", serviceName, time.Now())

	// Create a new server instance.
	s := grpc.NewServer()

	pb.RegisterArticleServiceServer(s, &article{})

	reflection.Register(s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", x.NewService(serviceName).GetPort()))
	if err != nil {
		return err
	}

	go func() {
		if err := s.Serve(lis); err != nil {
			fmt.Println(err)
			return
		}
	}()

	return nil
}
