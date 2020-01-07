package main

import (
	"net/http"

	routes "./routes"
	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func main() {
	routes.Handler()
	http.ListenAndServe(":8080", nil)
}
