package api

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kishankpatel/usp_client/models"
)

// USPServerAPI - will communicate with the USPServer and return the response
type USPServerAPI interface {
	RegisterAgent(agentID string) ([]byte, error)
	SendMessage(message models.Message) (string, error)
}

type uspServerAPI struct {
	URL string
}

// NewUSPServerAPI - constructor to create new USPServerAPI
func NewUSPServerAPI() (USPServerAPI, error) {
	godotenv.Load()
	url, ok := os.LookupEnv("CRYPTO_SERVER_URL")
	if !ok {
		return nil, errors.New("Env variable CRYPTO_SERVER_URL is not defined")
	}
	return uspServerAPI{
		URL: url,
	}, nil
}

func (server uspServerAPI) RegisterAgent(agentID string) ([]byte, error) {
	response, err := http.Post(server.URL+"/register/"+string(agentID), "application/json", nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return nil, readErr
	}
	fmt.Printf("body:%s", body)
	return body, nil
}

func (server uspServerAPI) SendMessage(message models.Message) (string, error) {
	fmt.Println(message.EncryptedText)
	fmt.Println(message.AgentID())
	url := server.URL + "/send_message/" + string(message.AgentID())
	jsonData := fmt.Sprintf(`{"message":"%s"}`, message.EncryptedText)
	var jsonStr = []byte(jsonData)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, errors := client.Do(req)
	if errors != nil {
		panic(errors)
	}
	defer resp.Body.Close()
	messageBody, _ := ioutil.ReadAll(resp.Body)
	return string(messageBody), nil
}
