package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/prestonbourne/goserve/controllers"
	"github.com/prestonbourne/goserve/db"
)

func main() {

	db, err := db.NewPostgresStore(context.Background(), "postgresql://postgres:password@localhost:5432/postgres")
	if err != nil {
		fmt.Printf("%v", err)
	}

	db.GetPlayers()

	port := 5000
	addr := fmt.Sprintf(":%v", port)

	fmt.Println("Server running on " + addr)
	defaultStore := &controllers.InMemoryPlayerStore{Scores: make(map[string]int)}
	playerHandler := &controllers.PlayerServer{Store: defaultStore, DB: db}
	mux := http.NewServeMux()
	mux.Handle("/players", playerHandler)

	err = http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal(err)
	}

}
