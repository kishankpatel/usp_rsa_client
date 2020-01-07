package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// EncryptMessage sd
func EncryptMessage(id, normalText string) string {
	key, _ := base64.URLEncoding.DecodeString(id)
	message := []byte(normalText)

	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := gcm.Seal(nonce, nonce, message, nil)
	enocdeText := base64.URLEncoding.EncodeToString(ciphertext)
	return enocdeText
}
