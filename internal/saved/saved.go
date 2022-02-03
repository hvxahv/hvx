package saved

import (
	pb "github.com/hvxahv/hvxahv/api/saved/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"strconv"
)

// Change the upload logic.
// The upload should be encrypted in the client and uploaded to the IPFS API,
// and then the returned hash should be submitted to the hvxahv server for storage.
// hvxahv should not store the user's file key,
// only the saved hash value is used to format the storage list.

type Saves struct {
	gorm.Model

	AccountID   uint   `gorm:"primaryKey;type:bigint;accounts_id"`
	Name        string `gorm:"type:text;name"`
	Description string `gorm:"type:text;description"`
	Hash        string `gorm:"type:text;hash"`
	Types       string `gorm:"type:text;types"`
}

func (s *saved) Create(ctx context.Context, in *pb.NewSavedCreate) (*pb.SavedReply, error) {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Saves{}); err != nil {
		return nil, err
	}
	id, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}
	c := NewSaves(uint(id), in.Name, in.Description, in.Hash, in.Types)
	if err := db.Debug().Table("saves").Create(&c).Error; err != nil {
		return nil, err
	}
	return &pb.SavedReply{Code: "200", Reply: "ok"}, nil
}

func (s *saved) GetSaves(ctx context.Context, in *pb.NewSavedAccountID) (*pb.GetSavesReply, error) {
	db := cockroach.GetDB()
	id, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}
	var saves []*pb.SavedData
	if err := db.Debug().Table("saves").Where("account_id = ?", uint(id)).Find(&saves).Error; err != nil {
		return nil, err
	}
	return &pb.GetSavesReply{Code: "200", Saves: saves}, nil
}

func (s *saved) GetSaved(ctx context.Context, in *pb.NewSavedID) (*pb.SavedData, error) {
	db := cockroach.GetDB()
	id, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}
	var saves Saves
	if err := db.Debug().Table("saves").Where("id = ?", uint(id)).First(&saves).Error; err != nil {
		return nil, err
	}
	return &pb.SavedData{
		Id:          strconv.Itoa(int(saves.ID)),
		Name:        saves.Name,
		Description: saves.Description,
		Hash:        saves.Hash,
		Types:       saves.Types,
	}, nil
}

func NewSaves(accountID uint, name string, description string, hash string, types string) *Saves {
	return &Saves{AccountID: accountID, Name: name, Description: description, Hash: hash, Types: types}
}
