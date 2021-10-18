// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"log"
	"strings"

	"github.com/google/uuid"
)

// gameType, one for each game
type gameType struct {
	age            int
	bluePopulation int
	blueScore      int
	redPopulation  int
	redScore       int
	gameKey        string
	gameId         int
}

const maxGames = 5

// gameArrayType contains all games
type gameArrayType [maxGames]*gameType

func newGame(gameId int, gameKey string) *gameType {
	result := gameType{gameId: gameId}

	if len(gameKey) > 0 {
		result.gameKey = gameKey
	} else {
		result.gameKey = uuid.NewString()
	}

	return &result
}

func gameAdd(gt *gameType, gat *gameArrayType) int {
	log.Printf("gameAdd:%d %s", gt.gameId, gt.gameKey)

	for ndx := 0; ndx < maxGames; ndx++ {
		if gat[ndx] == nil {
			gat[ndx] = gt
			return ndx
		}
	}

	return -1
}

func gameDelete(target string, gat *gameArrayType) int {
	log.Printf("ganeDelete:%s", target)

	for ndx := 0; ndx < maxGames; ndx++ {
		if gat[ndx] != nil {
			if strings.Compare(gat[ndx].gameKey, target) == 0 {
				gat[ndx] = nil
				return ndx
			}
		}
	}

	return -1
}

func gameDump(gat gameArrayType) {
	log.Println("=-=-=-= gameDump =-=-=-=")

	for ndx := 0; ndx < maxGames; ndx++ {
		if gat[ndx] == nil {
			log.Printf("%d nil", ndx)
		} else {
			log.Printf("%d %d %d %d %d %s", ndx, gat[ndx].gameKey, gat[ndx].age, gat[ndx].bluePopulation, gat[ndx].redPopulation, gat[ndx].gameKey)
		}
	}

	log.Println("=-=-=-= gameDump =-=-=-=")
}

func gameFind(target string, gat gameArrayType) int {
	for ndx := 0; ndx < maxGames; ndx++ {
		if gat[ndx] != nil {
			if strings.Compare(gat[ndx].gameKey, target) == 0 {
				return ndx
			}
		}
	}

	return -1
}
