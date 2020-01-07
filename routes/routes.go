package handlers

import (
	"net/http"

	common "../common"
	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

// Handler method declaration
func Handler() {
	router.HandleFunc("/", common.IndexPageHandler) // GET

	router.HandleFunc("/register", common.RegisterHandler).Methods("POST")

	http.Handle("/", router)
}
