package models

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"

	"github.com/kishankpatel/usp_client/utils"
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
	publicKey, _ := utils.StringToPublicKey(message.PublicKey())
	label := []byte("")
	hash := sha256.New()
	byteMessage := []byte(message.PlainText)
	ciphertext, _ := rsa.EncryptOAEP(
		hash,
		rand.Reader,
		publicKey,
		byteMessage,
		label,
	)
	encodeCiphertext := base64.URLEncoding.EncodeToString(ciphertext)
	message.EncryptedText = encodeCiphertext
	return nil
}

// AgentID - getter used to fetch the Agent ID info
func (message Message) AgentID() string {
	return message.Agent.AgentID
}

// PublicKey - getter used to fetch the PublicKey info
func (message Message) PublicKey() string {
	return message.Agent.PublicKey
}
