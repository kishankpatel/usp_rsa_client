package models

// Agent struct Declaration
type Agent struct {
	AgentID string `json:"AgentID"`
	Key     string `json:"Key"`
}

// NewAgent method declaration
func NewAgent(agentID, Key string) Agent {
	return Agent{
		AgentID: agentID,
		Key:     Key,
	}
}
