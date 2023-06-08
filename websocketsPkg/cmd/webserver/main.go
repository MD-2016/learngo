package main

import (
	"log"
	"net/http"
	"os"
	"websocketsPkg"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := websocketsPkg.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	game := websocketsPkg.NewTexasHoldem(websocketsPkg.BlindAlerterFunc(websocketsPkg.Alerter), store)

	server, err := websocketsPkg.NewPlayerServer(store, game)

	if err != nil {
		log.Fatalf("problem creating player server %v", err)
	}

	log.Fatal(http.ListenAndServe(":5000", server))
}
