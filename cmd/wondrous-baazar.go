package main

import (
	"log"
	"wondrousBaazar/internal/http"
)

func main() {
	s := http.NewServer(":3000")
	log.Panic(s.Listen())
}
