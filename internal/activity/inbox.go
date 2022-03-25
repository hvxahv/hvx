package activity

import (
	"context"
	actor "github.com/hvxahv/hvxahv/api/account/v1alpha1"
	pb "github.com/hvxahv/hvxahv/api/activity/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/account"
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

	CurrentAccountID uint   `gorm:"primaryKey;type:bigint;account_id"`
	ActivityID       string `gorm:"primaryKey;type:text;activity_id"`
	ActivityActorID  uint   `gorm:"type:bigint;activity_actor_id"`
	ActivityType     string `gorm:"type:text;activity_type"`
	ActivityBody     string `gorm:"type:text;activity_body"`
}

func NewInboxes(currentAccountID uint, activityID string, activityActorID uint, activityType string, activityBody []byte) *Inboxes {
	return &Inboxes{CurrentAccountID: currentAccountID, ActivityID: activityID, ActivityActorID: activityActorID, ActivityType: activityType, ActivityBody: string(activityBody)}
}

func (a *activity) CreateInbox(ctx context.Context, in *pb.CreateInboxRequest) (*pb.CreateInboxResponse, error) {
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

	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Inboxes{}); err != nil {
		return nil, err
	}
	if err := db.Debug().
		Table(InboxTableName).
		Create(NewInboxes(uint(aid), ibx.ActivityId, uint(addressID), ibx.ActivityType, ibx.ActivityData)).
		Error; err != nil {
		return nil, err
	}

	return &pb.CreateInboxResponse{Code: "200", Response: "success."}, nil
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
		ActorId:      strconv.Itoa(int(a.Inboxes.ActivityActorID)),
		ActivityId:   a.Inboxes.ActivityID,
		ActivityType: a.Inboxes.ActivityType,
		ActivityBody: []byte(a.Inboxes.ActivityBody),
	}
	return &pb.GetInboxByActivityIDResponse{Code: "200", Inbox: &ibx}, nil
}
