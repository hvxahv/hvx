package internal

import (
	"encoding/json"
	"fmt"
)

// All Activity Types inherit the properties of the base Activity type.
// Some specific Activity Types are subtypes or specializations of more generalized Activity Types
// (for instance, the Invite Activity Type is a more specific form of the Offer Activity Type).
// The Activity Types include:
// https://www.w3.org/TR/activitystreams-vocabulary/#activity-types

type Activity struct {
	Context string      `json:"@context"`
	ID      string      `json:"id"`
	Type    string      `json:"type"`
	Actor   string      `json:"actor"`
	Object  interface{} `json:"object"`
}

func NewActivityInbox(name string, body []byte) *Inbox {
	a := Activity{}
	if err := json.Unmarshal(body, &a); err != nil {
		fmt.Printf("Unmarshal acticity type:%v", err)
	}
	return &Inbox{
		CurrentUsername: name,
		ActivityActor:   a.Actor,
		ActivityType:    a.Type,
		ActivityId:      a.ID,
		ActivityData:    body,
	}
}
