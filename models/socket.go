package models

import (
	"github.com/gorilla/websocket"
)

// type Todo struct {
// 	ID          int    `json:"id,omitempty"`
// 	Description string `json:"description,omitempty"`
// 	Done        bool   `json:"done,omitempty"`
// }

// type Todos []Todo

// type ClientRequest struct {
// 	ID   int    `json:"id,omitempty"`
// 	Type string `json:"type,omitempty"`
// 	Todo `json:"todo,omitempty"`
// }

// type ClientResponse struct {
// 	Todos `json:"todos,omitempty"`
// }

type ClientConn struct {
	Conn  *websocket.Conn
	Email string
}

type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

type Messages []Message

type ClientRequest struct {
	// ID      int `json:"id,omitempty"`
	// Message `json:"message,omitempty"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

type ClientResponse struct {
	Messages `json:"messages,omitempty"`
}
