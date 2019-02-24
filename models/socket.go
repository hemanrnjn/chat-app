package models

type Todo struct {
	ID          int    `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Done        bool   `json:"done,omitempty"`
}

type Todos []Todo

type ClientRequest struct {
	ID   int    `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	Todo `json:"todo,omitempty"`
}

type ClientResponse struct {
	Todos `json:"todos,omitempty"`
}
