package model

type Response struct {
	Status  string `json:"status"`
	Data    any    `json:"data,omitempty"`
	Version int    `json:"version,omitempty"`
	Message string `json:"message,omitempty"`
}
