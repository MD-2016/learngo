package main

import (
	"fmt"
	"log"
	"os"
	"websocketsPkg"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := websocketsPkg.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := websocketsPkg.NewTexasHoldem(websocketsPkg.BlindAlerterFunc(websocketsPkg.Alerter), store)
	cli := websocketsPkg.NewCLI(os.Stdin, os.Stdout, game)

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	cli.PlayPoker()
}
