// Package mebros defines routes and their handlers for RESTFul API
package mebros

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// webhookLink handle request and response to client for webhook path
func webhookLink(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	w.Write([]byte(fmt.Sprintf("title: %s", token)))
}

// RunServer runs RESTFul API service of mebros
func RunServer() {
	router := chi.NewRouter()

	// Uses middleware
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)

	router.Get("/webhook/{token}", webhookLink)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))
}
