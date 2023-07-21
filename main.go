package main

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prestonbourne/goserve/store"
	"github.com/prestonbourne/goserve/todos"
	"github.com/prestonbourne/goserve/users"
	"github.com/prestonbourne/goserve/utils"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}

	}
}

func newAPIServer(listenAddr string, store store.PostgresStore) *APIServer {
	return &APIServer{
		listenAddr,
		store,
	}
}

type APIServer struct {
	listenAddr string
	store      store.PostgresStore
}

type ApiError struct {
	Error string `json:"error"`
}

func (s *APIServer) serve() {

	todoController := todos.NewTodoController(s.store)
	userController := users.NewUserController(s.store)

	router := mux.NewRouter()
	//gotta be a dryer way to do this 😅
	router.HandleFunc("/todos", makeHTTPHandleFunc(todoController.GetAll)).Methods("GET")
	router.HandleFunc("/todos/{id}", makeHTTPHandleFunc(todoController.GetById)).Methods("GET")

	router.HandleFunc("/login", makeHTTPHandleFunc(userController.Login)).Methods("POST")
	router.HandleFunc("/users", makeHTTPHandleFunc(userController.Add)).Methods("POST")
	router.HandleFunc("/users", makeHTTPHandleFunc(userController.GetAll)).Methods("GET")
	router.HandleFunc("/users/{id}", makeHTTPHandleFunc(userController.GetById)).Methods("GET")
	router.HandleFunc("/users/{id}", makeHTTPHandleFunc(userController.Delete)).Methods("DELETE")

	utils.Success("Running on port " + s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)

}

func main() {
	//look into environment variables
	const postgresURL string = "postgresql://postgres:password@localhost:5432/testing"
	store, err := store.NewPostgresStore(context.Background(), postgresURL)
	if err != nil {
		utils.Throw("Could not connect to Postgres", err)
	}

	server := newAPIServer(":5000", *store)
	server.serve()

}
