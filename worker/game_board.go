// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"log"
)

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

func (bat boardArrayType) boardDump() {
	log.Println("=-=-=-= boardDump =-=-=-=")

	var buffer string

	for yy := maxBoardSideY - 1; yy >= 0; yy-- {
		buffer = ""

		for xx := 0; xx < maxBoardSideX; xx++ {
			if bat[yy][xx] == nil {
				buffer += "o"
			} else {
				buffer += boardCellToken(*bat[yy][xx])
			}
		}

		log.Printf("%2.2d %s", yy+1, buffer)
	}

	log.Println("=-=-=-= boardDump =-=-=-=")
}

// return a random location for ships
func (bat boardArrayType) randomShipLocation() *locationType {
	for ndx := 0; ndx < 100; ndx++ {
		position := randomLocation(maxBoardSideY, maxBoardSideX)
		boardCell := bat[position.yy][position.xx]
		if testForEmpty(*boardCell) {
			return position
		}
	}

	log.Println("unable to generate random ship location")

	return nil
}

// return a random location for stars and planets
func (bat boardArrayType) randomCelestialLocation() *locationType {
	for ndx := 0; ndx < 100; ndx++ {
		position := randomLocation(maxBoardSideY, maxBoardSideX)

		/*
			// cannot have celestial objects adjacent to stargates
			_, locNdx := starGateAdjacent(position)
			if locNdx >= 0 {
				//log.Printf("stargate adjacent:%d %d %d", gateNdx, locNdx, ndx)
				continue
			}
		*/

		boardCell := bat[position.yy][position.xx]
		if !testForCelestial(*boardCell) {
			return position
		}
	}

	log.Println("unable to generate random celestial location")

	return nil
}
