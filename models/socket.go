package models

import (
	"github.com/gorilla/websocket"
	u "github.com/hemanrnjn/chat-app/utils"
	"github.com/jinzhu/gorm"
)

type ClientConn struct {
	Conn *websocket.Conn
	Id   int64
}

type Messages []Message

type ClientRequest struct {
	gorm.Model
	Timestamp string `json:"timeStamp"`
	From      int64  `json:"from"`
	To        int64  `json:"to"`
	Username  string `json:"username"`
	Message   string `json:"message"`
	Is_Read   bool   `json:"is_read"`
}

func (message *Message) AddMessage() map[string]interface{} {

	GetDB().Create(message)

	if message.ID <= 0 {
		return u.Message(false, "Failed to add Message, connection error.")
	}

	response := u.Message(true, "Message has been created")
	response["message"] = message
	return response
}
