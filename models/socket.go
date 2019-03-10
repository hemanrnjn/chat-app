package models

import (
	"github.com/gorilla/websocket"
)

type ClientConn struct {
	Conn  *websocket.Conn
	Email string
}

type Message struct {
	Timestamp string `json:"timeStamp"`
	From      string `json:"from"`
	To        string `json:"to"`
	Username  string `json:"username"`
	Message   string `json:"message"`
}

type Messages []Message

type ClientRequest struct {
	Timestamp string `json:"timeStamp"`
	From      string `json:"from"`
	To        string `json:"to"`
	Username  string `json:"username"`
	Message   string `json:"message"`
}

// type ClientResponse struct {
// 	Messages `json:"messages,omitempty"`
// }
