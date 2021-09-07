// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found
// in the LICENSE file.
package main

import (
	"math"
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

/*
   map origin lower left 1, 1

   0 1 2  (gate indices and relative locations)
   3 4 5
   6 7 8
*/

// starGateAdjacent discover if next to a SG.  Returns index into starGateDestinations
func starGateAdjacent(shipPosition, starGatePosition *locationType) int {
	var x, y int

	for ndx := 0; ndx < 9; ndx++ {
		switch ndx {
		case 0:
			x = starGatePosition.xx - 1
			y = starGatePosition.yy + 1
		case 1:
			x = starGatePosition.xx
			y = starGatePosition.yy + 1
		case 2:
			x = starGatePosition.xx + 1
			y = starGatePosition.yy + 1
		case 3:
			x = starGatePosition.xx - 1
			y = starGatePosition.yy
		case 4: // should never match
			x = starGatePosition.xx
			y = starGatePosition.yy
		case 5:
			x = starGatePosition.xx + 1
			y = starGatePosition.yy
		case 6:
			x = starGatePosition.xx - 1
			y = starGatePosition.yy - 1
		case 7:
			x = starGatePosition.xx
			y = starGatePosition.yy - 1
		case 8:
			x = starGatePosition.xx + 1
			y = starGatePosition.yy - 1
		}

		temp := newLocation(y, x)
		if temp.xx == shipPosition.xx && temp.yy == shipPosition.yy {
			return ndx
		}

	}

	return -1
}
