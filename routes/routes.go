package routes

import (
	"net/http"

	"../controllers"
	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

// Handler method declaration
func Handler() {
	router.HandleFunc("/", controllers.IndexPageHandler) // GET

	router.HandleFunc("/register", controllers.RegisterHandler).Methods("POST")

	http.Handle("/", router)
}
