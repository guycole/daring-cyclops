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
	standardBoard                  // standard game map w/stars, planets, stargates and voids
)

// must match order for boardTypeEnum
func (bte boardTypeEnum) string() string {
	return [...]string{"unknown", "empty", "standard"}[bte]
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

func boardGenerator(gt *gameType) {
	switch gt.boardType {
	case emptyBoard:
		log.Println("generating empty board")
	case standardBoard:
		log.Println("generating standard board")
		starGatesAdd(gt)
		starsAdd(gt)
		planetsAdd(gt)
	default:
		log.Println("unsupported boardType in boardGenerator")
	}
}
