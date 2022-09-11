package activity

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"
)

type Object struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Actor  string `json:"actor"`
	Object string `json:"object"`
}

type Handler struct {
	inbox      string
	aAddr      string
	privateKey string
	actorId    uint
}

func NewHandler(inbox string, aAddr string, privateKey string, actorId uint) *Handler {
	return &Handler{inbox: inbox, aAddr: aAddr, privateKey: privateKey, actorId: actorId}
}

type AHandler interface {
	Follow() (*pb.ActivityResponse, error)
	Accept(data []byte) (*pb.ActivityResponse, error)
	Reject(data []byte) (*pb.ActivityResponse, error)
	Undo(data []byte) (*pb.ActivityResponse, error)
}

func response(notok, ok []string) (*pb.ActivityResponse, error) {
	if len(notok) != 0 {
		return &pb.ActivityResponse{
			Code:   "200",
			Status: FailedToDelivery,
			Inbox:  notok,
		}, nil
	}
	return &pb.ActivityResponse{
		Code:   "200",
		Status: DeliverySuccessfully,
		Inbox:  ok,
	}, nil
}
