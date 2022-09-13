package inbox

import (
	"encoding/json"
	"fmt"
	"github.com/hvxahv/hvx/activitypub"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/cmd/activity/internal/activity"
	"github.com/hvxahv/hvx/cmd/activity/internal/friendship"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
	"strconv"
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

	// Go to the account server and get the ActorId from the account name received by inbox.
	ctx := context.Background()
	account, err := clientv1.New(ctx, microsvc.AccountServiceName).GetAccountByUsername(name)
	if err != nil {
		errors.Throw("failed to connect to account server during inbox processing.", err)
		return nil, errors.New(errors.ErrAccountGetByUsername)
	}

	actorId, err := strconv.Atoi(account.ActorId)
	if err != nil {
		return nil, err
	}
	return NewHandler(uint(actorId), body), nil
}

func (h *Handler) Handler() error {
	a := activity.Activity{}
	if err := json.Unmarshal(h.body, &a); err != nil {
		return errors.New("UNMARSHAL_ACTIVITY")
	}
	switch a.Type {
	case "Follow":
		fmt.Println("Follow")
		var f activitypub.Follow
		if err := json.Unmarshal(h.body, &f); err != nil {
			return err
		}
		if err := NewInboxes(h.actorId, f.Id, f.Actor, f.Type, string(h.body)).Create(); err != nil {
			return err
		}

	case "Undo":
		fmt.Println("Undo")
		undo := activitypub.Undo{}
		if err := json.Unmarshal(h.body, &undo); err != nil {
			return errors.New("UNMARSHAL_ACTIVITY_UNDO")
		}
		if err := NewInboxesActivityId(undo.Object.Id).Delete(); err != nil {
			return err
		}

		switch undo.Object.Type {
		case "Follow":
			if err := NewInboxes(h.actorId, undo.Id, undo.Actor, undo.Type, string(h.body)).Create(); err != nil {
				return err
			}
			objectId, err := GetObjectActorId(undo.Actor)
			if err != nil {
				return err
			}
			// UNFOLLOWING ...
			if err := friendship.NewFollower(h.actorId, objectId).UNFollow(); err != nil {
				return err
			}

		default:
			fmt.Println("Unknown Undo")
		}

	case "Accept":
		accept := activitypub.Accept{}
		if err := json.Unmarshal(h.body, &accept); err != nil {
			fmt.Println("ERR", err)
			return err
		}
		switch accept.Object.Type {
		case "Follow":
			if err := NewInboxes(h.actorId, accept.Id, accept.Actor, accept.Type, string(h.body)).Create(); err != nil {
				return err
			}

			objectId, err := GetObjectActorId(accept.Actor)
			if err != nil {
				return err
			}

			// The following should be added because the person accepted your follow request.
			if err := friendship.NewFollowing(h.actorId, objectId).Follow(); err != nil {
				return err
			}

		default:
			fmt.Println("Unknown Accept")
		}

	case "Delete":
		fmt.Println("DELETE REQUEST...")

	case "Reject":
		fmt.Println("REJECT REQUEST...")

		reject := activitypub.Reject{}
		if err := json.Unmarshal(h.body, &reject); err != nil {
			return err
		}

		switch reject.Object.Type {
		case "Follow":
			// TODO - CREATE INBOX
			if err := NewInboxes(h.actorId, reject.Id, reject.Actor, reject.Type, string(h.body)).Create(); err != nil {
				return err
			}
			objectId, err := GetObjectActorId(reject.Actor)
			if err != nil {
				return err
			}

			// UN FOLLOWER...
			if err := friendship.NewFollowing(h.actorId, objectId).UNFollow(); err != nil {
				return err
			}

		default:
			fmt.Println("Unknown Accept")
		}
	default:
		fmt.Println("Unknown")
	}

	return nil
}

func GetObjectActorId(actor string) (uint, error) {
	object, err := clientv1.New(context.Background(), microsvc.ActorServiceName).GetActorByAddress(actor)
	if err != nil {
		return 0, err
	}
	objectId, err := strconv.Atoi(object.Id)
	if err != nil {
		return 0, err
	}

	return uint(objectId), nil

}

//
//	switch ibx.ActivityType {
//	case "Follow":
//		fmt.Println("Follow")
//		if err := NewInboxes(uint(aid), ibx.ActivityId, uint(addressID), ibx.ActivityType, ibx.ActivityData).Create(); err != nil {
//			return err
//		}
//	case "Undo":
//		undo := activitypub.Undo{}
//		if err := json.Unmarshal(in.Data, &undo); err != nil {
//			fmt.Println(err)
//		}
//		fmt.Println("Undo")
//		if err := NewActivityID(undo.Object.Object).DeleteByActivityID(); err != nil {
//			return err
//		}
//	case "Accept":
//		accept := activitypub.Accept{}
//		if err := json.Unmarshal(in.Data, &accept); err != nil {
//			return err
//		}
//		fmt.Println(accept)
//		fmt.Println("Accept")
//	case "Reject":
//		reject := activitypub.Reject{}
//		if err := json.Unmarshal(in.Data, &reject); err != nil {
//			fmt.Println(err)
//		}
//		fmt.Println("Reject")
//	case "Create":
//		fmt.Println("Create")
//	case "Announce":
//		fmt.Println("Announce")
//	case "Like":
//		fmt.Println("Like")
//	case "Dislike":
//		fmt.Println("Dislike")
//	case "Delete":
//		fmt.Println("Delete")
//	case "Update":
//		fmt.Println("Update")
//	case "Add":
//		fmt.Println("Add")
//	case "Remove":
//		fmt.Println("Remove")
//	case "Move":
//		fmt.Println("Move")
//	case "Block":
//		fmt.Println("Block")
//	case "Unblock":
//		fmt.Println("Unblock")
//	case "Flag":
//		fmt.Println("Flag")
//	case "Unflag":
//		fmt.Println("Unflag")
//	default:
//		fmt.Println("default")
//	}
//
//	return nil
//}
