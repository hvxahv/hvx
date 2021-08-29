package channel

import (
	"fmt"
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/disism/hvxahv/pkg/ipfs"
	"log"
	"strings"
)

type Broadcasts struct {
	Author  string `gorm:"type:varchar(100);author"`
	Article string `gorm:"type:varchar(3000);article"`
	Cid     string `gorm:"type:varchar(3000);cid"`
}

func (b *Broadcasts) New() {

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
<p>%s</p>
</div>
</body>
</html>
`, b.Author, b.Article)
	cid, err := ipfs.GetIPFS().Add(strings.NewReader(broad))
	if err != nil {
		fmt.Printf("ipfs add error: %v", err)
	}

	// Save data and cid to database.
	db := cockroach.GetDB()
	err2 := db.AutoMigrate(Broadcasts{})
	if err2 != nil {
		return 
	}

	data := Broadcasts{
		Author:  b.Author,
		Article: broad,
		Cid:     cid,
	}

	if err1 := db.Debug().Table("broadcasts").Create(&data).Error; err1 != nil {
		log.Printf("an error occurred while creating the broadcasts: %v", err)
	}

	fmt.Printf("NEW BROADCASTS SUCCESS, CID = %s", cid)
}

func NewBroadcast(author string, article string) *Broadcasts {
	return &Broadcasts{Author: author, Article: article}
}

type Broadcast interface {
	// New Create broadcast Articles.
	// Synchronize to ipfs return url.
	New()
}
