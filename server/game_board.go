// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import "fmt"

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

const maxBoardSideX = uint16(75)
const maxBoardSideY = uint16(75)

type boardArrayType [maxBoardSideX][maxBoardSideY]*boardCellType

func newGameBoard(bte boardTypeEnum) *boardArrayType {
	var bat boardArrayType

	for yy := uint16(0); yy < maxBoardSideY; yy++ {
		for xx := uint16(0); xx < maxBoardSideX; xx++ {
			bat[yy][xx] = newBoardCell()
		}
	}

	if bte != emptyBoard {
		fmt.Println("fix me only empty board supported")
	}

	return &bat
}

/*
func (gt *gameType) boardGenerator() {
	gt.board = newBoard()

	switch gt.boardType {
	case emptyBoard:
			log.Println("generating empty board")
	case standardBoard:
			log.Println("generating standard board")
			//              starGatesAdd(&gt.starGates, &gt.board)
			//              starsAdd(&gt.stars, &gt.board)
			//              planetsAdd(&gt.planets, &gt.board)
	default:
		log.Println("unsupported boardType in boardGenerator")
	}
}
*/
