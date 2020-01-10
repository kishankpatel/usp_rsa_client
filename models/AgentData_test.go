package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAgentData(t *testing.T) {
	resData := `{"data":{"AgentID":"50","PublicKey":"testPublicKey","PrivateKey":""}}`
	response := NewAgentData([]byte(resData))
	assert.Equal(t, "50", response.AgentObj.AgentID)
	assert.Equal(t, "testPublicKey", response.AgentObj.PublicKey)

	assert.NotEmpty(t, response)
	fmt.Println("|")
}
