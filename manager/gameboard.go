package manager

import "log"

type BoardTokenType int

const (
	VacantToken BoardTokenType = iota
	MineToken
	PlanetToken
	ShipToken
	StarGateToken
)

func (btt BoardTokenType) String() string {
	return [...]string{"Vacant", "Mine", "Planet", "Ship", "StarGate"}[btt]
}

const MaxBoardSideX = 75
const MaxBoardSideY = 75

// BoardCell describes a location on GameBoard
type BoardCell struct {
	AcheronVoid bool
	BlackHole   bool
	Planet      bool
	StarGate    bool

	Position Location
}

func getFreshBoardCell(position Location) BoardCell {
	var result BoardCell

	result.AcheronVoid = false
	result.BlackHole = false
	result.Planet = false
	result.StarGate = false

	result.Position = position

	return result
}

type GameBoard struct {
	// creation time
	Uuid string

	BoardArray [MaxBoardSideX][MaxBoardSideY]BoardCell
}

func freshGameBoard() GameBoard {
	log.Println("fresh game board")

	var gameBoard GameBoard

	// generate empty gameboard
	for yy := 0; yy < MaxBoardSideY; yy++ {
		for xx := 0; xx < MaxBoardSideX; xx++ {
			var location = getFreshLocation(yy, xx)
			var boardCell = getFreshBoardCell(location)
			gameBoard.BoardArray[yy][xx] = boardCell
		}
	}

	// add planets, etc
	for yy := 0; yy < MaxBoardSideY; yy++ {
		for xx := 0; xx < MaxBoardSideX; xx++ {
			var boardCell = gameBoard.BoardArray[yy][xx]
			//test(boardCell)
		}
	}

	return gameBoard
}
