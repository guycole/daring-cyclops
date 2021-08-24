package game

import (
	"log"
)

const maxTeamPlayers = 5
const maxPlayers = maxTeamPlayers * 2

// WorkerType main game structure
type WorkerType struct {
	//gameBoard   GameBoardType
	players     [maxPlayers]*PlayerType
	turnCounter int
	uuid        string
}

// NewGame creates new game (should be init)
func NewWorker(id string) *WorkerType {
	log.Println("new game:", id)

	wt := WorkerType{uuid: id}
	//	result.gameBoard = freshGameBoard()
	return &wt
}

/*
// PlayerAdd add fresh player to game
func PlayerAdd(pt *PlayerType, gt *GameType) {
	log.Println(pt)

	switch pt.team {
	case BlueTeam:
		for ndx := 0; ndx < maxTeamPlayers; ndx++ {
			if gt.playersBlue[ndx] == nil {
				log.Println("assign", ndx)
				gt.playersBlue[ndx] = pt
				return
			}
		}

		// TOODO throw
		log.Println("team full blue")
	case RedTeam:
		for ndx := 0; ndx < maxTeamPlayers; ndx++ {
			if gt.playersRed[ndx] == nil {
				log.Println("assign", ndx)
				gt.playersRed[ndx] = pt
				return
			}
		}

		//TODO throw
		log.Println("team full red")
	default:
		log.Println("unkown team")
	}
}
*/

/*
// PlayerDelete remove player from game
func PlayerDelete(target string, gt *GameType) {
	for ndx := 0; ndx < maxTeamPlayers; ndx++ {
		if strings.Compare(gt.playersBlue[ndx].uuid, target) == 0 {
			gt.playersBlue[ndx] = nil
			return
		}
	}

	log.Println("no match blue")

	for ndx := 0; ndx < maxTeamPlayers; ndx++ {
		if strings.Compare(gt.playersRed[ndx].uuid, target) == 0 {
			gt.playersRed[ndx] = nil
			return
		}
	}

	log.Println("no match red")
}
*/

/*
// PlayerFind remove player from game
func PlayerFind(target string, gt *GameType) *PlayerType {
	for ndx := 0; ndx < maxTeamPlayers; ndx++ {
		if strings.Compare(gt.playersBlue[ndx].uuid, target) == 0 {
			return gt.playersBlue[ndx]
		}
	}

	log.Println("no match blue")

	for ndx := 0; ndx < maxTeamPlayers; ndx++ {
		if strings.Compare(gt.playersRed[ndx].uuid, target) == 0 {
			return gt.playersRed[ndx]
		}
	}

	log.Println("no match red")
	return nil
}
*/

/*
// ShipAdd add fresh ship to game
func ShipAdd(st *ShipType, gt *GameType) {
	// find player
	// discover if legal ship type for rank
	// discover if legal ship population
	// discover if ship name available
	// add ship to hashmap
}

// ShipDelete remove ship from game
func ShipDelete(target string, gt *GameType) {
}
*/
