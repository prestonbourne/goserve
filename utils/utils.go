package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

const reset string = "\u001b[0m"

func LogError(text string, err error) {
	const red string = "\u001b[31m"
	output := fmt.Sprintf(red + "[Failure]: " + reset + text + "\n")
	fmt.Println(output+"%v", err)
}

func Throw(text string, err error) {
	const red string = "\u001b[31m"
	output := fmt.Sprintf(red + "[Error]: " + reset + text + "\n%v")
	log.Fatalln(output, err)
}

func Success(text string) {
	const green string = "\u001b[32m"
	output := fmt.Sprintf(green + "[Success]: " + reset + text)
	fmt.Println(output)
}

func GetId(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("Invalid ID, expected integer. Received %v", idStr)
	}
	return id, nil
}
