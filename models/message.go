package models

import (
	u "github.com/hemanrnjn/chat-app/utils"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type Message struct {
	gorm.Model
	Timestamp string `json:"timeStamp"`
	From_User uint   `json:"from_user"`
	To_User   uint   `json:"to_user"`
	Username  string `json:"username"`
	Message   string `json:"message"`
	Is_Read   bool   `json:"is_read"`
}

type Messages []Message

func (message *Message) AddMessage() map[string]interface{} {
	log.Info("Message: ", message)

	if err := GetDB().Create(message); err != nil {
		return u.Message(false, "Failed to save Message, connection error.")
	}

	response := u.Message(true, "Message has been saved")
	response["message"] = message
	return response
}

func GetMessagesForUser(userId uint) map[string]interface{} {

	var messages Messages

	if err := GetDB().Where("to_user = ?", userId).Find(&messages); err != nil {
		response := u.Message(false, "Message not retrived!")
		return response
	}

	response := u.Message(true, "Message retrived")
	response["messages"] = messages
	return response
}
