/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package article

import (
	pb "github.com/hvxahv/hvxahv/api/article/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/microservices"
	"google.golang.org/grpc"
)

func GetArticleClient() (pb.ArticleServiceClient, error) {
	conn, err := grpc.Dial(microservices.NewService(serviceName).GetAddress(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return pb.NewArticleServiceClient(conn), nil
}
