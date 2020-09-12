package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../environments/dev.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		prifle := os.Getenv("PROFILE")
		log.Print("Profile: ", prifle)
	}
}
