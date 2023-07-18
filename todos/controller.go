package todos

import (
	"fmt"
	"net/http"

	"github.com/prestonbourne/goserve/store"
	"github.com/prestonbourne/goserve/utils"
)

type TodoController struct {
	store store.PostgresStore
}

func NewTodoController(store store.PostgresStore) *TodoController {
	return &TodoController{
		store: store,
	}
}

func logRequest(r *http.Request) {
	log := fmt.Sprintf("[Request Type]: %v\n[Request Body]: %v\n[Request Path]: %v", r.Method, r.Body, r.URL.Path)
	fmt.Println(log)
}

func (c *TodoController) GetAll(w http.ResponseWriter, r *http.Request) error {
	//id := mux.Vars(r)["id"]
	todo := NewTodo("Nothing here yet...")
	return utils.WriteJSON(w, http.StatusAccepted, todo)
}
func (c *TodoController) GetById(w http.ResponseWriter, r *http.Request) error {
	//id := mux.Vars(r)["id"]
	todo := NewTodo("Nothing here yet...")
	return utils.WriteJSON(w, http.StatusAccepted, todo)
}

func (c *TodoController) Add(w http.ResponseWriter, r *http.Request) error {

	todo := NewTodo("Nothing here yet...")
	return utils.WriteJSON(w, http.StatusAccepted, todo)
}
func (c *TodoController) Delete(w http.ResponseWriter, r *http.Request) error {

	todo := NewTodo("Nothing here yet...")
	return utils.WriteJSON(w, http.StatusAccepted, todo)
}
func (c *TodoController) Update(w http.ResponseWriter, r *http.Request) error {

	todo := NewTodo("Nothing here yet...")
	return utils.WriteJSON(w, http.StatusAccepted, todo)
}

// func Controller(w http.ResponseWriter, r *http.Request) error {

// 	switch r.Method {
// 	case http.MethodPost:
// 		// POST /players

// 		var jsonRequest Todo

// 		dec := json.NewDecoder(r.Body)
// 		dec.Decode(&jsonRequest)

// 	case http.MethodGet:
// 		id := mux.Vars(r)["id"]
// 		hasId := id == ""
// 		if hasId {
// 			response := handleGetById(w, id)
// 			return response
// 		}
// 		response := handleGet(w)
// 		return response

// 	case http.MethodPatch, http.MethodPut:

// 	}
// 	return fmt.Errorf("")
// }

// func handleGet(w http.ResponseWriter) error {
// 	todo := NewTodo("Nothing here yet...")
// 	return utils.WriteJSON(w, http.StatusAccepted, todo)
// }

// func handleGetById(w http.ResponseWriter, id string) error {
// 	todo := NewTodo("Nothing here yet...")
// 	return utils.WriteJSON(w, http.StatusAccepted, todo)
// }
