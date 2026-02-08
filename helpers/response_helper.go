package helpers

import (
	"encoding/json"
	"kasir-api/models"
	"net/http"
	"strconv"
)

func SendResponse(w http.ResponseWriter, code int, message string, data any) {
	status := "success"
	if code >= 400 {
		status = "error"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(models.BaseResponse{
		Code:    strconv.Itoa(code),
		Status:  status,
		Message: message,
		Data:    data,
	})
}
