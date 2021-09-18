// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"log"
	"strings"

	"github.com/google/uuid"
)

// gameWorkerType, one for each game
type gameWorkerType struct {
	uuid string // game identifier
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
			log.Printf("%d %s", ndx, gat[ndx].uuid)
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
