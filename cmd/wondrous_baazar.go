package main

import (
	"log"
	"wondrousBaazar/http"
)

func main() {
	log.Printf("Starting Server")

	s := http.NewServer(":3000")
	log.Panic(s.Listen())
}
