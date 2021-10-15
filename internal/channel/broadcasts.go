package channel

import (
	"fmt"
	pb "github.com/disism/hvxahv/api/accounts/v1alpha1"
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/disism/hvxahv/pkg/ipfs"
	"github.com/disism/hvxahv/pkg/microservices/client"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"log"
	"strings"
)

type Broadcasts struct {
	gorm.Model
	ChannelID uint   `gorm:"primarykey;channel_id"`
	AuthorID  uint   `gorm:"type:bigint;author_id`
	Title     string `gorm:"type:text;article"`
	Summary   string `gorm:"type:text;summary"`
	Article   string `gorm:"type:text;article"`
	NSFW      bool   `gorm:"type:boolean;nsfw"`
	IpfsCID   string `gorm:"type:text;ipfs_cid"`
}

func (b *Broadcasts) QueryLisByCID() (*[]Broadcasts, error) {
	db := cockroach.GetDB()

	var br []Broadcasts
	if err := db.Debug().Table("broadcasts").Where("channel_id = ?", b.ChannelID).Find(&br); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if ok {
			return nil, errors.Errorf("BROADCASTS_NOT_FOUND")
		}
	}
	return &br, nil
}

func (b *Broadcasts) New() error {

	cli, conn, err := client.Accounts()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	r, err := cli.FindActorByID(context.Background(), &pb.ActorID{
		ActorID: uint64(b.AuthorID),
	})
	if err != nil {
		log.Printf("failed to send message to accounts server: %v", err)
	}
	author := r.PreferredUsername

	broad := fmt.Sprintf(`
<!doctype html>
<html>
<head>
<meta charset='UTF-8'><meta name='viewport' content='width=device-width initial-scale=1'>
</style><title></title>
</head>
<body>
<div>
<h1>%s</h1>
<p>%s</p>
<p>%s</p>
</div>
</body>
</html>
`, b.Title, author, b.Article)
	cid, err := ipfs.GetIPFS().Add(strings.NewReader(broad))
	if err != nil {
		fmt.Printf("ipfs add error: %v", err)
	}

	// Save data and cid to database.
	db := cockroach.GetDB()
	err2 := db.AutoMigrate(Broadcasts{})
	if err2 != nil {
		return nil
	}

	data := Broadcasts{
		ChannelID: b.ChannelID,
		AuthorID:  b.AuthorID,
		Article:   broad,
		IpfsCID:   cid,
	}

	if err1 := db.Debug().Table("broadcasts").Create(&data).Error; err1 != nil {
		log.Printf("an error occurred while creating the broadcasts: %v", err)
	}

	fmt.Printf("NEW BROADCASTS SUCCESS, CID = %s", cid)
	return nil
}

type Broadcast interface {
	// New Create broadcast Articles.
	// Synchronize to ipfs return ipfs id.
	New() error

	// QueryLisByCID Fetch the content list in the channel by channel id.
	QueryLisByCID() (*[]Broadcasts, error)
}

func NewBroadcast(title, article string, channelID, actorID uint) (*Broadcasts, error) {
	db := cockroach.GetDB()
	if err := db.Table("administrators").Where("channel_id = ?", channelID).Where("actor_id = ?", actorID).First(&Administrators{}); err != nil {
		if cockroach.IsNotFound(err.Error) {
			return nil, errors.Errorf("You are not the moderator of this channel")
		}
	}

	return &Broadcasts{AuthorID: actorID, Title: title, Article: article, ChannelID: channelID}, nil
}

func NewBroadcastCID(channelId uint) *Broadcasts {
	return &Broadcasts{ChannelID: channelId}
}
