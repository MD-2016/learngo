package main

import (
	"log"
	"net/http"

	"github.com/MD-2016/go-specs-greet/adapters/httpserver"
)

func main() {
	log.Fatal(http.ListenAndServe(":8000", httpserver.NewHandler()))
}
