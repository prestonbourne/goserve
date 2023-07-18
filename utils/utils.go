package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(v)
}

func LogRequest(r *http.Request) {
	log := fmt.Sprintf("[Request Type]: %v\n[Request Body]: %v\n[Request Path]: %v", r.Method, r.Body, r.URL.Path)
	fmt.Println(log)
}

func DecodeAndWrite(r *http.Request, val any) error {
	if err := json.NewDecoder(r.Body).Decode(val); err != nil {
		return err
	}
	return nil
}
