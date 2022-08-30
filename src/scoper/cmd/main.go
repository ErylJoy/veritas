package main

import (
	"fmt"
	"os"

	service "github.com/eryljoy/veritas/src/scoper/internal"
)

func main() {
	fmt.Println("Hello Worlds")
	service.Useful()
	fileString, err := os.ReadFile("./main.c")
	if err != nil {
		return
	}

	fmt.Printf(string(fileString))
}
