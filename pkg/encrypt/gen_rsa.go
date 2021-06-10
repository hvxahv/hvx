package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/pkg/errors"
)

// GenRSA The method calls the GenerateKey function to generate a 2068-bit ras key
// to return the public key and private key.
func GenRSA() (string, string, error) {
	privateKey, publicKey, err := GenerateKey(2048)
	if err != nil {
		return "", "", errors.Errorf("Generate key is error: %v", err)
	}
	private := EncodePrivateKey(privateKey)
	public, err := EncodePublicKey(publicKey)
	if err != nil {
		return "", "", errors.Errorf("Encode Public Key is error: %v", err)
	}

	return string(private), string(public), err
}

// GenerateKey Generate key pair.
func GenerateKey(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	private, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return private, &private.PublicKey, nil
}

// EncodePrivateKey Encode Private Key.
func EncodePrivateKey(private *rsa.PrivateKey) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Bytes: x509.MarshalPKCS1PrivateKey(private),
		Type:  "RSA PRIVATE KEY",
	})
}

// EncodePublicKey Encode Public Key.
func EncodePublicKey(public *rsa.PublicKey) ([]byte, error) {
	publicBytes, err := x509.MarshalPKIXPublicKey(public)
	if err != nil {
		return nil, err
	}
	return pem.EncodeToMemory(&pem.Block{
		Bytes: publicBytes,
		Type:  "PUBLIC KEY",
	}), nil
}
