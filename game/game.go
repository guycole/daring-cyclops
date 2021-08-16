package game

import (
	"log"

	"github.com/google/uuid"
)

// Game describes game
type GameType struct {
	active       bool
	gameBoard    GameBoardType
	score        scoreType
	sequentialID int
	uuid         string
}

// NewGame creates new game
func NewGame(id int) *GameType {
	log.Println("fresh game")
	result := GameType{active: true, sequentialID: id}

	result.sequentialID = id
	result.gameBoard = freshGameBoard()
	result.uuid = uuid.NewString()

	return &result
}
