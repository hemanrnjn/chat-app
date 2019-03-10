package controllers

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/hemanrnjn/chat-app/models"
	log "github.com/sirupsen/logrus"
)

var upgrader websocket.Upgrader
var msg models.Message

var clients []models.ClientConn
var broadcast = make(chan *models.ClientRequest)

func HandleSocketMessages(w http.ResponseWriter, r *http.Request) {

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
		if contains(clients, creq.From) == false {
			n := models.ClientConn{Conn: conn, Email: creq.From}
			clients = append(clients, n)
		}
		// } else {
		// 	clients = replace(clients, creq, conn)
		// }
		log.Infof("Clients Connections: %v", clients)
		if err != nil {
			log.Println(err)
			break
		}
		log.Infof("Message from client : %v", creq)
		if creq.To != "" {
			broadcast <- creq
		}

	}
}

func HandleMessages() {
	for {

		msg := <-broadcast

		for _, client := range clients {
			if msg.To == client.Email {
				log.Info(client, msg)
				err := client.Conn.WriteJSON(msg)
				if err != nil {
					log.Printf("error: %v", err)
					client.Conn.Close()
					clients = delete(clients, client.Email)
				}
				break
			}
		}
	}
}

func contains(list []models.ClientConn, email string) bool {
	if len(list) > 1 {
		for _, ele := range list {
			if ele.Email == email {
				log.Info("True")
				return true
			}
		}
		log.Info("False")
		return false
	}
	return false
}

func delete(list []models.ClientConn, email string) []models.ClientConn {
	for i, ele := range list {
		if ele.Email == email {
			list[i] = list[len(list)-1]
			list = list[:len(list)-1]
			break
		}
	}
	return list
}

// func replace(list []models.ClientConn, creq *models.ClientRequest, conn *websocket.Conn) []models.ClientConn {
// 	for i, ele := range list {
// 		if ele.Email == creq.From {
// 			list[i].Conn = conn
// 			break
// 		}
// 	}
// 	return list
// }
