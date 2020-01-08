package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kishankpatel/usp_client/routes"
)

var router = mux.NewRouter()

func main() {
	routes.Handler()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
