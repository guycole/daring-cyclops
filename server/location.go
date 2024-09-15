// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package server

import (
	"math"
	"math/rand"
	"strconv"
)

// row column origin = 0,0 lower left corner of map
type locationType struct {
	xx int // column
	yy int // row
}

func newLocation(y, x int) *locationType {
	result := locationType{yy: y, xx: x}
	return &result
}

func randomLocation(limitY, limitX int) *locationType {
	xx := rand.Intn(limitX)
	yy := rand.Intn(limitY)
	return newLocation(yy, xx)
}

func stringLocation(y, x string) *locationType {
	yy, err1 := strconv.Atoi(y)
	xx, err2 := strconv.Atoi(x)

	if err1 == nil && err2 == nil {
		return newLocation(yy, xx)
	}

	return nil
}

func (origin *locationType) getDistance(destination *locationType) int {
	dx := float64(destination.xx - origin.xx)
	dy := float64(destination.yy - origin.yy)
	result := math.Hypot(dx, dy)
	return int(result)
}

/*
   map origin lower left 1, 1

   0 1 2  (gate indices and relative locations)
   3 4 5
   6 7 8
*/
// return index of target location relative to origin location
func (origin *locationType) testForAdjacency(target *locationType) int {
	var x, y int

	for ndx := 0; ndx < 9; ndx++ {
		switch ndx {
		case 0:
			x = origin.xx - 1
			y = origin.yy + 1
		case 1:
			x = origin.xx
			y = origin.yy + 1
		case 2:
			x = origin.xx + 1
			y = origin.yy + 1
		case 3:
			x = origin.xx - 1
			y = origin.yy
		case 4:
			// should never match
			continue
		case 5:
			x = origin.xx + 1
			y = origin.yy
		case 6:
			x = origin.xx - 1
			y = origin.yy - 1
		case 7:
			x = origin.xx
			y = origin.yy - 1
		case 8:
			x = origin.xx + 1
			y = origin.yy - 1
		}

		if x == target.xx && y == target.yy {
			return ndx
		}
	}

	return -1
}
