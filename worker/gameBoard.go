package main

import (
	"log"
)

type boardTypeEnum int

const (
	unknownBoardType boardTypeEnum = iota
	emptyBoard01                   // no mines, planets, ships, stargates, voids
	standardBoard01                // standard game map w/planets, stargates and voids
	randomBoard01                  // standard game map w/randomly generated planets
)

// must match order for boardTypeEnum
func (bte boardTypeEnum) string() string {
	return [...]string{"Unknown", "empty01", "standard01", "random01"}[bte]
}

const maxBoardSideX = 75
const maxBoardSideY = 75

type gameBoardType struct {
	boardType boardTypeEnum

	boardArray [maxBoardSideX][maxBoardSideY]boardCellType
}

func newGameBoard(bt boardTypeEnum) gameBoardType {
	log.Println("fresh game board")

	var gb gameBoardType
	gb.boardType = bt

	return gb
}

func boardDump(gt gameType) {
	log.Println("=-=-=-= boardDump =-=-=-=")

	boardType := gt.gameBoard.boardType.string()
	log.Printf("boardType:%s", boardType)

	/*
		for ndx := 0; ndx < maxPlayers; ndx++ {
			rank := gt.players[ndx].rank.string()
			team := gt.players[ndx].team.string()

			log.Printf("%d %t %s %s %s %s", ndx, gt.players[ndx].active, gt.players[ndx].name, rank, team, gt.players[ndx].uuid)
		}
	*/

	log.Println("=-=-=-= boardDump =-=-=-=")
}

func boardGenerator(gt *gameType) {
	boardType := gt.gameBoard.boardType

	switch boardType {
	case emptyBoard01:
		log.Println("generating empty board")
		break
	case standardBoard01:
		log.Println("generating standard board")
		addPlanets(gt)
		addStarGates(gt)
	case randomBoard01:
		log.Println("generating random board")
		addRandomPlanets(gt)
	default:
		log.Println("unsupported boardType in boardGenerator")
	}
}

func addPlanets(gt *gameType) {
}

func addStarGates(gt *gameType) {
	/*
		for ndx := 0; ndx < 9; ndx++ {
			sg := newStarGate(ndx)
			log.Println(sg)

			gb.StarGates[sg.uuid] = *sg
		}
	*/
}

func addRandomPlanets(gt *gameType) {

}
