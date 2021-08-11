package manager

import (
	"log"

	"github.com/google/uuid"
)

const Banner = "Daring Cyclops V0.0"

const MaxGames = 5

type Game struct {
	Active       bool
	RandomId     string
	SequentialId int
}

var allGames [MaxGames]Game

var maxSequentialId = -1

func getSequentialId() int {
	var currentId = maxSequentialId + 1
	maxSequentialId = currentId
	return maxSequentialId
}

func setupGame(ndx int) {
	log.Println("setupGame:", ndx)
	allGames[ndx].Active = false
	allGames[ndx].RandomId = uuid.NewString()
	allGames[ndx].SequentialId = getSequentialId()
}

func setup() {
	for ndx := 0; ndx < MaxGames; ndx++ {
		setupGame(ndx)
	}
}

func Manager() {
	log.Println(Banner)
	setup()

	// fall into event loop
}
