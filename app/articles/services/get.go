package services

import (
	pb "hvxahv/api/hvxahv/v1alpha1"
	"hvxahv/pkg/maria"
)

// ShowStatusLis ...
func GetArticleHandler(in *pb.GetArticleData) ([]*pb.ArticleData, string) {
	db := maria.GetMaria()
	var a []*pb.ArticleData
	if db.Debug().Table("articles").Where("author = ?", in.Name).Find(&a).RecordNotFound() {
		return nil, "没有通过用户名获取到文章"
	}
	return a, "ok"


}