package controllers

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/hemanrnjn/chat-app/models"
	log "github.com/sirupsen/logrus"
)

var upgrader websocket.Upgrader

func HandleSocketMessages(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}

	for {
		creq := &models.ClientRequest{}
		err := conn.ReadJSON(creq)
		if err != nil {
			log.Println(err)
			return
		}
		log.Infof("Message from client : %v", creq)
		// cresp := &models.ClientResponse{Todos: todos}
		conn.WriteJSON("{'todos': []}")
	}
}
