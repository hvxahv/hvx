package cli

import (
	article "github.com/hvxahv/hvx/api/grpc/proto/article/v1alpha1"
	"google.golang.org/grpc"
)

type Article interface {
	article.ArticleClient
}

type art struct {
	article.ArticleClient
}

func NewArticle(conn *grpc.ClientConn) Article {
	return &art{
		ArticleClient: article.NewArticleClient(conn),
	}
}
