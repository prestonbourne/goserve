package main

import (
	"fmt"
	"log"
	"net/http"

	controllers "github.com/prestonbourne/goserve/controllers"
)

func main() {
	fmt.Println("running")
	server := &controllers.PlayerServer{}
	log.Fatal(http.ListenAndServe(":5000", server))

}
