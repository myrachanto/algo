package main

import (
	"log"

	"github.com/myrachanto/algo/gobook/internal/server"
)

func main() {
	log.Println("Starting web server at port 8080")
	server.CreateServer()
}
