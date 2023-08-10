package main

import (
	"log"

	"github.com/ChouE/go-commandline-app/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
