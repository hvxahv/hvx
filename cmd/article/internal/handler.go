package internal

import (
	"context"
	"fmt"
	"github.com/hvxahv/hvx/APIs/v1alpha1/actor"
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/article"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/microsvc"
	"google.golang.org/protobuf/types/known/emptypb"
	"strconv"
)

func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	status := func() bool {
		if in.GetSummary() != "" && in.GetTitle() != "" {
			return true
		}
		return false
	}

	a := &Articles{
		ActorId:     parse.ActorId,
		Title:       in.GetTitle(),
		Summary:     in.Summary,
		Article:     in.Article,
		Tags:        in.GetTags(),
		Statuses:    status(),
		NSFW:        in.GetNsfw(),
		Attachments: in.GetAttachments(),
		TO:          in.GetTo(),
		CC:          in.GetCc(),
		BTO:         in.Bto,
		Audience:    in.Audience,
		Visibility:  in.GetVisibility(),
	}
	create, err := a.Create()
	if err != nil {
		return nil, err
	}

	fmt.Println(create)
	return &pb.CreateResponse{
		Code:   "200",
		Status: "ok",
	}, nil
}

func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	id, err := strconv.Atoi(in.GetArticleId())
	if err != nil {
		return nil, err
	}

	a, err := NewArticlesId(uint(id)).Get()
	if err != nil {
		return nil, err
	}

	ad, err := clientv1.New(ctx, microsvc.ActorServiceName).GetActor(int64(a.ActorId))
	if err != nil {
		return nil, err
	}
	return &pb.GetResponse{
		Code: "200",
		Actor: &actor.ActorData{
			Id:                ad.Actor.Id,
			PreferredUsername: ad.Actor.PreferredUsername,
			Domain:            ad.Actor.Domain,
			Avatar:            ad.Actor.Avatar,
			Name:              ad.Actor.Name,
			Summary:           ad.Actor.Summary,
			Inbox:             ad.Actor.Inbox,
			Address:           ad.Actor.Address,
			PublicKey:         ad.Actor.PublicKey,
			ActorType:         ad.Actor.ActorType,
			IsRemote:          ad.Actor.IsRemote,
		},
		Article: &pb.ArticleInfo{
			Title:       a.Title,
			Summary:     a.Summary,
			Article:     a.Article,
			Tags:        a.Tags,
			Nsfw:        a.NSFW,
			To:          a.TO,
			Cc:          a.CC,
			Bto:         a.BTO,
			Audience:    a.Audience,
			Attachments: a.Attachments,
			Visibility:  a.Visibility,
			Id:          int64(a.ID),
			CreateAt:    a.CreatedAt.String(),
			UpdateAt:    a.UpdatedAt.String(),
		},
	}, nil
}

func (s *server) GetArticles(ctx context.Context, in *emptypb.Empty) (*pb.GetArticlesResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	articles, err := NewArticlesActorId(parse.ActorId).GetArticles()
	if err != nil {
		return nil, err
	}
	var articlesInfo []*pb.ArticleInfo
	for _, a := range articles {
		articlesInfo = append(articlesInfo, &pb.ArticleInfo{
			Title:       a.Title,
			Summary:     a.Summary,
			Article:     a.Article,
			Tags:        a.Tags,
			Nsfw:        a.NSFW,
			To:          a.TO,
			Cc:          a.CC,
			Bto:         a.BTO,
			Audience:    a.Audience,
			Attachments: a.Attachments,
			Visibility:  a.Visibility,
			Id:          int64(a.ID),
			CreateAt:    a.CreatedAt.String(),
			UpdateAt:    a.UpdatedAt.String(),
		})
	}
	return &pb.GetArticlesResponse{
		Code:     "200",
		Articles: articlesInfo,
	}, nil
}

func (s *server) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	articles := new(Articles)
	switch {
	case in.Title != "":
		articles.EditTitle(in.Title)
	case in.Summary != "":
		articles.EditSummary(in.Summary)
	case in.Article != "":
		articles.EditArticle(in.Article)
	case len(in.Tags) != 0:
		articles.EditTags(in.Tags)
	case len(in.Attachments) != 0:
		articles.EditAttachments(in.Attachments)
	}
	articles.ID = uint(in.GetId())
	articles.ActorId = parse.ActorId
	if err := articles.Edit(); err != nil {
		return nil, err
	}
	return &pb.UpdateResponse{
		Code:   "200",
		Status: "ok",
	}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	id, err := strconv.Atoi(in.GetArticleId())
	if err != nil {
		return nil, err
	}
	if err := NewArticlesDelete(uint(id), parse.ActorId).Delete(); err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{
		Code:   "200",
		Status: "ok",
	}, nil
}

func (s *server) DeleteArticles(ctx context.Context, in *pb.DeleteArticlesRequest) (*pb.DeleteArticlesResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	if err := NewArticlesActorId(parse.ActorId).DeleteArticles(); err != nil {
		return nil, err
	}
	return &pb.DeleteArticlesResponse{
		Code:   "200",
		Status: "ok",
	}, nil
}
