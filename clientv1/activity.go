package clientv1

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"
	"github.com/hvxahv/hvx/APIs/v1alpha1/article"
)

type Activity interface {
	Inbox(name string, body []byte) (*pb.InboxResponse, error)
	ArticleActivity(accountId, actorId, articleId int64, article *article.CreateRequest) (*pb.ActivityResponse, error)
}

func (svc *Svc) Inbox(name string, body []byte) (*pb.InboxResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	i, err := pb.NewInboxClient(c.Conn).Inbox(svc.ctx, &pb.InboxRequest{
		Name: name,
		Data: body,
	})
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (svc *Svc) ArticleActivity(accountId, actorId, articleId int64, article *article.CreateRequest) (*pb.ActivityResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	a, err := pb.NewActivityClient(c.Conn).ArticleCreateActivity(svc.ctx, &pb.ArticleCreateActivityRequest{
		AccountId: accountId,
		ActorId:   actorId,
		ArticleId: articleId,
		Article:   article,
	})
	if err != nil {
		return nil, err
	}
	return a, nil
}
