// go mod init github.com/guycole/daring-cyclops/worker

package main

import (
	"encoding/json"
	"log"

	"github.com/guycole/daring-cyclops/worker/game"
)

// Banner splash message
const Banner = "Daring Cyclops Worker V0.0"

func main() {
	log.Println(Banner)

	worker := game.NewWorker("gameId")
	log.Println(worker)

	message1 := `{"command":["newPlayer", "player1uuid", "CaptainRank", "BlueTeam"]}`

	var result map[string]interface{}
	json.Unmarshal([]byte(message1), &result)

	game.DispatchCommand(message1, *worker)

	/*
		player1 := game.NewPlayer("player1", game.Player1, game.CaptainRank, game.BlueTeam)
		log.Println(player1)
		game.PlayerAdd(player1, game1)
		playerTest := game.PlayerFind(game.Player1, game1)
		log.Println(playerTest)
	*/

	/*
		ship1 := game.NewShip("shipName", game.Player1, game.FighterShip, game.BlueTeam)
		game.ShipAdd(ship1, game1)

		test := `{"command":["m", "one", "two", "three"]}`
		log.Println(test)

		var result map[string]interface{}
		json.Unmarshal([]byte(test), &result)
		log.Println(result)
		log.Println(result["command"])

		//zzz := game.NewJsonCommand(test, player1)

		rawCommand := game.NewRawCommand(game.Player1, test)
		log.Println(rawCommand)
	*/

	//	demoCommand := game.NewTextCommand("gate", "bogus")
	//	log.Println(demoCommand)

	//	game.DispatchCommand(demoCommand, *demoGame)

	//game.CommandGa("shipId", *demoGame)
}
