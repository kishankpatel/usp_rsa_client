package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMessage(t *testing.T) {
	resData := `{"data":{"AgentID":"50","PublicKey":"testPublicKey","PrivateKey":""}}`
	agentResponse := NewAgentData([]byte(resData))
	plainText := "Hello World"
	messageResponse := NewMessage(agentResponse.AgentObj, plainText)

	assert.Equal(t, "50", messageResponse.Agent.AgentID)
	assert.Equal(t, "testPublicKey", messageResponse.Agent.PublicKey)
	assert.Equal(t, "Hello World", messageResponse.PlainText)
	assert.Equal(t, "", messageResponse.EncryptedText)

	assert.NotEmpty(t, messageResponse)
	fmt.Println("|")
}

func TestEncrypt(t *testing.T) {
	resData := `{"data":{"AgentID":"50","PublicKey":"-----BEGIN RSA public Key-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3xpF3mpCSgQ+HAyf9qBo\nuGX5V0Bdb5B0ZzfqcUQxUsDI66z4S7x5BwbxMOwBC/ZoKxtuEN2vkqUeEJHzV8Y2\nnz6xhSGJQjKfjK1qNh+md9Xrho1Ucl/dJrzM1woD2hv9RtWx8ioLp2gmFDqaSTPH\nZZW7DJ+mNXjpV/GmqDa0q48Xq2+3z1oaYi8k6te+puQ9w0jKFduihdq9UCG0phcf\nUGLbOYFeJGTI48cUjd5AW2dE/nu1FpZOi+e+nTpFK5UqJneAqpZu8AXtX3dNrQyn\nYvCAnpX9GLl3qjWvcsy9Lo09qfSjCRxfBNYz28kzzUDe0517vR/yT+pMRNs8yddT\nmwIDAQAB\n-----END RSA public Key-----\n","PrivateKey":""}}`
	agentResponse := NewAgentData([]byte(resData))
	plainText := "Hello World"
	message := NewMessage(agentResponse.AgentObj, plainText)

	message.Encrypt()
	assert.NotEmpty(t, message.EncryptedText)

}

func TestAgentID(t *testing.T) {
	resData := `{"data":{"AgentID":"50","PublicKey":"-----BEGIN RSA public Key-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3xpF3mpCSgQ+HAyf9qBo\nuGX5V0Bdb5B0ZzfqcUQxUsDI66z4S7x5BwbxMOwBC/ZoKxtuEN2vkqUeEJHzV8Y2\nnz6xhSGJQjKfjK1qNh+md9Xrho1Ucl/dJrzM1woD2hv9RtWx8ioLp2gmFDqaSTPH\nZZW7DJ+mNXjpV/GmqDa0q48Xq2+3z1oaYi8k6te+puQ9w0jKFduihdq9UCG0phcf\nUGLbOYFeJGTI48cUjd5AW2dE/nu1FpZOi+e+nTpFK5UqJneAqpZu8AXtX3dNrQyn\nYvCAnpX9GLl3qjWvcsy9Lo09qfSjCRxfBNYz28kzzUDe0517vR/yT+pMRNs8yddT\nmwIDAQAB\n-----END RSA public Key-----\n","PrivateKey":""}}`
	agentResponse := NewAgentData([]byte(resData))
	plainText := "Hello World"
	message := NewMessage(agentResponse.AgentObj, plainText)

	message.Encrypt()
	assert.NotEmpty(t, message.Agent.AgentID)
	assert.NotEmpty(t, "50", message.Agent.AgentID)

}
