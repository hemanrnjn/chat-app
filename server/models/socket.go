package models

import (
	"github.com/gorilla/websocket"
)

// ClientConn Client connection obj
type ClientConn struct {
	Conn *websocket.Conn
	ID   uint
}

// ClientRequest object
type ClientRequest struct {
	Timestamp string `json:"timeStamp"`
	FromUser  uint   `json:"from_user"`
	ToUser    uint   `json:"to_user"`
	Username  string `json:"username"`
	Message   string `json:"message"`
	IsRead    bool   `json:"is_read"`
}
