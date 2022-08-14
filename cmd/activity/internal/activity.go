package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hvxahv/hvx/APIs/v1alpha1/account"
	"github.com/hvxahv/hvx/activitypub"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/microsvc"
	"strconv"
)

// All Activity Types inherit the properties of the base Activity type.
// Some specific Activity Types are subtypes or specializations of more generalized Activity Types
// (for instance, the Invite Activity Type is a more specific form of the Offer Activity Type).
// The Activity Types include:
// https://www.w3.org/TR/activitystreams-vocabulary/#activity-types
//

type Activity struct {
	Context string      `json:"@context"`
	ID      string      `json:"id"`
	Type    string      `json:"type"`
	Actor   string      `json:"actor"`
	Object  interface{} `json:"object"`
}

type InboxActivity struct {
	ReceiverId   uint
	SenderId     string
	ActivityType string
	ActivityId   string
	ActivityData []byte
}

func NewActivity(name string, body []byte) (*InboxActivity, error) {
	a := Activity{}
	if err := json.Unmarshal(body, &a); err != nil {
		return nil, errors.New("UNMARSHAL_ACTIVITY")
	}

	ctx := context.Background()
	client, err := clientv1.New(ctx,
		microsvc.NewGRPCAddress("account").Get(),
	)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	acct, err := account.NewAccountsClient(client.Conn).GetByUsername(ctx, &account.GetByUsernameRequest{
		Username: name,
	})
	rec, err := strconv.Atoi(acct.ActorId)
	if err != nil {
		return nil, err
	}

	return &InboxActivity{
		ReceiverId:   uint(rec),
		SenderId:     a.Actor,
		ActivityType: a.Type,
		ActivityId:   a.ID,
		ActivityData: body,
	}, nil
}

func (ibx *InboxActivity) Activity() error {
	switch ibx.ActivityType {
	case "Follow":
		fmt.Println("Follow")
		if err := NewInboxes(ibx.ReceiverId, ibx.SenderId, ibx.ActivityId, ibx.ActivityType, ibx.ActivityData).Create(); err != nil {
			return err
		}
	case "Undo":
		fmt.Println("Undo")
		undo := activitypub.Undo{}
		if err := json.Unmarshal(ibx.ActivityData, &undo); err != nil {
			return errors.New("UNMARSHAL_ACTIVITY_UNDO")
		}
		if err := NewInboxesActivityId(undo.Object.Id).Delete(); err != nil {
			return err
		}
	default:
		fmt.Println("Unknown")
	}

	return nil
}
