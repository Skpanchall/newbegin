package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string      `json:"message"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func SendSuccessResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	res := Response{
		Status: "Success",
		Data:   data,
	}
	json.NewEncoder(w).Encode(res)
}
func SendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	res := Response{
		Status: "Error",
		Error:  message,
	}
	json.NewEncoder(w).Encode(res)
}
