// Package mebros defines routes and their handlers for RESTFul API
package mebros

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// webhookLink handle request and response to client for webhook path
func webhookLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Webhook")
}

// RunServer runs RESTFul API service of mebros
func RunServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/webhook", webhookLink)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))
}
