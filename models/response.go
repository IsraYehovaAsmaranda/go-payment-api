package models

type CommonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorData struct {
	Error string `json:"error"`
}
