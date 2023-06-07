package main

import (
	"fmt"
	"log"
	"os"
	"timeserver"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := timeserver.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := timeserver.NewTexasHoldem(timeserver.BlindAlerterFunc(timeserver.StdOutAlerter), store)
	cli := timeserver.NewCLI(os.Stdin, os.Stdout, game)

	fmt.Println("Let's play timeserver")
	fmt.Println("Type {Name} wins to record a win")
	cli.PlayPoker()
}
