package activity

import pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"

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

type AHandler interface {
	Follow() (*pb.ActivityResponse, error)

	Accept(data []byte) (*pb.ActivityResponse, error)
	Reject(data []byte) (*pb.ActivityResponse, error)
	Undo(data []byte) (*pb.ActivityResponse, error)

	Create(address string, in *pb.ArticleCreateActivityRequest)
	//Delete(data []byte, to []string) (*pb.ActivityResponse, error)
}
