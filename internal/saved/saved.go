package saved

import (
	pb "github.com/hvxahv/hvxahv/api/saved/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"strconv"
)

// Change the upload logic. The upload should be encrypted
// in the client and uploaded to the  IPFS API, and then
// the returned hash should be submitted to the hvxahv
// server for storage. hvxahv should not store the user's
// file key, only the saved hash value is used to format
// the storage list.

// Saves a file to the database.
type Saves struct {
	gorm.Model

	AccountID uint `gorm:"primaryKey;type:bigint;accounts_id"`

	// Name of the file.
	Name string `gorm:"type:text;name"`

	// Comments on the file.
	Description string `gorm:"type:text;description"`
	// Hash IPFS CID.
	Hash string `gorm:"type:text;hash"`

	// Types Is the file type identifier.
	Types string `gorm:"type:text;types"`
}

func (s *saved) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
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
	return &pb.CreateResponse{Code: "200", Reply: "ok"}, nil
}

func (s *saved) GetSaves(ctx context.Context, in *pb.GetSavesRequest) (*pb.GetSavesResponse, error) {
	db := cockroach.GetDB()
	id, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}
	var saves []*pb.Save
	if err := db.Debug().Table("saves").Where("account_id = ?", uint(id)).Find(&saves).Error; err != nil {
		return nil, err
	}
	return &pb.GetSavesResponse{Code: "200", Saves: saves}, nil
}

func (a *Saves) SetSavedName(name string) *Saves {
	a.Name = name
	return a
}

func (a *Saves) SetSavedDescription(description string) *Saves {
	a.Description = description
	return a
}

func (s *saved) GetSaved(ctx context.Context, in *pb.GetSavedRequest) (*pb.Save, error) {
	db := cockroach.GetDB()
	id, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}
	var saves Saves
	if err := db.Debug().Table("saves").Where("id = ?", uint(id)).First(&saves).Error; err != nil {
		return nil, err
	}
	return &pb.Save{
		Id:          strconv.Itoa(int(saves.ID)),
		Name:        saves.Name,
		Description: saves.Description,
		Hash:        saves.Hash,
		Types:       saves.Types,
	}, nil
}

func (s *saved) EditSaved(ctx context.Context, in *pb.EditSavedRequest) (*pb.EditSavedResponse, error) {
	save := new(Saves)
	if in.Name != "" {
		save.SetSavedName(in.Name)
	}
	if in.Description != "" {
		save.SetSavedDescription(in.Description)
	}

	db := cockroach.GetDB()
	id, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}
	if err := db.Debug().Table("saves").Where("id = ?", uint(id)).Updates(&save).Error; err != nil {
		return nil, err
	}
	return &pb.EditSavedResponse{Code: "200", Reply: "ok"}, nil
}

func NewSaves(accountID uint, name string, description string, hash string, types string) *Saves {
	return &Saves{AccountID: accountID, Name: name, Description: description, Hash: hash, Types: types}
}
