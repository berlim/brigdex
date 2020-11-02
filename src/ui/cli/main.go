package main

import (
	"fmt"
	"log"
	"os"

	"github.com/berlim/bridgex/src/core/usecases/cmd"
	"github.com/berlim/bridgex/src/core/usecases/helpers"
)

func main() {

	env := getEnv()

	if len(os.Args) < 3 {
		fmt.Println("parameters missing")
		fmt.Println("ex: <comand> <option>")
		return
	}

	configData, err := helpers.LoadData(env)

	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	err = cmd.SwitchCmd(configData)

	if err != nil {
		log.Fatalf("ERROR: %v \n", err)
	}
}

func getEnv() string {
	env := os.Getenv("APP_ENV")

	if len(env) < 1 {
		return "development"
	}

	return env
}
