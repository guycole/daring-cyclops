// go mod init github.com/guycole/daring-cyclops
// go install github.com/guycole/daring-cyclops

package main

import (
	"log"

	"github.com/guycole/daring-cyclops/game"
)

func main() {
	log.Println("Banner")
	game.NewGame(123)
	//manager.Manager()
}
