package article

import (
	pb "github.com/hvxahv/hvx/APIs/grpc-go/article/v1alpha1"
	"github.com/hvxahv/hvx/clientv1"
)

type Article interface {
	pb.ArticleClient
}

type article struct {
	pb.ArticleClient
}

func NewArticle(c *clientv1.Client) Article {
	return &article{
		ArticleClient: pb.NewArticleClient(c.conn),
	}
}
