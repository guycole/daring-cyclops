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

const maxBoardSideRow = uint16(75)
const maxBoardSideCol = uint16(75)

type boardArrayType [maxBoardSideRow][maxBoardSideCol]*boardCellType

func newGameBoard(bte boardTypeEnum) *boardArrayType {
	var bat boardArrayType

	for row := uint16(0); row < maxBoardSideRow; row++ {
		for col := uint16(0); col < maxBoardSideCol; col++ {
			bat[row][col] = newBoardCell()
		}
	}

	if bte != emptyBoard {
		fmt.Println("fix me only empty board supported")
	}

	switch bte {
	case emptyBoard:
		fmt.Println("generating empty board")
	case standardBoard:
		fmt.Println("generating standard board")
		//              starGatesAdd(&gt.starGates, &gt.board)
		//              starsAdd(&gt.stars, &gt.board)
		//              planetsAdd(&gt.planets, &gt.board)
	default:
		fmt.Println("unsupported boardType in boardGenerator")
	}

	return &bat
}
