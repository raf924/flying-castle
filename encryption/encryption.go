package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"io"
)

var aesGCM cipher.AEAD

func DecodeKey(key string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(key)
}

func EncodeKey(key []byte) string {
	return base64.StdEncoding.EncodeToString(key)
}

func MustDecodeKey(key string) []byte {
	r, err := DecodeKey(key)
	if err != nil {
		panic(err)
	}
	return r
}

func Decrypt(packet []byte) []byte {
	nonce := packet[:12]
	data := packet[12:]
	decrypted, err := aesGCM.Open(nil, nonce, data, nil)
	if err != nil {
		panic(err)
	}
	return decrypted
}

func Encrypt(packet []byte) []byte {
	nonce := make([]byte, 12)
	_, err := io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err)
	}
	encrypted := aesGCM.Seal(nil, nonce, packet, nil)
	return append(nonce, encrypted...)
}

func MustUpdateKey(newKey string) {
	var key = MustDecodeKey(newKey)
	block1, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	aesGCM, err = cipher.NewGCM(block1)
	if err != nil {
		panic(err)
	}
}

func MustHash(text string, salt []byte) []byte {
	bytes, err := Hash(text, salt)
	if err != nil {
		panic(err)
	}
	return bytes
}

func Hash(text string, salt []byte) ([]byte, error) {
	hash, err := scrypt.Key(append(salt, []byte(text)...), salt, 32768, 8, 1, 16)
	if err != nil {
		return nil, err
	}
	return append(salt, hash...), nil
}

func GenerateKey() ([]byte, error) {
	var key = make([]byte, 256)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	var salt = make([]byte, 256)
	_, err = rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return scrypt.Key(key, salt, 32768, 8, 1, 32)
}
