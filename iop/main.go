package iop

import (
	"log"
	"net/http"
	"os"
)

const db = "game.db.json"

func main() {
	dbf, err := os.OpenFile(db, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", db, err)
	}

	store, err := NewFileSystemPlayerStore(dbf)

	if err != nil {
		log.Fatalf("problem creating file system player store %v", err)
	}

	server := NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":5000", server))
}
