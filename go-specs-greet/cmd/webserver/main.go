package main

import (
	"log"
	"net/http"

	"github.com/MD-2016/go-specs-greet/adapters/webserver"
)

func main() {
	handle, err := webserver.NewHandler()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8081", handle))
}
