package main

import (
	"fmt"

	"github.com/prestonbourne/goserve/services"
)

func main() {
	services.ReadWholeFile("file.json")
	fmt.Println("Exit Code 0")
}
