package models

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// Message struct Declaration
type Message struct {
	EncryptedText string
	PlainText     string
	Agent         Agent
}

// NewMessage - constructor to create a new message
func NewMessage(agent Agent, plainText string) Message {
	return Message{
		PlainText: plainText,
		Agent:     agent,
	}
}

// Encrypt - encrypts the message and returns the encrypted text
func (message *Message) Encrypt() error {
	block, _ := aes.NewCipher(message.Key())
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(message.PlainText), nil)
	encodedText := base64.URLEncoding.EncodeToString(ciphertext)
	message.EncryptedText = encodedText
	fmt.Printf("encrypted text %s", encodedText)
	return nil
}

// AgentID - getter used to fetch the Agent ID info
func (message Message) AgentID() string {
	return message.Agent.AgentID
}

// Key - getter used to fetch the Key info
func (message Message) Key() []byte {
	fmt.Println("Key...", message.Agent.Key)
	key, err := base64.URLEncoding.DecodeString(message.Agent.Key)
	if err != nil {
		return []byte{}
	}
	return key
}
