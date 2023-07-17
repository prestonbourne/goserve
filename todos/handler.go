package todos

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (server *TodoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	confirmation := fmt.Sprintf("[request]: %v", r.Method)
	fmt.Println(confirmation)
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		// POST /players

		var jsonRequest TodoModel

		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&jsonRequest)

	case http.MethodGet:
		// GET /players/<name/id>

	case http.MethodPatch, http.MethodPut:

	}
}
