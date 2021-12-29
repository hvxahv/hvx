package channels

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hvxahv/hvxahv/internal/accounts"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/hvxahv/hvxahv/pkg/ipfs"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

type Broadcasts struct {
	gorm.Model
	ChannelID uint   `gorm:"primaryKey;type:bigint;channel_id"`
	ActorID   uint   `gorm:"type:bigint;actor_id"`
	CID       string `gorm:"index;type:text;cid"`
}

type BroadcastsIPFSData struct {
	PreferredUsername string `json:"preferred_username"`
	Avatar            string `json:"avatar"`
	URL               string `json:"url"`
	Title             string `json:"title"`
	Summary           string `json:"summary"`
	Article           string `json:"article"`
	NSFW              bool   `json:"nsfw"`
}

func (b *Broadcasts) GetBroadcastsByChannelID() (*[]Broadcasts, error) {
	db := cockroach.GetDB()

	var r []Broadcasts
	if err := db.Debug().Table("broadcasts").Where("channel_id = ?", b.ChannelID).Find(&r); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if ok {
			return nil, errors.Errorf("BROADCASTS_NOT_FOUND")
		}
	}
	return &r, nil
}

func (b *Broadcasts) Create() error {
	// Save data and cid to database.
	db := cockroach.GetDB()
	err2 := db.AutoMigrate(Broadcasts{})
	if err2 != nil {
		return nil
	}

	bc := &Broadcasts{
		ChannelID: b.ChannelID,
		ActorID:   b.ActorID,
		CID:       b.CID,
	}

	if err := db.Debug().Table("broadcasts").Create(&bc).Error; err != nil {
		log.Printf("an error occurred while creating the broadcasts: %v", err)
	}

	fmt.Printf("NEW BROADCASTS SUCCESS, CID = %s", b.CID)
	return nil
}

type Broadcast interface {
	// Create broadcast Articles.
	// Synchronize to ipfs return ipfs id.
	Create() error

	// GetBroadcastsByChannelID Fetch the content list in the channels by channel id.
	GetBroadcastsByChannelID() (*[]Broadcasts, error)
}

func NewBroadcasts(channelID uint, actorID uint, cid string) *Broadcasts {
	return &Broadcasts{ChannelID: channelID, ActorID: actorID, CID: cid}
}

// NewBroadcastsIPFSCID Upload Broadcasts to IPFS and return CID.
func NewBroadcastsIPFSCID(channelID, actorID uint, title, summary, article string, NSFW bool) (string, error) {
	if err := NewAdministrators(channelID, actorID).IsAdmin(); err != nil {
		return "", err
	}

	actor, err := accounts.NewActorID(actorID).GetByActorID()
	if err != nil {
		return "", err
	}

	broad := BroadcastsIPFSData{
		PreferredUsername: actor.PreferredUsername,
		Avatar:            actor.Avatar,
		URL:               actor.Url,
		Title:             title,
		Summary:           summary,
		Article:           article,
		NSFW:              NSFW,
	}
	data, err := json.Marshal(broad)
	if err != nil {
		return "", err
	}
	cid, err := ipfs.GetIPFS().Add(bytes.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("ipfs add error: %v", err)
	}

	return cid, nil
}

func NewBroadcastsChannelID(channelId uint) *Broadcasts {
	return &Broadcasts{ChannelID: channelId}
}
