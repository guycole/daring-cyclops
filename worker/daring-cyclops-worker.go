// go mod init github.com/guycole/daring-cyclops/worker

package main

import (
	"log"

	"github.com/guycole/daring-cyclops/worker/game"
)

// Banner splash message
const Banner = "Daring Cyclops Worker V0.0"

func main() {
	log.Println(Banner)

	game1 := game.NewGame(123)
	log.Println(game1)

	player1 := game.NewPlayer("player1", "88265291-fcf5-47e6-ad41-e20fa712e0f7", game.CaptainRank, game.BlueTeam)
	game.PlayerAdd(player1, game1)

	ship1 := game.NewShip("shipName", "88265291-fcf5-47e6-ad41-e20fa712e0f7", game.FighterShip, game.BlueTeam)
	game.ShipAdd(ship1, game1)

	//	demoCommand := game.NewTextCommand("gate", "bogus")
	//	log.Println(demoCommand)

	//	game.DispatchCommand(demoCommand, *demoGame)

	//game.CommandGa("shipId", *demoGame)
}
