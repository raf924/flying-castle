package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"testing"
)

func TestCipher(t *testing.T) {
	var key = make([]byte, 32)
	var nonce = make([]byte, 12)
	_, _ = rand.Read(key)
	_, _ = rand.Read(nonce)
	block1, _ := aes.NewCipher(key)
	aesGCM1, _ := cipher.NewGCM(block1)
	block2, _ := aes.NewCipher(key)
	aesGCM2, _ := cipher.NewGCM(block2)
	var encry = aesGCM1.Seal(nil, nonce, []byte("hello"), nil)
	decry, _ := aesGCM2.Open(nil, nonce, encry, nil)
	if !bytes.Equal(decry, []byte("hello")) {
		t.Fatal("Very different")
	}
}
