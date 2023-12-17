package main

import (
	"log"
	"todo/cmd"
	"todo/config"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("main: could not initialize config due to error ", err)
	}

	log.Printf("main: config loaded successfully %+v", config.Value)

	cmd.NewCLI().Execute()
}
