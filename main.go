package main

import (
	"log"

	"github.com/kingtingthegreat/top-fetch-old/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	server := server.Server()
	log.Println("Server running at http://localhost:8080")
	err := server.ListenAndServe()
	log.Fatal(err.Error())
}
