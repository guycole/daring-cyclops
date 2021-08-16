package game

import (
	"log"

	"github.com/google/uuid"
)

const maxBoardSideX = 75
const maxBoardSideY = 75

// GameBoard comment
type GameBoardType struct {
	// creation time
	UUID string

	BoardArray [maxBoardSideX][maxBoardSideY]*boardCellType

	Planets   map[string]planetType
	StarGates map[string]starGateType

	Ships map[string]shipType
}

func freshGameBoard() GameBoardType {
	log.Println("fresh game board")

	var gameBoard GameBoardType

	gameBoard.UUID = uuid.NewString()

	gameBoard.StarGates = make(map[string]starGateType)

	// initialize empty gameboard
	for yy := 0; yy < maxBoardSideY; yy++ {
		for xx := 0; xx < maxBoardSideX; xx++ {
			location := newLocation(yy, xx)
			gameBoard.BoardArray[yy][xx] = newBoardCell(location)

			/*
				var location = newLocation(yy, xx)
				var boardCell = newBoardCell(location)
				gameBoard.BoardArray[yy][xx] = boardCell
			*/
		}
	}

	addStarGates(gameBoard)

	// add planets, etc
	/*
		for yy := 0; yy < MaxBoardSideY; yy++ {
			for xx := 0; xx < MaxBoardSideX; xx++ {
				var boardCell = gameBoard.BoardArray[yy][xx]
				//test(boardCell)
			}
		}
	*/

	return gameBoard
}

func addStarGates(gb GameBoardType) {
	for ndx := 0; ndx < 9; ndx++ {
		sg := newStarGate(ndx)
		log.Println(sg)
	}

	/*
		// stargate 1
		boardCell := gb.BoardArray[12][34]

		setStarGate(boardCell, sg.uuid)
		gb.StarGates[sg.uuid] = *sg
	*/
}
