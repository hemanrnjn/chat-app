package controllers

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/hemanrnjn/chat-app/models"
	log "github.com/sirupsen/logrus"
)

var upgrader websocket.Upgrader
var msg models.Message

// var clients = make(map[*websocket.Conn]bool)
var clients []models.ClientConn
var broadcast = make(chan *models.ClientRequest)

// type message struct {
// 	Message string `json:message`
// }

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
	// clients[conn] = true
	n := models.ClientConn{Conn: conn, Email: "abc@abc.com"}
	clients = append(clients, n)
	log.Infof("%v", clients)

	for {
		creq := &models.ClientRequest{}
		err := conn.ReadJSON(creq)
		if err != nil {
			log.Println(err)
			clients = clients[:len(clients)-1]
			// delete(clients, conn)
			break
		}
		log.Infof("Message from client : %v", creq)

		broadcast <- creq

		// cresp := &models.ClientResponse{Todos: todos}
		// str := "Hello World Back"
		// msg.Message = str
		// json.Unmarshal([]byte(str), &msg)
		// conn.WriteJSON(msg)
	}
}

func HandleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.Conn.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Conn.Close()
				// delete(clients, client)

			}
		}
	}
}
