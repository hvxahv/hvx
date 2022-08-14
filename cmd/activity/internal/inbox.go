package internal

import (
	"github.com/hvxahv/hvx/cockroach"
	"gorm.io/gorm"
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

func NewInboxesGetInbox(receiverId uint, activityId string) *Inboxes {
	return &Inboxes{ReceiverId: receiverId}
}
func (i *Inboxes) GetInbox() (*Inboxes, error) {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(InboxTableName).
		Where("receiver_id = ? AND activity_id = ?", i.ReceiverId, i.ActivityId).
		First(&i).
		Error; err != nil {
		return nil, err
	}

	return i, nil
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
