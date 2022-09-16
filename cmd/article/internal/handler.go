package internal

import (
	"context"
	"github.com/hvxahv/hvx/APIs/v1alpha1/actor"
	"strconv"

	"github.com/hvxahv/hvx/clientv1"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/hvxahv/hvx/APIs/v1alpha1/article"
	"github.com/hvxahv/hvx/microsvc"
)

func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	visibility, err := strconv.Atoi(in.Visibility)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	var (
		To = StringArrayToInt64Array(in.To)
		Cc = StringArrayToInt64Array(in.Cc)
	)
	if in.State {
		if err := NewStatus(parse.ActorId,
			in.Article,
			in.Tags,
			in.AttachmentType,
			in.Attachments,
			To,
			Cc,
			in.Nsfw,
			uint(visibility),
		).Create(); err != nil {
			return nil, err
		}
	} else {
		if err := NewArticles(parse.ActorId,
			in.Title,
			in.Summary,
			in.Article,
			in.Tags,
			in.AttachmentType,
			in.Attachments,
			To,
			Cc,
			in.Nsfw,
			uint(visibility),
		).Create(); err != nil {
			return nil, err
		}
	}

	// TODO - ACTIVITYPUB FOLLOWER PUSH.

	return &pb.CreateResponse{Code: "200", Status: "ok"}, nil
}

func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	id, err := strconv.Atoi(in.GetArticleId())
	if err != nil {
		return nil, err
	}
	a, err := NewArticlesId(uint(id)).Get(parse.ActorId)
	if err != nil {
		return nil, err
	}

	// GET ACTOR DATA
	data, err := clientv1.New(ctx, microsvc.ActorServiceName).GetActor(int64(a.ActorId))
	if err != nil {
		return nil, err
	}
	return &pb.GetResponse{
		Code: "200",
		Actor: &actor.ActorData{
			Id:                data.Actor.Id,
			PreferredUsername: data.Actor.PreferredUsername,
			Domain:            data.Actor.Domain,
			Avatar:            data.Actor.Avatar,
			Name:              data.Actor.Name,
			Summary:           data.Actor.Summary,
			Inbox:             data.Actor.Inbox,
			Address:           data.Actor.Address,
			PublicKey:         data.Actor.PublicKey,
			ActorType:         data.Actor.ActorType,
			IsRemote:          data.Actor.IsRemote,
		},
		Article: &pb.ArticleInfo{
			Id:             int64(a.ID),
			Title:          a.Title,
			Summary:        a.Summary,
			Article:        a.Article,
			Tags:           a.Tags,
			AttachmentType: a.AttachmentType,
			Attachments:    a.Attachments,
			To:             Int64ArrayToStringArray(a.TO),
			Cc:             Int64ArrayToStringArray(a.CC),
			State:          a.Statuses,
			Nsfw:           a.NSFW,
			Visibility:     strconv.Itoa(int(a.Visibility)),
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
			Id:             int64(a.ID),
			Title:          a.Title,
			Summary:        a.Summary,
			Article:        a.Article,
			Tags:           a.Tags,
			AttachmentType: a.AttachmentType,
			Attachments:    a.Attachments,
			To:             Int64ArrayToStringArray(a.TO),
			Cc:             Int64ArrayToStringArray(a.CC),
			State:          a.Statuses,
			Nsfw:           a.NSFW,
			Visibility:     strconv.Itoa(int(a.Visibility)),
		})
	}
	return &pb.GetArticlesResponse{
		Code: "200",
		Data: articlesInfo,
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
	case in.AttachmentType != "":
		articles.EditAttachmentType(in.AttachmentType)
	case len(in.Attachments) != 0:
		articles.EditAttachments(in.Attachments)
	case in.Nsfw == "":
		nsfw, err := strconv.ParseBool(in.Nsfw)
		if err != nil {
			return nil, err
		}
		articles.EditNSFW(nsfw)
	case in.Visibility != "":
		v, err := strconv.Atoi(in.Visibility)
		if err != nil {
			return nil, err
		}
		articles.EditVisibility(uint(v))
	}

	id, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}
	if err := articles.Update(uint(id), parse.ActorId); err != nil {
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
	if err := NewArticlesDelete(uint(in.GetArticleId()), parse.ActorId).Delete(); err != nil {
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
