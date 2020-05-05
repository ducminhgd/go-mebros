// Package handlers handles request to other services
package handlers

import (
	"log"
	"net/http"
	"strings"
)

type SlackHandler struct {
	Webhook string
}

func (h *SlackHandler) SendMessage(p string) bool {
	req, err := http.NewRequest("POST", h.Webhook, strings.NewReader(p))
	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer rsp.Body.Close()

	if rsp.StatusCode == 200 {
		return true
	}
	return false
}
