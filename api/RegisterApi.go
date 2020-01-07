package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"../models"
	"github.com/joho/godotenv"
)

// RegisterAgent method declaaration
func RegisterAgent(agentID string) string {
	godotenv.Load()
	url, _ := os.LookupEnv("CRYPTO_SERVER_URL")
	response, err := http.Post(url+"/register/"+string(agentID), "application/json", nil)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		panic(readErr)
	}
	println("body:", body)

	var agentData models.AgentData
	readErr = json.Unmarshal(body, &agentData)

	if readErr != nil {
		panic(readErr)
	}
	agent := agentData.Data
	return agent.Key
}
