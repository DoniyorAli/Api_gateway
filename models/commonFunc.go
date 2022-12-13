package models


type JSONRespons struct {
    Message string       `json:"message"`
    Data    interface{}  `json:"data"`
}

type JSONErrorRespons struct {
	Message string `json:"message"`
	Error string `json:"error"`
}