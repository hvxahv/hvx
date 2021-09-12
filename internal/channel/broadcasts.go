package channel

import (
	"fmt"
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/disism/hvxahv/pkg/ipfs"
	"github.com/disism/hvxahv/pkg/microservices/client"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
	"strings"
)

type Broadcasts struct {
	gorm.Model
	CID     uint   `gorm:"primarykey"`
	Author  string `gorm:"type:varchar(100);author"`
	Title   string `gorm:"type:varchar(999);article"`
	Article string `gorm:"type:varchar(3000);article"`
	IpfsCID string `gorm:"type:varchar(3000);ipfs_cid"`
}

func (b *Broadcasts) QueryLisByCID() (*[]Broadcasts, error) {
	db := cockroach.GetDB()

	var br []Broadcasts
	if err := db.Debug().Table("broadcasts").Where("c_id = ?", b.CID).Find(&br); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if ok {
			return nil, errors.Errorf("BROADCASTS_NOT_FOUND")
		}
	}
	return &br, nil
}

func (b *Broadcasts) New() error {

	broad := fmt.Sprintf(`
<!doctype html>
<html>
<head>
<meta charset='UTF-8'><meta name='viewport' content='width=device-width initial-scale=1'>
</style><title></title>
</head>
<body>
<div>
<p>AUTHOR: %s</p>
<h1>%s</h1>
<p>%s</p>
</div>
</body>
</html>
`, b.Author, b.Title, b.Article)
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
		CID:     b.CID,
		Author:  b.Author,
		Article: broad,
		IpfsCID: cid,
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

func NewBroadcast(title, article string, cid, aid uint) (*Broadcasts, error) {
	db := cockroach.GetDB()
	if err := db.Table("administrators").Where("c_id = ?", cid).Where("a_id = ?", aid).First(&Administrators{}); err != nil {
		if cockroach.IsNotFound(err.Error) {
			return nil, errors.Errorf("You are not the moderator of this channel")
		}
	}

	author, err := client.FetchAccountNameByID(aid)
	if err != nil {
		return nil, errors.Errorf("Failed to get author.")
	}
	return &Broadcasts{Author: author, Title: title, Article: article, CID: cid}, nil
}

func NewBroadcastCID(cid uint) *Broadcasts {
	return &Broadcasts{CID: cid}
}
