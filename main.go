package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("wire tutorial\n")
	e, err := InitializeEvent("aaa")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
}
