package utils

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, message string) {
	response := map[string]string{"message": message}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
