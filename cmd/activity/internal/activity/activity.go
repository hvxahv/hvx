package activity

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"
)

// Activity Types inherit the properties of the base Activity type.
// Some specific Activity Types are subtypes or specializations of more generalized Activity Types
// (for instance, the Invite Activity Type is a more specific form of the Offer Activity Type).
// The Activity Types include:
// https://www.w3.org/TR/activitystreams-vocabulary/#activity-types
type Activity struct {
	Context interface{} `json:"@context"`
	ID      string      `json:"id"`
	Type    string      `json:"type"`
	Actor   string      `json:"actor"`
	Object  interface{} `json:"object"`
}

type Object struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Actor  string `json:"actor"`
	Object string `json:"object"`
}

type Handler struct {
	// inbox is address of the Actor object.
	// We need to know who is sending us the activity.
	inbox string

	// aAddr actor address.
	aAddr string

	// privateKey used to encrypt activitypub delivery requests.
	// It is stored in the account and the corresponding public key is stored in the actor.
	privateKey string

	// actorId of the activity recipient actor,
	// since the id is to be used as the primary key of the inboxes table.
	actorId uint
}

func NewHandler(inbox string, aAddr string, privateKey string, actorId uint) *Handler {
	return &Handler{inbox: inbox, aAddr: aAddr, privateKey: privateKey, actorId: actorId}
}

type AHandler interface {
	Follow() (*pb.ActivityResponse, error)
	Accept(data []byte) (*pb.ActivityResponse, error)
	Reject(data []byte) (*pb.ActivityResponse, error)
	Undo(data []byte) (*pb.ActivityResponse, error)
	Create(data []byte, to, cc []string) (*pb.ActivityResponse, error)
	Delete(data []byte, to []string) (*pb.ActivityResponse, error)
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
