package main

import (
	"storage-links-app/repository"
	"storage-links-app/web"
)

func main() {
	repository.ConnectDB()
	web.StartListening()
}
