package activity

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"fmt"
	"github.com/go-fed/httpsig"
	"github.com/hvxahv/hvx/errors"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// Sender Server to Server Interactions
type Sender struct {
	actor      string
	body       []byte
	privateKey string
}

func NewSender(body []byte, actor, privateKey string) *Sender {
	return &Sender{actor: actor, body: body, privateKey: privateKey}
}

func (i *Sender) Do(inbox string) (*http.Response, error) {
	hostname, err := parseInboxHostname(inbox)
	if err != nil {
		fmt.Println(err)
	}
	client := http.Client{}
	
	singer, _, _ := httpsig.NewSigner([]httpsig.Algorithm{httpsig.RSA_SHA256}, "SHA-256", []string{"(request-target)", "date", "host", "digest"}, httpsig.Signature, 120)

	req, err := http.NewRequest("POST", inbox, bytes.NewReader(i.body))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/activity+json")
	req.Header.Add("Content-Type", "application/ld+json")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Date", time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05")+" GMT")
	req.Header.Add("User-Agent", "hvxahv/0.0.0")
	req.Header.Add("Host", hostname)
	req.Header.Add("Accept", "application/activity+json; charset=utf-8")

	if err := singer.SignRequest(GetPrivateKey(i.privateKey), fmt.Sprintf("%s#main-key", i.actor), req, i.body); err != nil {
		fmt.Println(err)
	}
	if err := Verify(req, i.privateKey); err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer res.Body.Close()
	stdout := os.Stdout
	_, err = io.Copy(stdout, res.Body)

	return res, err
}

func parseInboxHostname(inbox string) (string, error) {
	u, err := url.Parse(inbox)
	if err != nil {
		return "", err
	}
	hostname := strings.TrimPrefix(u.Hostname(), "www.")
	return hostname, nil
}

func IsLocalInstance(inbox string) (bool, error) {
	domain := viper.GetString("domain")

	hostname, err := parseInboxHostname(inbox)
	if err != nil {
		return false, err
	}
	if domain != hostname {
		return false, nil
	}
	return true, nil
}

func GetPrivateKey(privateKey string) crypto.PrivateKey {
	var cpk crypto.PrivateKey
	key, err := ssh.ParseRawPrivateKey([]byte(privateKey))
	cpk = key.(*rsa.PrivateKey)
	if err != nil {
		errors.Throw("", err)
	}
	return cpk
}

func Verify(r *http.Request, privateKey string) error {
	verifier, err := httpsig.NewVerifier(r)
	if err != nil {
		return err
	}

	priv := GetPrivateKey(privateKey)
	pub := priv.(*rsa.PrivateKey).PublicKey
	var algo httpsig.Algorithm = httpsig.RSA_SHA256
	var pubKey crypto.PublicKey = &pub

	return verifier.Verify(pubKey, algo)
}
