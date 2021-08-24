package remote

import (
	"bytes"
	"fmt"
	"github.com/disism/hvxahv/internal/accounts"
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/disism/hvxahv/pkg/security"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/url"
	"time"
)

func SendRequest(addr string, body []byte) {
	payload := bytes.NewBuffer(body)
	cli := &http.Client{}
	// addr end url
	req, err := http.NewRequest("POST", "https://mas.to/users/hvturingga/inbox", payload)
	if err != nil {
		fmt.Println(err)
	}

	h, err := url.Parse("https://mas.to/users/hvturingga/inbox")
	if err != nil {
		log.Fatal(err)
	}

	date := time.Now().UTC().Format(http.TimeFormat)
	fmt.Println( h.Hostname())
	req.Header.Add("Host", h.Hostname())
	req.Header.Add("Date", date)

	acct := &accounts.Accounts{}
	if err2 := cockroach.GetDB().
		Debug().
		Table("accounts").
		Where("username = ?", "hvturingga").
		First(acct).Error; err2 != nil {
		log.Println(gorm.ErrMissingWhereClause)
	}

	block := security.PrivateKey{
		Key: []byte(acct.PrivateKey),
	}
	uad := fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), "hvturingga")
	fmt.Println(uad)
	security.SignRequest(uad, block, req, body)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	req = req.WithContext(ctx)
	res, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(req)
	res.Body.Close()
	fmt.Println(res)
	fmt.Println(res.StatusCode)
}
