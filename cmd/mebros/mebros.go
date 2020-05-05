// Package mebros defines routes and their handlers for RESTFul API
package mebros

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// webhookLink handles request and response to client for webhook path
func webhookLink(w http.ResponseWriter, r *http.Request) {
	// token := chi.URLParam(r, "token")

	payload, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println(err)
	}

	data := make(map[string][]interface{})
	json.Unmarshal(payload, &data)
	if slacks, found := data["slacks"]; found {
		for _, b := range slacks {
			// webhook := b.(map[string]string)["webhook"]
			// sh := handlers.SlackHandler{Webhook: webhook}
			p := b.(map[string]interface{})["payload"]
			w.Write([]byte(fmt.Sprintf("Payload: %s", p)))
			// sh.SendMessage(payload)
		}
	} else {
		w.Write([]byte(fmt.Sprintln("Failed")))
	}

}

// RunServer runs RESTFul API service of mebros
func RunServer() {
	router := chi.NewRouter()

	// Uses middleware
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)

	router.Post("/webhook/{token}", webhookLink)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))
}
