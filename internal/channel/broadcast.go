package channel

import (
	"context"
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/channel/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"gorm.io/gorm"
	"strconv"
)

const (
	BroadcastsTableName = "broadcasts"
)

type Broadcasts struct {
	gorm.Model
	ChannelID uint   `gorm:"primaryKey;type:bigint;channel_id"`
	AdminID   uint   `gorm:"type:bigint;admin_id"`
	ArticleID uint   `gorm:"type:bigint;article_id"`
	CID       string `gorm:"type:text;cid"`
}

func (c *channel) CreateBroadcast(ctx context.Context, in *pb.CreateBroadcastRequest) (*pb.CreateBroadcastResponse, error) {
	administrator, err := c.IsChannelAdministrator(ctx, &pb.IsChannelAdministratorRequest{
		ChannelId: in.ChannelId,
		AdminId:   in.AdminId,
	})
	if err != nil {
		return nil, err
	}
	if !administrator.IsAdministrator {
		return nil, fmt.Errorf("%s", NotAdmin)
	}
	// TODO: sync with ipfs.
	hash := "IPFS_CID"

	cid, err := strconv.Atoi(in.GetChannelId())
	if err != nil {
		return nil, err
	}
	aid, err := strconv.Atoi(in.GetAdminId())
	if err != nil {
		return nil, err
	}
	article, err := strconv.Atoi(in.GetArticleId())
	if err != nil {
		return nil, err
	}

	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Broadcasts{}); err != nil {
		return nil, err
	}
	if err := db.Debug().Table(BroadcastsTableName).Create(&Broadcasts{
		ChannelID: uint(cid),
		AdminID:   uint(aid),
		ArticleID: uint(article),
		CID:       hash,
	}).Error; err != nil {
		return nil, err
	}

	return &pb.CreateBroadcastResponse{Code: "200", Reply: "ok"}, nil
}

func (c *channel) GetAllBroadcasts(ctx context.Context, in *pb.GetAllBroadcastsRequest) (*pb.GetAllBroadcastsResponse, error) {
	cid, err := strconv.Atoi(in.GetChannelId())
	if err != nil {
		return nil, err
	}
	var b []Broadcasts
	db := cockroach.GetDB()
	if err := db.Debug().Table(BroadcastsTableName).Where("channel_id = ?", uint(cid)).Find(&b).Error; err != nil {
		return nil, err
	}
	var reply []*pb.Broadcast
	for _, v := range b {
		reply = append(reply, &pb.Broadcast{
			Id:        strconv.Itoa(int(v.ID)),
			ChannelId: strconv.Itoa(int(v.ChannelID)),
			AdminId:   strconv.Itoa(int(v.AdminID)),
			ArticleId: strconv.Itoa(int(v.ArticleID)),
			Cid:       v.CID,
		})
	}
	return &pb.GetAllBroadcastsResponse{Code: "200", Broadcasts: reply}, nil
}

func (c *channel) DeleteBroadcast(ctx context.Context, in *pb.DeleteBroadcastRequest) (*pb.DeleteBroadcastResponse, error) {
	administrator, err := c.IsChannelAdministrator(ctx, &pb.IsChannelAdministratorRequest{
		ChannelId: in.ChannelId,
		AdminId:   in.AdminId,
	})
	if err != nil {
		return nil, err
	}
	if !administrator.IsAdministrator {
		return nil, fmt.Errorf("%s", NotAdmin)
	}
	bid, err := strconv.Atoi(in.GetBroadcastId())
	if err != nil {
		return nil, err
	}
	cid, err := strconv.Atoi(in.GetChannelId())
	if err != nil {
		return nil, err
	}

	db := cockroach.GetDB()
	if err := db.Debug().
		Table(BroadcastsTableName).
		Where("id = ? AND channel_id = ?", uint(bid), uint(cid)).
		Unscoped().
		Delete(&Broadcasts{}).
		Error; err != nil {
		return nil, err
	}

	return &pb.DeleteBroadcastResponse{Code: "200", Reply: "ok"}, nil
}
