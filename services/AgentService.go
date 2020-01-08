package services

import (
	"fmt"

	"../models"
)

// AgentDataService interface declaration
type AgentDataService interface {
	parseAgentData()
}

type agentDataService struct {
	ResData []byte
}

// NewAgentDataService method declaration
func NewAgentDataService(resData []byte) {
	var ads AgentDataService = agentDataService{resData}
	ads.parseAgentData()
}

func (ds agentDataService) parseAgentData() {

	a := models.NewAgentData([]byte(ds.ResData))
	fmt.Printf("===========> %+v", a)
}
