package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prestonbourne/goserve/controllers"
)

func main() {

	port := 5000
	addr := fmt.Sprintf(":%v", port)
	fmt.Println("Server running on " + addr)
	defaultStore := &controllers.InMemoryPlayerStore{Scores: make(map[string]int)}
	server := &controllers.PlayerServer{Store: defaultStore}
	err := http.ListenAndServe(addr, server)
	if err != nil {
		log.Fatal(err)
	}

}
