package models

import "encoding/json"

// AgentData struct Declaration
type AgentData struct {
	Data Agent `json:"data"`
}

// NewAgentData method declaration
func NewAgentData(resData []byte) *AgentData {
	var agentData AgentData
	json.Unmarshal(resData, &agentData)
	agent := agentData.Data

	return &AgentData{
		Data: NewAgent(agent.AgentID, agent.Key),
	}
}
