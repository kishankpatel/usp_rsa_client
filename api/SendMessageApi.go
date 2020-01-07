package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// SendMessage method declaaration
func SendMessage(agentID, encryptedMessage string) string {
	godotenv.Load()
	url, _ := os.LookupEnv("CRYPTO_SERVER_URL")
	url = url + "/send_message/" + string(agentID)
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
