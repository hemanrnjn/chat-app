package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/hemanrnjn/chat-app/models"
	u "github.com/hemanrnjn/chat-app/utils"
	log "github.com/sirupsen/logrus"
)

func GetUserMessages(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.GetMessagesForUser(account.ID)
	log.Info("Retrived Messages: ", resp)
	u.Respond(w, resp)
}
