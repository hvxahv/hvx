package internal

import (
	"github.com/hvxahv/hvx/cockroach"
	"github.com/hvxahv/hvx/errors"

	"gorm.io/gorm"
)

const (
	BroadcastsTableName = "broadcasts"
)

type Broadcasts struct {
	gorm.Model

	ChannelId uint   `gorm:"primaryKey;type:bigint;channel_id"`
	AdminId   uint   `gorm:"type:bigint;admin_id"`
	ArticleId uint   `gorm:"type:bigint;article_id"`
	CID       string `gorm:"type:text;cid"`
}

type Broadcast interface {
	CreateBroadcast() error
	GetBroadcasts() (*[]Broadcasts, error) // Get all broadcasts.
	DeleteBroadcast() error
}

func NewBoardcast(channelId, adminId, articleId uint, cid string) *Broadcasts {
	return &Broadcasts{ChannelId: channelId, AdminId: adminId, ArticleId: articleId, CID: cid}
}

func (b *Broadcasts) CreateBroadcast() error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Broadcasts{}); err != nil {
		return errors.NewDatabaseCreate("broadcasts")
	}
	if err := db.Debug().
		Table(BroadcastsTableName).
		Create(&b).
		Error; err != nil {
		return err
	}

	return nil
}

func NewBoardcastsChannelId(channelId uint) *Broadcasts {
	return &Broadcasts{ChannelId: channelId}
}

func (b *Broadcasts) GetBroadcasts() (*[]Broadcasts, error) {
	db := cockroach.GetDB()
	var broadcasts []Broadcasts

	if err := db.Debug().
		Table(BroadcastsTableName).
		Where("channel_id = ?", b.ChannelId).
		Find(&broadcasts).
		Error; err != nil {
		return nil, err
	}

	return &broadcasts, nil
}

func NewDeleteBroadcast(id, channelId, adminId uint) *Broadcasts {
	return &Broadcasts{
		Model:     gorm.Model{ID: id},
		ChannelId: channelId,
		AdminId:   adminId,
	}
}

func (b *Broadcasts) DeleteBroadcast() error {
	isAdmin := NewAdministratesPermission(b.AdminId, b.ChannelId).IsAdministrator()
	if !isAdmin {
		return errors.New(errors.ErrNotAchannelAdministrator)
	}

	db := cockroach.GetDB()
	if err := db.Debug().
		Table(BroadcastsTableName).
		Where("id = ? AND channel_id = ?", b.ID, b.ChannelId).
		Unscoped().
		Delete(&Broadcasts{}).
		Error; err != nil {
		return err
	}

	return nil
}
