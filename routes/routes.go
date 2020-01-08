package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kishankpatel/usp_client/api"
	"github.com/kishankpatel/usp_client/models"
	"github.com/kishankpatel/usp_client/utils"
)

var router = mux.NewRouter()

// Handler method declaration
func Handler() {

	uspServerAPI, err := api.NewUSPServerAPI()
	if err != nil {
		panic(err)
	}

	router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		var indexBody, err = utils.LoadFile("templates/index.html")
		if err != nil {
			fmt.Println(err.Error())
			fmt.Fprintf(response, "Some error happened in the server")
		}
		fmt.Fprintf(response, indexBody)
	})

	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		agentID := r.FormValue("agent_id")
		plainText := r.FormValue("message")

		response, err := uspServerAPI.RegisterAgent(agentID)
		if err != nil {
			panic(err)
		}

		agentResponse := models.NewAgentData(response)
		message := models.NewMessage(agentResponse.AgentObj, plainText)

		err = message.Encrypt()
		if err != nil {
			panic(err)
		}

		responseStr, err := uspServerAPI.SendMessage(message)
		if err != nil {
			panic(err)
		}

		fmt.Println("Response from server for the message:", responseStr)

		http.Redirect(w, r, r.Header.Get("Referer"), 302)
	}).Methods("POST")

	http.Handle("/", router)
}
