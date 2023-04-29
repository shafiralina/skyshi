package models

import (
	"encoding/json"
	"net/http"
)

func Message(status string, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func RespondError(w http.ResponseWriter, data map[string]interface{}, errcode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(errcode)
	json.NewEncoder(w).Encode(data)
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
