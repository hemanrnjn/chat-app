package controllers

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/hemanrnjn/chat-app/models"
	log "github.com/sirupsen/logrus"
)

var upgrader websocket.Upgrader

var clients []models.ClientConn
var broadcast = make(chan *models.ClientRequest)
var respWriter http.ResponseWriter

// HandleSocketMessages handles socket messages
func HandleSocketMessages(w http.ResponseWriter, r *http.Request) {

	respWriter = w

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}

	defer conn.Close()

	for {
		creq := &models.ClientRequest{}
		err := conn.ReadJSON(creq)

		if contains(clients, creq.FromUser) == false {
			n := models.ClientConn{Conn: conn, ID: creq.FromUser}
			clients = append(clients, n)
		}

		log.Infof("Clients Connections: %v", clients)
		if err != nil {
			log.Println(err)
			break
		}
		log.Infof("Message from client : %v", creq)
		if creq.ToUser != 0 {
			broadcast <- creq
		}

	}
}

// HandleMessages handles each message
func HandleMessages() {
	for {

		message := <-broadcast

		msg := models.Message{}

		msg.Timestamp = message.Timestamp
		msg.FromUser = message.FromUser
		msg.ToUser = message.ToUser
		msg.Username = message.Username
		msg.Message = message.Message
		msg.IsRead = message.IsRead

		log.Infof("Message Model: %v", msg)

		resp := msg.AddMessage() //Add Message
		log.Info("Message Status: ", resp["status"].(bool))

		for _, client := range clients {
			if message.ToUser == client.ID {
				log.Info(client, message)
				err := client.Conn.WriteJSON(message)
				if err != nil {
					log.Printf("error: %v", err)
          _ = client.Conn.Close()
					clients = deleteConn(clients, client.ID)
				}
				break
			}
		}
	}
}

func contains(list []models.ClientConn, id uint) bool {
	for _, ele := range list {
		if ele.ID == id || id == 0 {
			return true
		}
	}

	return false
}

func deleteConn(list []models.ClientConn, id uint) []models.ClientConn {
	for i, ele := range list {
		if ele.ID == id {
			list[i] = list[len(list)-1]
			list = list[:len(list)-1]
			break
		}
	}
	return list
}

// func replace(list []models.ClientConn, creq *models.ClientRequest, conn *websocket.Conn) []models.ClientConn {
// 	for i, ele := range list {
// 		if ele.Email == creq.From_User {
// 			list[i].Conn = conn
// 			break
// 		}
// 	}
// 	return list
// }
