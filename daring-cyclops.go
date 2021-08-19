// go mod init github.com/guycole/daring-cyclops
// go install github.com/guycole/daring-cyclops

package main

import (
	"log"

	"github.com/guycole/daring-cyclops/game"
)

func main() {
	log.Println("Banner")
	demoGame := game.NewGame(123)
	log.Println(demoGame)

	demoCommand := game.NewTextCommand("gate", "bogus")
	log.Println(demoCommand)

	game.DispatchCommand(demoCommand, *demoGame)

	//game.CommandGa("shipId", *demoGame)
}
