package activity

import (
	"context"
	"encoding/json"
	"fmt"
	actor "github.com/hvxahv/hvxahv/api/account/v1alpha1"
	pb "github.com/hvxahv/hvxahv/api/activity/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/account"
	"github.com/hvxahv/hvxahv/pkg/activitypub"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"gorm.io/gorm"
	"strconv"
)

const (
	InboxTableName = "inboxes"
)

type Inbox struct {
	CurrentUsername string
	ActivityActor   string
	ActivityType    string
	ActivityId      string
	ActivityData    []byte
}

type Inboxes struct {
	gorm.Model

	// AccountID is current account id.
	AccountID  uint   `gorm:"primaryKey;type:bigint;account_id"`
	ActivityID string `gorm:"primaryKey;type:text;activity_id"`

	// ActorID is activity actor id.
	ActorID      uint   `gorm:"type:bigint;actor_id"`
	ActivityType string `gorm:"type:text;activity_type"`
	ActivityBody string `gorm:"type:text;activity_body"`
}

func NewInboxes(accountID uint, activityID string, actorID uint, activityType string, activityBody []byte) *Inboxes {
	return &Inboxes{AccountID: accountID, ActivityID: activityID, ActorID: actorID, ActivityType: activityType, ActivityBody: string(activityBody)}
}

func (i *Inboxes) Create() error {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Inboxes{}); err != nil {
		return err
	}
	if err := db.Debug().
		Table(InboxTableName).
		Create(NewInboxes(i.AccountID, i.ActivityID, i.ActorID, i.ActivityType, []byte(i.ActivityBody))).
		Error; err != nil {
		return err
	}
	return nil
}

func NewActivityID(activityID string) *Inboxes {
	return &Inboxes{ActivityID: activityID}
}

func (i *Inboxes) DeleteByActivityID() error {
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(InboxTableName).
		Where("activity_id = ?", i.ActivityID).
		Unscoped().
		Delete(Inboxes{}).
		Error; err != nil {
		return err
	}
	return nil
}

func (a *activity) Inbox(ctx context.Context, in *pb.InboxRequest) (*pb.InboxResponse, error) {
	ibx := NewActivityInbox(in.GetName(), in.GetData())
	client, err := account.GetActorClient()
	if err != nil {
		return nil, err
	}
	act, err := client.GetActorByAccountUsername(ctx, &actor.GetActorByAccountUsernameRequest{
		Username: ibx.CurrentUsername,
	})
	if err != nil {
		return nil, err
	}
	address, err := client.GetActorByAddress(ctx, &actor.GetActorByAddressRequest{
		Address: ibx.ActivityActor,
	})
	if err != nil {
		return nil, err
	}

	aid, err := strconv.Atoi(act.Id)
	if err != nil {
		return nil, err
	}
	addressID, err := strconv.Atoi(address.Id)
	if err != nil {
		return nil, err
	}

	switch ibx.ActivityType {
	case "Follow":
		fmt.Println("Follow")
		if err := NewInboxes(uint(aid), ibx.ActivityId, uint(addressID), ibx.ActivityType, ibx.ActivityData).Create(); err != nil {
			return nil, err
		}
	case "Undo":
		undo := activitypub.Undo{}
		if err := json.Unmarshal(in.Data, &undo); err != nil {
			fmt.Println(err)
		}
		fmt.Println("Undo")
		if err := NewActivityID(undo.Object.Object).DeleteByActivityID(); err != nil {
			return nil, err
		}
	case "Accept":
		accept := activitypub.Accept{}
		if err := json.Unmarshal(in.Data, &accept); err != nil {
			return nil, err
		}
		fmt.Println(accept)
		fmt.Println("Accept")
	case "Reject":
		reject := activitypub.Reject{}
		if err := json.Unmarshal(in.Data, &reject); err != nil {
			fmt.Println(err)
		}
		fmt.Println("Reject")
	case "Create":
		fmt.Println("Create")
	case "Announce":
		fmt.Println("Announce")
	case "Like":
		fmt.Println("Like")
	case "Dislike":
		fmt.Println("Dislike")
	case "Delete":
		fmt.Println("Delete")
	case "Update":
		fmt.Println("Update")
	case "Add":
		fmt.Println("Add")
	case "Remove":
		fmt.Println("Remove")
	case "Move":
		fmt.Println("Move")
	case "Block":
		fmt.Println("Block")
	case "Unblock":
		fmt.Println("Unblock")
	case "Flag":
		fmt.Println("Flag")
	case "Unflag":
		fmt.Println("Unflag")
	default:
		fmt.Println("default")
	}

	return &pb.InboxResponse{
		Code:     "200",
		Response: "ok",
	}, nil
}

func (a *activity) GetInboxByActivityID(ctx context.Context, in *pb.GetInboxByActivityIDRequest) (*pb.GetInboxByActivityIDResponse, error) {
	aid, err := strconv.Atoi(in.GetActivityId())
	if err != nil {
		return nil, err
	}
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(InboxTableName).
		Where("id = ?", uint(aid)).
		First(&a.Inboxes).
		Error; err != nil {
		return nil, err
	}
	ibx := pb.Inboxes{
		Id:           strconv.Itoa(int(a.Inboxes.ID)),
		ActorId:      strconv.Itoa(int(a.Inboxes.ActorID)),
		ActivityId:   a.Inboxes.ActivityID,
		ActivityType: a.Inboxes.ActivityType,
		ActivityBody: []byte(a.Inboxes.ActivityBody),
	}
	return &pb.GetInboxByActivityIDResponse{Code: "200", Inbox: &ibx}, nil
}

func (a *activity) GetInboxesByAccountID(ctx context.Context, in *pb.GetInboxesByAccountIDRequest) (*pb.GetInboxesByAccountIDResponse, error) {
	aid, err := strconv.Atoi(in.GetAccountId())
	if err != nil {
		return nil, err
	}
	var inboxes []Inboxes
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(InboxTableName).
		Where("account_id = ?", uint(aid)).
		Find(&inboxes).
		Error; err != nil {
		return nil, err
	}
	var ibx []*pb.Inboxes
	for _, v := range inboxes {
		ibx = append(ibx, &pb.Inboxes{
			Id:           strconv.Itoa(int(v.ID)),
			ActorId:      strconv.Itoa(int(v.ActorID)),
			ActivityId:   v.ActivityID,
			ActivityType: v.ActivityType,
			ActivityBody: []byte(v.ActivityBody),
		})
	}

	return &pb.GetInboxesByAccountIDResponse{Code: "200", Inboxes: ibx}, nil
}

func (a *activity) DeleteInboxByInboxesID(ctx context.Context, in *pb.DeleteInboxByInboxesIDRequest) (*pb.DeleteInboxByInboxesIDResponse, error) {
	id, err := strconv.Atoi(in.GetActivityId())
	if err != nil {
		return nil, err
	}
	aid, err := strconv.Atoi(in.GetAccountId())
	if err != nil {
		return nil, err
	}
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(InboxTableName).
		Where("id = ? AND account_id = ?", uint(id), uint(aid)).
		Unscoped().
		Delete(&Inboxes{}).
		Error; err != nil {
		return nil, err
	}
	return &pb.DeleteInboxByInboxesIDResponse{Code: "200", Reply: "ok"}, nil
}
