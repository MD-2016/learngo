package cli

import (
	"commandline"
	"fmt"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := commandline.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play timeserver")
	fmt.Println("Type {Name} wins to record a win")
	commandline.NewCLI(store, os.Stdin).Playtimeserver()
}
