package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prestonbourne/goserve/store"
	"github.com/prestonbourne/goserve/todos"
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

func (s *APIServer) Init() {

	todoController := todos.NewTodoController(s.store)
	router := mux.NewRouter()
	router.HandleFunc("/todos", makeHTTPHandleFunc(todoController.GetAll)).Methods("GET")
	router.HandleFunc("/todos/{id}", makeHTTPHandleFunc(todoController.GetById)).Methods("GET")
	fmt.Println("[Success]: Running on", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func main() {
	//look into environment variables
	const postgresURL string = "postgresql://postgres:password@localhost:5432/testing"
	store, err := store.NewPostgresStore(context.Background(), postgresURL)
	if err != nil {
		log.Fatalf("%v", err)
	}

	server := newAPIServer(":5000", *store)
	server.Init()

	// port := "5000"
	// initStr := fmt.Sprintf("Server running on :%v", port)

	// fmt.Println(initStr)

	// todoHandler := &todos.TodoServer{DB: db}
	// mux := http.NewServeMux()
	// mux.Handle("/todos", todoHandler)

	// err = http.ListenAndServe("localhost:5000", mux)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
