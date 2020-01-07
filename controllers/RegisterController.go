package controllers

import (
	"fmt"
	"net/http"

	"../api"

	"../services"
)

// RegisterHandler POST
func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	agentID := r.FormValue("agent_id")
	message := r.FormValue("message")

	agentKey := api.RegisterAgent(agentID)

	encryptedMessage := services.EncryptMessage(agentKey, message)

	messageBody := api.SendMessage(agentID, encryptedMessage)
	fmt.Println("response Body:", messageBody)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
