package models

import "encoding/json"

// AgentData struct Declaration
type AgentData struct {
	AgentObj Agent `json:"data"`
}

// NewAgentData - Parse the response body and create new AgentData
func NewAgentData(resData []byte) *AgentData {
	var agentData AgentData
	json.Unmarshal(resData, &agentData)
	return &agentData
}
