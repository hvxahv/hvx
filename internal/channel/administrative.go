package channel

import (
	"context"
	"strconv"

	pb "github.com/hvxahv/hvxahv/api/channel/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Administrates struct {
	gorm.Model

	ChannelID uint `gorm:"primaryKey;channel_id"`
	AdminID   uint `gorm:"primaryKey;admin_id"`
	IsOwner   bool `gorm:"type:boolean;is_owner"`
}

const (
	// AdministrateTable is the table name for the administrates table.
	AdministrateTable = "administrates"
	NotAdmin          = "NOT_AN_ADMINISTRATOR"
	AlreadyAdmin      = "ADMINISTRATOR_ALREADY_EXISTS"
	NotFound          = "ADMINISTRATOR_NOT_FOUND"
)

func (c *channel) IsChannelAdministrator(ctx context.Context, in *pb.IsChannelAdministratorRequest) (*pb.IsChannelAdministratorResponse, error) {
	db := cockroach.GetDB()
	aid, err := strconv.Atoi(in.GetAccountId())
	if err != nil {
		return nil, err
	}
	cid, err := strconv.Atoi(in.GetChannelId())
	if err != nil {
		return nil, err
	}
	if err := db.Debug().
		Table("channels").
		Where("id = ? AND account_id = ?", uint(cid), uint(aid)).
		First(&c.Channels); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if !ok {
			return &pb.IsChannelAdministratorResponse{IsAdministrator: true}, nil
		}
	}
	return &pb.IsChannelAdministratorResponse{IsAdministrator: false}, nil
}

func (c *channel) AddAdministrator(ctx context.Context, in *pb.AddAdministratorRequest) (*pb.AddAdministratorResponse, error) {
	administrator, err := c.IsChannelAdministrator(ctx, &pb.IsChannelAdministratorRequest{
		AccountId: in.GetAdminAccountId(),
		ChannelId: in.GetChannelId(),
	})
	if err != nil {
		return nil, err
	}
	if !administrator.IsAdministrator {
		return nil, errors.New(NotAdmin)
	}
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Administrates{}); err != nil {
		return nil, err
	}
	aid, err := strconv.Atoi(in.GetAddAdminId())
	if err != nil {
		return nil, err
	}
	cid, err := strconv.Atoi(in.GetChannelId())
	if err != nil {
		return nil, err
	}
	adm := &Administrates{
		ChannelID: uint(cid),
		AdminID:   uint(aid),
		IsOwner:   in.IsOwner,
	}

	if err := db.Debug().
		Table(AdministrateTable).
		Where("channel_id = ? AND admin_id = ?", uint(cid), uint(aid)).
		First(&adm); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if !ok {
			return nil, errors.New(AlreadyAdmin)
		}
	}
	if err := db.Debug().
		Table(AdministrateTable).
		Create(adm).Error; err != nil {
		return nil, err
	}
	return &pb.AddAdministratorResponse{Code: "200", Reply: "ok"}, nil
}

func (c *channel) RemoveAdministrator(ctx context.Context, in *pb.RemoveAdministratorRequest) (*pb.RemoveAdministratorResponse, error) {
	s := &channel{}
	administrator, err := s.IsChannelAdministrator(ctx, &pb.IsChannelAdministratorRequest{
		ChannelId: in.GetChannelId(),
		AccountId: in.GetOwnerId(),
	})
	if err != nil {
		return nil, err
	}
	if !administrator.IsAdministrator {
		return &pb.RemoveAdministratorResponse{Code: "401", Reply: NotAdmin}, nil
	}

	db := cockroach.GetDB()

	aid, err := strconv.Atoi(in.GetRemoveAdminId())
	if err != nil {
		return nil, err
	}
	cid, err := strconv.Atoi(in.GetChannelId())
	if err != nil {
		return nil, err
	}

	if err := db.Debug().
		Table(AdministrateTable).
		Where("channel_id = ? AND admin_id = ?", uint(cid), uint(aid)).First(&Administrates{}).
		Unscoped().
		Delete(&c.Administrates); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if ok {
			return nil, errors.New(NotFound)
		}
	}
	return &pb.RemoveAdministratorResponse{Code: "200", Reply: "ok"}, nil
}

func (c *channel) GetAdministrators(ctx context.Context, in *pb.GetAdministratorsRequest) (*pb.GetAdministratorsResponse, error) {
	s := &channel{}
	administrator, err := s.IsChannelAdministrator(ctx, &pb.IsChannelAdministratorRequest{
		ChannelId: in.GetChannelId(),
		AccountId: in.GetAccountId(),
	})
	if err != nil {
		return nil, err
	}
	if !administrator.IsAdministrator {
		return &pb.GetAdministratorsResponse{Code: "401", Administrators: nil}, nil
	}
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Administrates{}); err != nil {
		return nil, err
	}
	cid, err := strconv.Atoi(in.GetChannelId())
	if err != nil {
		return nil, err

	}
	var admins []Administrates
	if err := db.Debug().
		Table(AdministrateTable).
		Where("channel_id = ?", uint(cid)).
		Find(&admins).Error; err != nil {
		return nil, err
	}
	var accountIds []string
	for _, adm := range admins {
		accountIds = append(accountIds, strconv.Itoa(int(adm.AdminID)))
	}
	return &pb.GetAdministratorsResponse{Code: "200", Administrators: accountIds}, nil
}
