package clientv1

import (
	pb "github.com/hvxahv/hvx/api/grpc/proto/article/v1alpha1"
)

type Article interface {
	pb.ArticleClient
}

type article struct {
	pb.ArticleClient
}

func NewArticle(c *Client) Article {
	return &article{
		ArticleClient: pb.NewArticleClient(c.conn),
	}
}
