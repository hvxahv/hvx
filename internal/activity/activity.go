package activity

import (
	"fmt"
	"github.com/disism/hvxahv/internal/accounts"
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"net/url"
)

type ActivityRequest struct {
	KeyID     string
	TargetURL string
	Local     string
	Data      []byte
	Key       []byte
}

type Request interface {
	// Send request to remote server.
	Send()

	// Follow ActivityPub follow method.
	Follow()

	// Accept ... TODO - Implement the method...
	Accept()
}


type InboxWithCtx struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Actor  string `json:"actor"`
	Object string `json:"object"`
}
type Receive interface {
	Inbox(name string)
}

func getPrivk() string {
	acct := &accounts.Accounts{}
	if err2 := cockroach.GetDB().
		Debug().
		Table("accounts").
		Where("username = ?", "hvturingga").
		First(acct).Error; err2 != nil {
		log.Println(gorm.ErrMissingWhereClause)
	}
	return acct.PrivateKey
}

// NewActivityRequest Receive the current actor name,
// the other party's URL,
// the requested data and the current user's private key.
func NewActivityRequest(actor string, object string, data []byte, key []byte) *ActivityRequest {
	h, err := url.Parse(object)
	if err != nil {
		log.Fatal(err)
	}

	targetURL := fmt.Sprintf("https://%s/inbox", h.Hostname())
	keyID := fmt.Sprintf("%s#main-key", actor)

	return &ActivityRequest{
		KeyID:     keyID,
		TargetURL: targetURL,
		Local:     fmt.Sprintf(viper.GetString("localhost")),
		Data:      data,
		Key:       key,
	}
}

func (a *ActivityRequest) Follow() {
	a.Send()
}


func (a *ActivityRequest) Accept() {
	a.Send()
}

func (a *ActivityRequest) Create() {
	a.Send()
}