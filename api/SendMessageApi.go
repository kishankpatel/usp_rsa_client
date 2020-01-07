package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SendMessage method declaaration
func SendMessage(agentID, encryptedMessage string) string {
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
	return string(messageBody)
}
