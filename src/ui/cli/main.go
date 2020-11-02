package main

import (
	"fmt"
	"log"
	"os"

	"github.com/berlim/bridgex/src/core/usecases/helpers"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("parameters missing")
		fmt.Println("ex: <comand> <option>")
		return
	}

	_, err := helpers.LoadData("development")

	if err != nil {
		log.Fatalf("ERRO: %v", err)
	}

	fmt.Printf("aeee %v", 30)
}
