package inbox

import (
	"encoding/json"
	"fmt"
	"github.com/hvxahv/hvx/APIs/v1alpha1/account"
	"github.com/hvxahv/hvx/activitypub"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/cockroach"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"strconv"
)

const (
	InboxTableName = "inboxes"
)

type Inboxes struct {
	gorm.Model

	ReceiverId   uint   `gorm:"primaryKey;type:bigint;receiver_id"`
	SenderAddr   string `gorm:"type:text;sender_addr"`
	ActivityId   string `gorm:"primaryKey;type:text;activity_id"`
	ActivityType string `gorm:"type:text;activity_type"`
	ActivityBody string `gorm:"type:text;activity_body"`
}

func NewInboxes(receiverId uint, senderId, activityId, activityType string, activityBody []byte) *Inboxes {
	return &Inboxes{
		ReceiverId:   receiverId,
		SenderAddr:   senderId,
		ActivityId:   activityId,
		ActivityType: activityType,
		ActivityBody: string(activityBody),
	}
}

type Ibx interface {
	Create() error
	Delete() error
	GetInbox() (*Inboxes, error)
	DeleteInbox() error
	GetInboxes() ([]*Inboxes, error)
}

func (i *Inboxes) Create() error {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Inboxes{}); err != nil {
		return err
	}

	if err := db.Debug().
		Table(InboxTableName).
		Create(i).Error; err != nil {
		return err
	}
	return nil
}

func NewInboxesActivityId(activityId string) *Inboxes {
	return &Inboxes{ActivityId: activityId}
}

func (i *Inboxes) Delete() error {
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(InboxTableName).
		Where("activity_id = ?", i.ActivityId).
		Unscoped().
		Delete(Inboxes{}).
		Error; err != nil {
		return err
	}
	return nil
}

func NewInboxesIdAndActorId(id, receiverId uint) *Inboxes {
	return &Inboxes{
		Model: gorm.Model{
			ID: id,
		},
		ReceiverId: receiverId,
	}
}

func (i *Inboxes) GetInbox() (*Inboxes, error) {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(InboxTableName).
		Where("id = ? AND receiver_id = ?", i.ID, i.ReceiverId).
		First(&i).
		Error; err != nil {
		return nil, err
	}

	return i, nil
}

func (i *Inboxes) DeleteInbox() error {
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(InboxTableName).
		Where("id = ? AND receiver_id = ?", i.ID, i.ReceiverId).
		Unscoped().
		Delete(Inboxes{}).Error; err != nil {
		return err
	}
	return nil
}

func NewInboxesReceiverId(receiverId uint) *Inboxes {
	return &Inboxes{ReceiverId: receiverId}
}

func (i *Inboxes) GetInboxes() ([]*Inboxes, error) {
	db := cockroach.GetDB()
	var inboxes []*Inboxes
	if err := db.Debug().
		Table(InboxTableName).
		Where("receiver_id = ?", i.ReceiverId).
		Find(&inboxes).Error; err != nil {
		return nil, err
	}
	return inboxes, nil
}

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

// NewActivity The received inbox data is constructed into an InboxActivity,
// Then handed over to the handler for further processing.
func NewActivity(name string, body []byte) (*InboxActivity, error) {
	fmt.Println(string(body))

	// Go to the account server and get the ActorId from the account name received by inbox.
	ctx := context.Background()
	c, err := clientv1.New(ctx,
		microsvc.NewGRPCAddress("account").Get(),
	)
	if err != nil {
		errors.Throw("failed to connect to account server during inbox processing.", err)
		return nil, err
	}
	defer c.Close()

	acct, err := account.NewAccountsClient(c.Conn).GetByUsername(ctx, &account.GetByUsernameRequest{
		Username: name,
	})
	if err != nil {
		return nil, errors.New(errors.ErrAccountGetByUsername)
	}

	receiver, err := strconv.Atoi(acct.ActorId)
	if err != nil {
		return nil, err
	}

	// TODO - Use sync.Pool to optimize performance.
	a := Activity{}
	if err := json.Unmarshal(body, &a); err != nil {
		return nil, errors.New("UNMARSHAL_ACTIVITY")
	}
	return &InboxActivity{
		ReceiverId:   uint(receiver),
		SenderId:     a.Actor,
		ActivityType: a.Type,
		ActivityId:   a.ID,
		ActivityData: body,
	}, nil
}

func (ibx *InboxActivity) Handler() error {
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

// func (a *activity) GetInboxByActivityID(ctx context.Context, in *pb.GetInboxByActivityIDRequest) (*pb.GetInboxByActivityIDResponse, error) {
// 	aid, err := strconv.Atoi(in.GetActivityId())
// 	if err != nil {
// 		return nil, err
// 	}
// 	db := cockroach.GetDB()
// 	if err := db.Debug().
// 		Table(InboxTableName).
// 		Where("id = ?", uint(aid)).
// 		First(&a.Inboxes).
// 		Error; err != nil {
// 		return nil, err
// 	}
// 	ibx := pb.Inboxes{
// 		Id:           strconv.Itoa(int(a.Inboxes.ID)),
// 		ActorId:      strconv.Itoa(int(a.Inboxes.ActorID)),
// 		ActivityId:   a.Inboxes.ActivityID,
// 		ActivityType: a.Inboxes.ActivityType,
// 		ActivityBody: []byte(a.Inboxes.ActivityBody),
// 	}
// 	return &pb.GetInboxByActivityIDResponse{Code: "200", Inbox: &ibx}, nil
// }

// func (a *activity) GetInboxesByAccountID(ctx context.Context, in *pb.GetInboxesByAccountIDRequest) (*pb.GetInboxesByAccountIDResponse, error) {
// 	aid, err := strconv.Atoi(in.GetAccountId())
// 	if err != nil {
// 		return nil, err
// 	}
// 	var inboxes []Inboxes
// 	db := cockroach.GetDB()
// 	if err := db.Debug().
// 		Table(InboxTableName).
// 		Where("account_id = ?", uint(aid)).
// 		Find(&inboxes).
// 		Error; err != nil {
// 		return nil, err
// 	}
// 	var ibx []*pb.Inboxes
// 	for _, v := range inboxes {
// 		ibx = append(ibx, &pb.Inboxes{
// 			Id:           strconv.Itoa(int(v.ID)),
// 			ActorId:      strconv.Itoa(int(v.ActorID)),
// 			ActivityId:   v.ActivityID,
// 			ActivityType: v.ActivityType,
// 			ActivityBody: []byte(v.ActivityBody),
// 		})
// 	}

// 	return &pb.GetInboxesByAccountIDResponse{Code: "200", Inboxes: ibx}, nil
// }

// func (a *activity) DeleteInboxByInboxesID(ctx context.Context, in *pb.DeleteInboxByInboxesIDRequest) (*pb.DeleteInboxByInboxesIDResponse, error) {
// 	id, err := strconv.Atoi(in.GetActivityId())
// 	if err != nil {
// 		return nil, err
// 	}
// 	aid, err := strconv.Atoi(in.GetAccountId())
// 	if err != nil {
// 		return nil, err
// 	}
// 	db := cockroach.GetDB()
// 	if err := db.Debug().
// 		Table(InboxTableName).
// 		Where("id = ? AND account_id = ?", uint(id), uint(aid)).
// 		Unscoped().
// 		Delete(&Inboxes{}).
// 		Error; err != nil {
// 		return nil, err
// 	}
// 	return &pb.DeleteInboxByInboxesIDResponse{Code: "200", Reply: "ok"}, nil
// }
