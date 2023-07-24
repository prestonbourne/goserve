package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prestonbourne/goserve/store"
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

func NewAPIServer(listenAddr string, store store.PostgresStore) *APIServer {
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

func (s *APIServer) ListenAndServe() {

	todoController := NewTodoController(s.store)
	userController := NewUserController(s.store)

	router := mux.NewRouter()

	router.HandleFunc("/todos", makeHTTPHandleFunc(todoController.GetAll)).Methods("GET")
	router.HandleFunc("/todos/{id}", makeHTTPHandleFunc(todoController.GetById)).Methods("GET")

	router.HandleFunc("/login", makeHTTPHandleFunc(userController.Login)).Methods("POST")
	router.HandleFunc("/users", makeHTTPHandleFunc(userController.Add)).Methods("POST")
	router.HandleFunc("/users", makeHTTPHandleFunc(userController.GetAll)).Methods("GET")
	router.HandleFunc("/users/{id}", makeHTTPHandleFunc(userController.GetById)).Methods("GET")
	router.HandleFunc("/users/{id}", makeHTTPHandleFunc(userController.Delete)).Methods("DELETE")

	utils.Success("Server running on " + s.listenAddr)

	err := http.ListenAndServe(s.listenAddr, router)

	if err != nil {
		utils.Throw("Could not start server", err)
	}

}
