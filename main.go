package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ququzone/0g-storage-example/cmd"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cmd.Execute()
}
