package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// GenPassword Use the bcrypt package to crypto the password and return the encrypted hash,
// which needs to be converted into a string.
func GenPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Encryption password failed: ", err)
	}
	return string(hash)
}

// ***************************************************

// GenRSA The method calls the GenerateKey function to generate a 2068-bit ras key
// to return the public key and private key.
func GenRSA() (string, string, error) {
	privateKey, publicKey, err := generateKey(2048)
	if err != nil {
		return "", "", errors.Errorf("Generate key is error: %v", err)
	}
	private := encodePrivateKey(privateKey)
	public, err := encodePublicKey(publicKey)
	if err != nil {
		return "", "", errors.Errorf("Encode Public Key is error: %v", err)
	}

	return string(private), string(public), err
}

// GenerateKey Generate key pair.
func generateKey(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	private, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return private, &private.PublicKey, nil
}

// encodePrivateKey Encode Private Key.
func encodePrivateKey(private *rsa.PrivateKey) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Bytes: x509.MarshalPKCS1PrivateKey(private),
		Type:  "RSA PRIVATE KEY",
	})
}

// encodePublicKey Encode Public Key.
func encodePublicKey(public *rsa.PublicKey) ([]byte, error) {
	publicBytes, err := x509.MarshalPKIXPublicKey(public)
	if err != nil {
		return nil, err
	}
	return pem.EncodeToMemory(&pem.Block{
		Bytes: publicBytes,
		Type:  "PUBLIC KEY",
	}), nil
}
// ***************************************************
type claims struct {
	Uuid   string
	User string
	jwt.StandardClaims
}

func NewClaims(uuid string, user string) *claims {
	expireTime := time.Now().Add(30 * 24 * time.Hour)
	c := &claims{
		Uuid: uuid,
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    viper.GetString("localhost"),
			Subject:   "token",
		},
	}
	return c
}

// GenToken After the user logs in and the password is successfully verified,
// this method will be used to generate a Token and return.
func GenToken(uuid, username string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, NewClaims(uuid, username))
	token, err := t.SignedString([]byte(viper.GetString("token_signed")))
	if err != nil {
		fmt.Println(err)
	}

	return token, nil
}




