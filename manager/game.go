// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"log"
	"strings"

	"github.com/google/uuid"
)

const maxTeamPlayers = 5
const maxPlayers = maxTeamPlayers * 2

type playerArrayType [maxTeamPlayers]string

// gameWorkerType, one for each game
type gameWorkerType struct {
	blueTeam playerArrayType
	redTeam  playerArrayType
	uuid     string // game identifier
}

const maxGames = 5

// gameArrayType contains all games
type gameArrayType [maxGames]*gameWorkerType

func newGame() *gameWorkerType {
	result := gameWorkerType{}
	result.uuid = uuid.NewString()
	return &result
}

// gameAdd
func gameAdd(wt *gameWorkerType, gat *gameArrayType) int {
	log.Printf("gameAdd:%s", wt.uuid)

	for ndx := 0; ndx < maxGames; ndx++ {
		if gat[ndx] == nil {
			gat[ndx] = wt
			return ndx
		}
	}

	return -1
}

// gameDelete
func gameDelete(target string, gat *gameArrayType) int {
	log.Printf("ganeDelete:%s", target)

	for ndx := 0; ndx < maxGames; ndx++ {
		if gat[ndx] != nil {
			if strings.Compare(gat[ndx].uuid, target) == 0 {
				gat[ndx] = nil
				return ndx
			}
		}
	}

	return -1
}

// gameDump diagnostic
func gameDump(gat gameArrayType) {
	log.Println("=-=-=-= gameDump =-=-=-=")

	for ndx := 0; ndx < maxGames; ndx++ {
		if gat[ndx] == nil {
			log.Printf("%d nil", ndx)
		} else {
			bluePopulation, redPopulation := gamePlayerCensus(*gat[ndx])
			log.Printf("%d %d %d %s", ndx, bluePopulation, redPopulation, gat[ndx].uuid)
		}
	}

	log.Println("=-=-=-= gameDump =-=-=-=")
}

// gameFind
func gameFind(target string, gat gameArrayType) int {
	for ndx := 0; ndx < maxGames; ndx++ {
		if gat[ndx] != nil {
			if strings.Compare(gat[ndx].uuid, target) == 0 {
				return ndx
			}
		}
	}

	return -1
}

func gamePlayerCensus(gwt gameWorkerType) (int, int) {
	bluePopulation := 0
	redPopulation := 0

	for ndx := 0; ndx < maxTeamPlayers; ndx++ {
		if len(gwt.blueTeam[ndx]) > 0 {
			bluePopulation++
		}

		if len(gwt.redTeam[ndx]) > 0 {
			redPopulation++
		}
	}

	return bluePopulation, redPopulation
}

func gamePlayerAdd(player PlayerType, blueArray *playerArrayType, redArray *playerArrayType) {
	// TODO issue 19 enforce team size limits

	switch player.Team {
	case blueTeam:
		for ndx := 0; ndx < maxTeamPlayers; ndx++ {
			if len(blueArray[ndx]) < 1 {
				blueArray[ndx] = player.Name
				return
			}
		}
	case redTeam:
		for ndx := 0; ndx < maxTeamPlayers; ndx++ {
			if len(redArray[ndx]) < 1 {
				redArray[ndx] = player.Name
				return
			}
		}
	}
}

func gamePlayerDelete(player PlayerType, blueArray *playerArrayType, redArray *playerArrayType) {
	switch player.Team {
	case blueTeam:
		for ndx := 0; ndx < maxTeamPlayers; ndx++ {
			if strings.Compare(blueArray[ndx], player.Name) == 0 {
				blueArray[ndx] = ""
				return
			}
		}
	case redTeam:
		for ndx := 0; ndx < maxTeamPlayers; ndx++ {
			if strings.Compare(redArray[ndx], player.Name) == 0 {
				redArray[ndx] = ""
				return
			}
		}
	}
}
