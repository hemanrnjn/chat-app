package models

import "github.com/jinzhu/gorm"

type Message struct {
	gorm.Model
	Timestamp string `json:"timeStamp"`
	From      int64  `json:"from"`
	To        int64  `json:"to"`
	Username  string `json:"username"`
	Message   string `json:"message"`
	Is_Read   bool   `json:"is_read" gorm:"default:false`
}
