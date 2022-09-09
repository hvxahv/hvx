package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

type R interface {
	Generate() (*Rsa, error)
}
type Rsa struct {
	bits       int
	PrivateKey string
	PublicKey  string
}

func NewRsa(bits int) *Rsa {
	return &Rsa{
		bits: bits,
	}
}

func (r *Rsa) Generate() (*Rsa, error) {
	pvk, err := rsa.GenerateKey(rand.Reader, r.bits)
	if err != nil {
		return nil, fmt.Errorf("generates an RSA keypair error: %v", err)
	}

	derStream := x509.MarshalPKCS1PrivateKey(pvk)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prv := pem.EncodeToMemory(block)

	publicKey := &pvk.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, fmt.Errorf("converts a public key error: %v", err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pub := pem.EncodeToMemory(block)

	return &Rsa{
		PrivateKey: string(prv),
		PublicKey:  string(pub),
	}, nil
}
