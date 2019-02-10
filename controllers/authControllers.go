package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/hemanrnjn/chat-app/models"
	u "github.com/hemanrnjn/chat-app/utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create() //Create account
	u.Respond(w, resp)
}

var GetAllUsers = func(w http.ResponseWriter, r *http.Request) {

	resp := models.GetAllUsers() //Get All Users
	// u.Respond(w, resp)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}
