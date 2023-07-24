package main

import (
	"context"

	"github.com/prestonbourne/goserve/server"
	"github.com/prestonbourne/goserve/store"
	"github.com/prestonbourne/goserve/utils"
)

func main() {

	//look into environment variables
	const postgresURL string = "postgresql://postgres:password@localhost:5432/testing"
	store, err := store.NewPostgresStore(context.Background(), postgresURL)
	if err != nil {
		utils.Throw("Could not connect to Postgres", err)
	}
	const listenAddr string = "localhost:5000"
	server := server.NewAPIServer(listenAddr, *store)
	server.ListenAndServe()
}
