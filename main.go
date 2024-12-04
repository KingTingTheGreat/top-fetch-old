package main

import (
	"embed"
	"log"

	"github.com/kingtingthegreat/top-fetch/server"

	"github.com/joho/godotenv"
)

//go:embed styles/output.css
var StaticStyles embed.FS

func main() {
	godotenv.Load()

	// log.Println("styles", styles)

	server := server.Server()
	log.Println("Server running at http://localhost:8080")
	err := server.ListenAndServe()
	log.Fatal(err.Error())
}
