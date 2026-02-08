package models

type BaseResponse struct {
	Code    string `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
