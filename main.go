package main

import (
	"github.com/joho/godotenv"
	"log"
	"storage-links-app/repository"
	"storage-links-app/web"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No found .env file")
	}
	repository.ConnectDB()
	web.StartListening()
}
