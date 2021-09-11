// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found
// in the LICENSE file.
package main

import (
	"log"
	"math"
	"math/rand"
)

//row column origin = 0,0 lower left corner of map
type locationType struct {
	xx int // column
	yy int // row
}

func newLocation(y, x int) *locationType {
	result := locationType{yy: y, xx: x}
	return &result
}

func getDistance(origin, destination *locationType) int {
	var x = float64(origin.xx - destination.xx)
	var y = float64(origin.yy - destination.yy)
	return int(math.Round(math.Sqrt(x*x + y*y)))
}

func randomLocation(limitY, limitX int) *locationType {
	xx := rand.Intn(limitX)
	yy := rand.Intn(limitY)
	return newLocation(yy, xx)
}

// return a random location for stars and planets
func randomCelestialLocation(bat boardArrayType) *locationType {
	for ndx := 0; ndx < 100; ndx++ {
		position := randomLocation(maxBoardSideY, maxBoardSideX)

		// cannot have celestial objects adjacent to stargates
		_, locNdx := starGateAdjacent(position)
		if locNdx >= 0 {
			//log.Printf("stargate adjacent:%d %d %d", gateNdx, locNdx, ndx)
			continue
		}

		boardCell := bat[position.yy][position.xx]
		if !testForCelestial(*boardCell) {
			return position
		}
	}

	log.Println("unable to generate random celestial location")

	return nil
}

// return a random location for ships
func randomShipLocation(bat boardArrayType) *locationType {
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

/*
   map origin lower left 1, 1

   0 1 2  (gate indices and relative locations)
   3 4 5
   6 7 8
*/
// return index of test location relative to reference location
func testForAdjacency(refPos, testPos *locationType) int {
	var x, y int

	for ndx := 0; ndx < 9; ndx++ {
		switch ndx {
		case 0:
			x = refPos.xx - 1
			y = refPos.yy + 1
		case 1:
			x = refPos.xx
			y = refPos.yy + 1
		case 2:
			x = refPos.xx + 1
			y = refPos.yy + 1
		case 3:
			x = refPos.xx - 1
			y = refPos.yy
		case 4: // should never match
			x = refPos.xx
			y = refPos.yy
		case 5:
			x = refPos.xx + 1
			y = refPos.yy
		case 6:
			x = refPos.xx - 1
			y = refPos.yy - 1
		case 7:
			x = refPos.xx
			y = refPos.yy - 1
		case 8:
			x = refPos.xx + 1
			y = refPos.yy - 1
		}

		if x == testPos.xx && y == testPos.yy {
			return ndx
		}
	}

	return -1
}
