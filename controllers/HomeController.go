package controllers

import (
	"fmt"
	"net/http"

	"../services"
)

// IndexPageHandler GET
func IndexPageHandler(response http.ResponseWriter, request *http.Request) {
	var indexBody, _ = services.LoadFile("templates/index.html")
	fmt.Fprintf(response, indexBody)
}
