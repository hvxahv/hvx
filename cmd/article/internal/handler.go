package internal

import (
	"context"
	"strconv"

	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/article"
	"github.com/hvxahv/hvx/microsvc"
)

func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{}, nil
}

func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{}, nil
}

func (s *server) GetArticles(ctx context.Context, in *pb.GetArticlesRequest) (*pb.GetArticlesResponse, error) {
	return &pb.GetArticlesResponse{}, nil
}

func (s *server) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	// Get the updated content.
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
	accountId, err := microsvc.GetAccountIDWithContext(ctx)
	if err != nil {
		return nil, err
	}
	id, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}
	if err := articles.Update(uint(id), accountId); err != nil {
		return nil, err
	}
	return &pb.UpdateResponse{
		Code:   "200",
		Status: "ok",
	}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return &pb.DeleteResponse{}, nil
}

func (s *server) DeleteArticles(ctx context.Context, in *pb.DeleteArticlesRequest) (*pb.DeleteArticlesResponse, error) {
	return &pb.DeleteArticlesResponse{}, nil
}
