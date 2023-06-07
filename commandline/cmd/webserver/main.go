package main

import (
	"log"
	"net/http"

	"commandline"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := commandline.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := commandline.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":5000", server))
}
