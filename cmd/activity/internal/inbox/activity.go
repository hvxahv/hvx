package inbox

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hvxahv/hvx/activitypub"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/cmd/activity/internal/activity"
	"github.com/hvxahv/hvx/cmd/activity/internal/friendship"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
)

type Handler struct {
	// actorId of the activity recipient actor,
	// since the id is to be used as the primary key of the inboxes table.
	actorId uint

	body []byte
}

func NewHandler(actorId uint, body []byte) *Handler {
	return &Handler{actorId: actorId, body: body}
}

// NewActivity The received inbox data is constructed into an InboxActivity,
// Then handed over to the handler for further processing.
func NewActivity(name string, body []byte) (*Handler, error) {
	fmt.Println(string(body))

	actor, err := clientv1.New(context.Background(), microsvc.ActorServiceName).GetActorByUsername(name)
	if err != nil {
		return nil, err
	}
	return NewHandler(uint(actor.Id), body), nil
}

func (h *Handler) Handler() error {
	a := activity.Activity{}
	if err := json.Unmarshal(h.body, &a); err != nil {
		errors.Throw("UNMARSHAL_ACTIVITY_INBOX", err)
		return errors.New("UNMARSHAL_ACTIVITY")
	}

	switch a.Type {
	case activitypub.FollowType:
		var f activitypub.Follow
		if err := json.Unmarshal(h.body, &f); err != nil {
			return err
		}
		// Requested object of attention
		actor, err := clientv1.New(context.Background(), microsvc.ActorServiceName).GetActorByAddress(f.Object)
		if err != nil {
			return err
		}

		// Request Followers.
		object, err := clientv1.New(context.Background(), microsvc.ActorServiceName).GetActorByAddress(f.Actor)
		if err != nil {
			return err
		}

		switch actor.ActorType {
		case activitypub.PersonType:
			if err := NewInboxes(h.actorId, f.Id, f.Actor, f.Type, string(h.body)).Create(); err != nil {
				return err
			}
		//	If the actor requesting the follow is a service, the default response is to agree to the follow.
		case activitypub.ChannelType:
			// SEND ACCEPT AND ADD CHANNEL SUB.

			marshal, err := json.Marshal(&activitypub.Accept{
				Context: "https://www.w3.org/ns/activitystreams",
				Id:      fmt.Sprintf("%s/%s", actor.Address, uuid.NewString()),
				Type:    activitypub.AcceptType,
				Actor:   actor.Address,
				Object: struct {
					Id     string `json:"id"`
					Type   string `json:"type"`
					Actor  string `json:"actor"`
					Object string `json:"object"`
				}{
					Id:     f.Id,
					Type:   f.Type,
					Actor:  f.Actor,
					Object: f.Object,
				},
			})
			if err != nil {
				fmt.Println(err)
			}
			c, err := clientv1.New(context.Background(), microsvc.ChannelServiceName).GetPrivateKeyByActorId(actor.Id)
			if err != nil {
				return err
			}
			// DELIVERY ...
			do, err := activity.NewDelivery(marshal, actor.Address, c.PrivateKey).Do(object.Inbox)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(do)
		}

	case activitypub.UndoType:
		undo := activitypub.Undo{}
		if err := json.Unmarshal(h.body, &undo); err != nil {
			return errors.New("UNMARSHAL_ACTIVITY_UNDO")
		}
		objectId, err := clientv1.New(context.Background(), microsvc.ActorServiceName).GetActorByAddress(undo.Actor)
		if err != nil {
			return err
		}
		if err := NewInboxesDeleteByActivityId(uint(objectId.GetId()), undo.Object.Id).DeleteByActivityId(); err != nil {
			return err
		}

		switch undo.Object.Type {
		case activitypub.FollowType:
			//if err := NewInboxes(h.actorId, undo.Id, undo.Actor, undo.Type, string(h.body)).Create(); err != nil {
			//	return err
			//}
			if err := friendship.NewFollower(h.actorId, uint(objectId.GetId())).UNFollow(); err != nil {
				return err
			}

		default:
			fmt.Println("Unknown Undo")
		}

	case activitypub.AcceptType:
		accept := activitypub.Accept{}
		if err := json.Unmarshal(h.body, &accept); err != nil {
			fmt.Println("ERR", err)
			return err
		}

		switch accept.Object.Type {
		case activitypub.FollowType:
			if err := NewInboxes(h.actorId, accept.Id, accept.Actor, accept.Type, string(h.body)).Create(); err != nil {
				return err
			}

			objectId, err := clientv1.New(context.Background(), microsvc.ActorServiceName).GetActorByAddress(accept.Actor)
			if err != nil {
				return err
			}
			// The following should be added because the person accepted your follow request.
			if err := friendship.NewFollowing(h.actorId, uint(objectId.GetId())).Follow(); err != nil {
				return err
			}

		default:
			fmt.Println("Unknown Accept")
		}

	case activitypub.RejectType:
		reject := activitypub.Reject{}
		if err := json.Unmarshal(h.body, &reject); err != nil {
			return err
		}

		switch reject.Object.Type {
		case activitypub.FollowType:
			// TODO - CREATE INBOX
			if err := NewInboxes(h.actorId, reject.Id, reject.Actor, reject.Type, string(h.body)).Create(); err != nil {
				return err
			}
			objectId, err := clientv1.New(context.Background(), microsvc.ActorServiceName).GetActorByAddress(reject.Actor)
			if err != nil {
				return err
			}

			// UN FOLLOWER...
			if err := friendship.NewFollowing(h.actorId, uint(objectId.GetId())).UNFollow(); err != nil {
				return err
			}

		default:
			fmt.Println("Unknown Accept")
		}

	case activitypub.DeleteType:
		d := activitypub.Delete{}
		if err := json.Unmarshal(h.body, &d); err != nil {
			return err
		}

		if err := NewInboxesDeleteByActivityId(h.actorId, d.Object.Id).DeleteByActivityId(); err != nil {
			return err
		}

	case activitypub.CreateType:
		create := activitypub.Create{}
		if err := json.Unmarshal(h.body, &create); err != nil {
			errors.Throw("UNMARSHAL_ACTIVITY_CREATE_INBOX", err)
			return err
		}

		if err := NewInboxes(h.actorId, create.Object.Id, create.Actor, create.Type, string(h.body)).Create(); err != nil {
			return err
		}

	default:
		fmt.Println("Unknown")
	}

	return nil
}
