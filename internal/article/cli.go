/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package article

import (
	pb "github.com/hvxahv/hvxahv/api/article/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/x"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetArticleClient() (pb.ArticleServiceClient, error) {
	conn, err := grpc.Dial(x.NewService(serviceName).GetAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewArticleServiceClient(conn), nil
}
