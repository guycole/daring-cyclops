// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found
// in the LICENSE file.
package main

import (
	"log"
)

// boardTypeEnum
type boardTypeEnum int

const (
	unknownBoardType boardTypeEnum = iota
	emptyBoard                     // no mines, planets, ships, stargates, voids
	standardBoard                  // standard game map w/planets, stargates and voids
	randomBoard                    // standard game map w/randomly generated planets
)

// must match order for boardTypeEnum
func (bte boardTypeEnum) string() string {
	return [...]string{"unknown", "empty", "standard", "random"}[bte]
}

const maxBoardSideX = 75
const maxBoardSideY = 75

type boardArrayType [maxBoardSideX][maxBoardSideY]*boardCellType

// newBoard create fresh empty game board
func newBoard() boardArrayType {
	var bat boardArrayType

	for yy := 0; yy < maxBoardSideY; yy++ {
		for xx := 0; xx < maxBoardSideX; xx++ {
			bat[yy][xx] = newBoardCell()
		}
	}

	return bat
}

func boardDump(bat boardArrayType) {
	log.Println("=-=-=-= boardDump =-=-=-=")

	var buffer string

	for yy := maxBoardSideY - 1; yy >= 0; yy-- {
		buffer = ""

		for xx := 0; xx < maxBoardSideX; xx++ {
			if bat[yy][xx] == nil {
				buffer += "oo"
			} else {
				buffer += boardCellToken(*bat[yy][xx])
			}
		}

		log.Printf("%2.2d %s", yy+1, buffer)
	}

	log.Println("=-=-=-= boardDump =-=-=-=")
}

func randomLocation3() *locationType {
	//xx := rand.Intn(limitX)
	//yy := rand.Intn(limitY)
	//return newLocation(yy, xx)
	return nil
}

func boardGenerator(gt *gameType) {
	switch gt.boardType {
	case emptyBoard:
		log.Println("generating empty board")
		break
	case standardBoard:
		log.Println("generating standard board")
		//		addPlanets(gt)
		//		addStars(gt)
		addStarGates(gt)
	case randomBoard:
		log.Println("generating random board")
	//	addRandomPlanets(gt)
	default:
		log.Println("unsupported boardType in boardGenerator")
	}
}

func addPlanets(gt *gameType) {
}

func addStars(gt *gameType) {
}

func addStarGates(gt *gameType) {
	for ndx := 0; ndx < maxStarGates; ndx++ {
		sg := newStarGate(ndx)
		gt.starGates[ndx] = sg
		gt.board[sg.position.yy][sg.position.xx].starGate = true
		gt.board[sg.position.yy][sg.position.xx].starGateID = sg.uuid
	}
}

func addRandomPlanets(gt *gameType) {

}
