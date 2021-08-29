package httpsig

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/http"
	"strings"
)

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
