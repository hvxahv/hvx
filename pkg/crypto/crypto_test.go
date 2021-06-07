package crypto

import (
	"testing"
)

func TestGenerateKey(t *testing.T) {
	privateKey, publicKey, err := GenerateKey(2048)
	if err != nil {
		return
	}

	private := EncodePrivateKey(privateKey)

	public, err := EncodePublicKey(publicKey)
	if err != nil {
		t.Log("Encode Public Key is error: ", err)
	}

	t.Log(string(private), string(public))
}