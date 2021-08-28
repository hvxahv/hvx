package httpsig

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

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