package models

import (
	u "github.com/hemanrnjn/chat-app/utils"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// Message struct of message
type Message struct {
	gorm.Model
	ID        uint   `json:"id"`
	Timestamp string `json:"timeStamp"`
	FromUser  uint   `json:"from_user"`
	ToUser    uint   `json:"to_user"`
	Username  string `json:"username"`
	Message   string `json:"message"`
	IsRead    bool   `json:"is_read"`
}

// Messages slice of message
type Messages []Message

// AddMessage adds message for user
func (message *Message) AddMessage() map[string]interface{} {
	log.Info("Message: ", message)

	GetDB().Create(message)

	if message.ID <= 0 {
		return u.Message(false, "Failed to save Message, connection error.")
	}

	response := u.Message(true, "Message has been saved")
	response["message"] = message
	return response
}

// GetMessagesForUser gets message
func GetMessagesForUser(userID uint) map[string]interface{} {

	var messages Messages

	GetDB().Where("to_user = ? OR from_user = ?", userID, userID).Find(&messages)

	response := u.Message(true, "Message retrived")
	response["messages"] = messages
	return response
}
