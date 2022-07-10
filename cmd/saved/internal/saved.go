/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package internal

import (
	"github.com/hvxahv/hvx/cockroach"
	"gorm.io/gorm"
)

// Change the upload logic. The upload should be encrypted
// in the cfg and uploaded to the  IPFS API, and then
// the returned cid should be submitted to the hvxahv
// server for storage. hvxahv should not store the user's
// file key, only the saved cid value is used to format
// the storage list.

// Saves a file to the database.
type Saves struct {
	gorm.Model

	AccountId uint `gorm:"primaryKey;type:bigint;account_id"`

	// Name of the file.
	Name string `gorm:"type:text;name"`

	// Comments on the file.
	Description string `gorm:"type:text;description"`
	// Cid IPFS CID.
	Cid string `gorm:"type:text;cid"`

	// Types Is the file type identifier.
	Types string `gorm:"type:text;types"`
}

type Saved interface {
	Create() error
}

func NewSaves(accountID uint, name string, description string, cid string, types string) Saves {
	return Saves{
		AccountId:   accountID,
		Name:        name,
		Description: description,
		Cid:         cid,
		Types:       types,
	}
}
func (s *Saves) Create() error {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Saves{}); err != nil {
		return err
	}
	c := NewSaves(s.AccountId, s.Name, s.Description, s.Cid, s.Types)
	if err := db.Debug().Table("saves").Create(&c).Error; err != nil {
		return err
	}
	return nil
}

//
//func (s *Saves) GetSaves(ctx context.Context, in *pb.GetSavesRequest) (*pb.GetSavesResponse, error) {
//	db := cockroach.GetDB()
//	id, err := strconv.Atoi(in.AccountId)
//	if err != nil {
//		return nil, err
//	}
//
//	var saves []*pb.Save
//	if err := db.Debug().Table("saves").Where("account_id = ?", uint(id)).Find(&saves).Error; err != nil {
//		return nil, err
//	}
//	return &pb.GetSavesResponse{Code: "200", Saves: saves}, nil
//}
//
//func (s *Saves) SetSavedName(name string) *Saves {
//	s.Name = name
//	return s
//}
//
//func (s *Saves) SetSavedDescription(description string) *Saves {
//	s.Description = description
//	return s
//}
//
//func (s *Saves) GetSaved(ctx context.Context, in *pb.GetSavedRequest) (*pb.Save, error) {
//	db := cockroach.GetDB()
//	id, err := strconv.Atoi(in.Id)
//	if err != nil {
//		return nil, err
//	}
//	aid, err := strconv.Atoi(in.AccountId)
//	if err != nil {
//		return nil, err
//	}
//
//	var saves Saves
//	if err := db.Debug().
//		Table("saves").
//		Where("id = ? AND account_id = ?", uint(id), uint(aid)).
//		First(&saves).
//		Error; err != nil {
//		return nil, err
//	}
//
//	return &pb.Save{
//		Id:          strconv.Itoa(int(saves.ID)),
//		Name:        saves.Name,
//		Description: saves.Description,
//		Cid:         saves.Cid,
//		Types:       saves.Types,
//	}, nil
//}
//
//func (s *Saves) EditSaved(ctx context.Context, in *pb.EditSavedRequest) (*pb.EditSavedResponse, error) {
//	save := new(Saves)
//	if in.Name != "" {
//		save.SetSavedName(in.Name)
//	}
//	if in.Description != "" {
//		save.SetSavedDescription(in.Description)
//	}
//
//	db := cockroach.GetDB()
//	id, err := strconv.Atoi(in.Id)
//	if err != nil {
//		return nil, err
//	}
//	aid, err := strconv.Atoi(in.AccountId)
//	if err != nil {
//		return nil, err
//	}
//
//	if err := db.Debug().
//		Table("saves").
//		Where("id = ? AND account_id = ?", uint(id), uint(aid)).
//		Updates(&save).
//		Error; err != nil {
//		return nil, err
//	}
//	return &pb.EditSavedResponse{Code: "200", Reply: "ok"}, nil
//}
//
//func (s *Saves) DeleteSaved(ctx context.Context, in *pb.DeleteSavedRequest) (*pb.DeleteSavedResponse, error) {
//	db := cockroach.GetDB()
//	id, err := strconv.Atoi(in.Id)
//	if err != nil {
//		return nil, err
//	}
//
//	aid, err := strconv.Atoi(in.AccountId)
//	if err != nil {
//		return nil, err
//	}
//
//	if err := db.Debug().
//		Table("saves").
//		Where("id = ? AND account_id = ?", uint(id), uint(aid)).
//		Unscoped().
//		Delete(&Saves{}).
//		Error; err != nil {
//		return nil, err
//	}
//	return &pb.DeleteSavedResponse{Code: "200", Reply: "ok"}, nil
//}
//
//func (s *Saves) DeleteAllSaves(ctx context.Context, in *pb.DeleteAllSavesRequest) (*pb.DeleteAllSavesResponse, error) {
//	db := cockroach.GetDB()
//	id, err := strconv.Atoi(in.AccountId)
//	if err != nil {
//		return nil, err
//	}
//
//	if err := db.Debug().
//		Table("saves").
//		Where("account_id = ?", uint(id)).
//		Unscoped().
//		Delete(&Saves{}).
//		Error; err != nil {
//		return nil, err
//	}
//	return &pb.DeleteAllSavesResponse{Code: "200", Reply: "ok"}, nil
//}
//
//func NewSaves(accountID uint, name string, description string, cid string, types string) *Saves {
//	return &Saves{AccountID: accountID, Name: name, Description: description, Cid: cid, Types: types}
//}
