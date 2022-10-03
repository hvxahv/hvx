package delivery

import (
	"bytes"
	"crypto"
	"crypto/rsa"
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

const (
	FailedToDelivery     = "THERE_ARE_REQUESTS_THAT_FAILED_TO_SEND"
	DeliverySuccessfully = "SEND_SUCCESSFULLY"
)

// Delivery ...
// https://www.w3.org/TR/activitypub/#delivery
type Delivery struct {
	PublicKeyId string
	PrivateKey  string
	Body        []byte
}

func New(publicKeyId string, privateKey string, body []byte) *Delivery {
	return &Delivery{PublicKeyId: publicKeyId, PrivateKey: privateKey, Body: body}
}

func (i *Delivery) Do(inbox string) (*http.Response, error) {
	remote, err := IsRemote(inbox)
	if err != nil {
		return nil, errors.New("INBOX_URL_FAILS_WHEN_PARSING")
	}

	if !remote {
		// TODO - SEND TO LOCAL USER INBOX...
		return &http.Response{StatusCode: 202}, nil
	}

	hostname, err := parseInboxHostname(inbox)
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	singer, _, _ := httpsig.NewSigner([]httpsig.Algorithm{httpsig.RSA_SHA256}, "SHA-256", []string{"(request-target)", "date", "host", "digest"}, httpsig.Signature, 120)

	req, err := http.NewRequest("POST", inbox, bytes.NewReader(i.Body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/activity+json")
	req.Header.Add("Content-Type", "application/ld+json")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Date", time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05")+" GMT")
	req.Header.Add("User-Agent", "hvxahv/0.0.0")
	req.Header.Add("Host", hostname)
	req.Header.Add("Accept", "application/activity+json; charset=utf-8")

	if err := singer.SignRequest(GetPrivateKey(i.PrivateKey), i.PublicKeyId, req, i.Body); err != nil {
		return nil, err
	}
	if err := Verify(req, i.PrivateKey); err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
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

// IsRemote ...
func IsRemote(inbox string) (bool, error) {
	domain := viper.GetString("domain")

	hostname, err := parseInboxHostname(inbox)
	if err != nil {
		return false, err
	}
	if domain != hostname {
		return true, nil
	}
	return false, nil
}

// GetPrivateKey ParseRawPrivateKey -> crypto.PrivateKey.
func GetPrivateKey(privateKey string) crypto.PrivateKey {
	var cpk crypto.PrivateKey
	key, err := ssh.ParseRawPrivateKey([]byte(privateKey))
	cpk = key.(*rsa.PrivateKey)
	if err != nil {
		errors.Throw("", err)
	}
	return cpk
}

// Verify ...
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
