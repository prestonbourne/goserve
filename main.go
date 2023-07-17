package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/prestonbourne/goserve/db"
	"github.com/prestonbourne/goserve/todos"
)

func main() {
	const postgres_url string = "postgresql://postgres:password@localhost:5432/postgres"
	db, err := db.Init(context.Background(), postgres_url)
	if err != nil {
		fmt.Printf("%v", err)
	}

	port := 5000
	initStr := fmt.Sprintf("Server running on :%v", port)

	fmt.Println(initStr)

	todoHandler := &todos.TodoServer{DB: db}
	mux := http.NewServeMux()
	mux.Handle("/todos", todoHandler)

	err = http.ListenAndServe(initStr, mux)
	if err != nil {
		log.Fatal(err)
	}

}
