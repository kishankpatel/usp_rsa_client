package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../models"

	services "../services"
)

// Handlers

// RegisterHandler POST
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	agentID := r.FormValue("agent_id")
	message := r.FormValue("message")
	response, err := http.Post("http://localhost:4040/register/"+string(agentID), "application/json", nil)
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
	println("Agent key: ", agent.Key)
	encryptedMessage := services.EncryptMessage(agent.Key, message)
	println("encryptedMessage: ", encryptedMessage)
	url := "http://localhost:4040/send_message/" + string(agentID)
	jsonData := fmt.Sprintf(`{"message":"%s"}`, encryptedMessage)
	var jsonStr = []byte(jsonData)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, errors := client.Do(req)
	if errors != nil {
		panic(errors)
	}
	defer resp.Body.Close()
	messageBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(messageBody))
	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

// IndexPageHandler GET
func IndexPageHandler(response http.ResponseWriter, request *http.Request) {
	var indexBody, _ = services.LoadFile("templates/index.html")
	fmt.Fprintf(response, indexBody)
}
