package security

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// Signing requests using HTTP Signatures.
// https://www.w3.org/wiki/SocialCG/ActivityPub/Authentication_Authorization

// When communicating over the Internet using the HTTP protocol, it can
// be desirable for a server or client to authenticate the sender of a
// particular message.  It can also be desirable to ensure that the
// message was not tampered with during transit.  This document
// describes a way for servers and clients to simultaneously add
// authentication and message integrity to HTTP activity by using a
// digital signature.
//https://datatracker.ietf.org/doc/html/draft-cavage-http-signatures-08

type KeyType int

const (
	None KeyType = iota
	RSA
	Ed25519
)

type PublicKey struct {
	Type KeyType
	Key  interface{}
}

type PriKEY struct {
	Type KeyType
	Key  []byte
}

type sign struct {
	KeyID string
	Key   PriKEY
	Req   *http.Request
	Data  []byte
}

type verify struct {
	Req  *http.Request
	Data []byte
	Key  PublicKey
}

func NewVerify(req *http.Request, data []byte, key PublicKey) *verify {
	return &verify{Req: req, Data: data, Key: key}
}

func NewSign(keyName string, key PriKEY, request *http.Request, data []byte) *sign {
	return &sign{KeyID: keyName, Key: key, Req: request, Data: data}
}

func BS642BYTE(s string) []byte {
	var buf bytes.Buffer
	b64 := base64.NewDecoder(base64.StdEncoding, strings.NewReader(s))
	io.Copy(&buf, b64)
	return buf.Bytes()
}

func BS64(data []byte) string {
	var builder strings.Builder
	b64 := base64.NewEncoder(base64.StdEncoding, &builder)
	b64.Write(data)
	b64.Close()
	return builder.String()
}

func BS642SHA256(data []byte) string {
	h := sha256.New()
	h.Write(data)
	return BS64(h.Sum(nil))
}


func (key *PriKEY) Sign(msg []byte) []byte {
	block, _ := pem.Decode(key.Key)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatalf("parses an RSA private key error: %v", err)
	}
	sig, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, msg)
	if err != nil {
		log.Fatalf("SIGN PKCS1v15 error : %v", err)
	}

	return sig
}

// SignRequest ...
func (si *sign) SignRequest() {
	headers := []string{"(request-target)", "date", "host", "content-type", "digest"}
	value := BuildSignatureData(si.Req, headers, si.Data)

	h := sha256.New()
	h.Write(value)
	sig := si.Key.Sign(h.Sum(nil))

	si.Req.Header.Set("Signature", bindSignature(headers, sig, si.KeyID))
}

func bindSignature(headers []string, sig []byte, id string) string {
	return fmt.Sprintf(`keyId="%s",algorithm="%s",headers="%s",signature="%s"`,
		id,
		"rsa-sha256",
		strings.Join(headers, " "),
		BS64(sig),
	)
}

func requestPath(req *http.Request) string {
	path := req.URL.Path
	if path == "" {
		path = "/"
	}
	if req.URL.RawQuery != "" {
		path += "?" + req.URL.RawQuery
	}
	return path
}

func BuildSignatureHeader(req *http.Request, headers []string, data []byte) string {
	if len(headers) == 0 {
		headers = []string{"date"}
	}
	values := make([]string, 0, len(headers))
	for _, h := range headers {
		switch h {
		case "(request-target)":
			values = append(values, fmt.Sprintf("%s: %s %s",
				h, strings.ToLower(req.Method), requestPath(req)))
		case "host":
			values = append(values, fmt.Sprintf("%s: %s", h, req.Host))
		case "date":
			if req.Header.Get(h) == "" {
				req.Header.Set(h, time.Now().UTC().Format(http.TimeFormat))
			}
			values = append(values, fmt.Sprintf("%s: %s", h, req.Header.Get(h)))
		case "digest":
			if req.Header.Get(h) == "" {
				req.Header.Set(h, "SHA-256=" + BS642SHA256(data))
			}
			values = append(values, fmt.Sprintf("%s: %s", h, req.Header.Get(h)))
		default:
			for _, value := range req.Header[http.CanonicalHeaderKey(h)] {
				values = append(values,
					fmt.Sprintf("%s: %s", h, strings.TrimSpace(value)))
			}
		}
	}

	return strings.Join(values, "\n")
}

func BuildSignatureData(req *http.Request, headers []string, data []byte) []byte {
	return []byte(BuildSignatureHeader(req, headers, data))
}
