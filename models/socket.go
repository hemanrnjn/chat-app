package models

import (
	"github.com/gorilla/websocket"
)

type ClientConn struct {
	Conn *websocket.Conn
	Id   uint
}

type ClientRequest struct {
	Timestamp string `json:"timeStamp"`
	From_User uint   `json:"from_user"`
	To_User   uint   `json:"to_user"`
	Username  string `json:"username"`
	Message   string `json:"message"`
	Is_Read   bool   `json:"is_read"`
}
