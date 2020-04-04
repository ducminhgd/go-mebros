package mebros

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func webhookLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Webhook")
}

func RunServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/webhook", webhookLink)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))
}
