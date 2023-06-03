package main

import (
	"log"
	"net/http"

	"learngo/go-specs-greet/adapters/httpserver"
)

func main() {
	log.Fatal(http.ListenAndServe(":8000", httpserver.NewHandler()))
}
