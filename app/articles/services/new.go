package services

import (
	pb "hvxahv/api/hvxahv/v1alpha1"
)

// CreateArticleHandler Articles 微服务服务端创建文章的 Handler
// 将数据库执行的结果发送给客户端，返回 string 类型的 error 或者 ok
func CreateArticleHandler(in *pb.ArticleData) string {

	go SendActivity(in)
	return "ok"
}
