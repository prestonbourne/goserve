package main

import (
	"context"
	"fmt"
	"log"
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
	Error string
}

func (s *APIServer) run() {

	todoController := todos.NewTodoController(s.store)
	userController := users.NewUserController(s.store)

	router := mux.NewRouter()
	router.HandleFunc("/todos", makeHTTPHandleFunc(todoController.GetAll)).Methods("GET")
	router.HandleFunc("/todos/{id}", makeHTTPHandleFunc(todoController.GetById)).Methods("GET")
	router.HandleFunc("/users", makeHTTPHandleFunc(userController.Add)).Methods("POST")

	http.ListenAndServe(s.listenAddr, router)
	fmt.Println("[Success]: Running on", s.listenAddr)
}

func main() {
	//look into environment variables
	const postgresURL string = "postgresql://postgres:password@localhost:5432/testing"
	store, err := store.NewPostgresStore(context.Background(), postgresURL)
	if err != nil {
		log.Fatalf("%v", err)
	}

	server := newAPIServer(":5000", *store)
	server.run()

}
