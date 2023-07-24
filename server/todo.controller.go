package server

import (
	"fmt"
	"net/http"

	"github.com/prestonbourne/goserve/models"
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
	todo := models.NewTodo("Nothing here yet...")
	return utils.WriteJSON(w, http.StatusAccepted, todo)
}
func (c *TodoController) GetById(w http.ResponseWriter, r *http.Request) error {
	//id := mux.Vars(r)["id"]
	todo := models.NewTodo("Nothing here yet...")
	return utils.WriteJSON(w, http.StatusAccepted, todo)
}

func (c *TodoController) Add(w http.ResponseWriter, r *http.Request) error {

	todo := models.NewTodo("Nothing here yet...")
	return utils.WriteJSON(w, http.StatusAccepted, todo)
}
func (c *TodoController) Delete(w http.ResponseWriter, r *http.Request) error {

	todo := models.NewTodo("Nothing here yet...")
	return utils.WriteJSON(w, http.StatusAccepted, todo)
}
func (c *TodoController) Update(w http.ResponseWriter, r *http.Request) error {

	todo := models.NewTodo("Nothing here yet...")
	return utils.WriteJSON(w, http.StatusAccepted, todo)
}
