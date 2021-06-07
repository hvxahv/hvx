package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"golang.org/x/crypto/ssh"
)
// GenerateKey ...
func GenerateKey(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	private, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return private, &private.PublicKey, nil

}

func EncodePrivateKey(private *rsa.PrivateKey) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Bytes: x509.MarshalPKCS1PrivateKey(private),
		Type:  "RSA PRIVATE KEY",
	})
}

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

func EncodeSSHKey(public *rsa.PublicKey) ([]byte, error) {
	publicKey, err := ssh.NewPublicKey(public)
	if err != nil {
		return nil, err
	}
	return ssh.MarshalAuthorizedKey(publicKey), nil
}

func MakeSSHKeyPair() (string, string, error) {
	pkey, pubkey, err := GenerateKey(2048)
	if err != nil {
		return "", "", err
	}

	pub, err := EncodeSSHKey(pubkey)
	if err != nil {
		return "", "", err
	}

	return string(EncodePrivateKey(pkey)), string(pub), nil
}

// GenRasKey 该方法调用 utils.GenerateKey 包生成 2068 位的 ras key 返回公钥和私钥
func GenRasKey() (string, string, error) {
	privateKey, publicKey, err := GenerateKey(2048)
	if err != nil {
		fmt.Printf("Generate key is error: %s", err)
	}

	private := EncodePrivateKey(privateKey)

	public, err := EncodePublicKey(publicKey)
	if err != nil {
		fmt.Println("Encode Public Key is error: ", err)
	}

	return string(private), string(public), err
}

