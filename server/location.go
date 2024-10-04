// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package server

import (
	"errors"
	"math"
	"math/rand"
	"strconv"
)

// row column origin = 0,0 lower left corner of map
type locationType struct {
	row uint16
	col uint16
}

func newLocation(row, col uint16) *locationType {
	result := locationType{row: row, col: col}
	return &result
}

func randomLocation(limitRow, limitCol uint16) *locationType {
	r1 := rand.Intn(int(limitRow))
	c1 := rand.Intn(int(limitCol))
	return newLocation(uint16(r1), uint16(c1))
}

func stringLocation(row, col string) *locationType {
	r1, err1 := strconv.Atoi(row)
	c1, err2 := strconv.Atoi(col)

	if err1 == nil && err2 == nil {
		return newLocation(uint16(r1), uint16(c1))
	}

	return nil
}

func (origin *locationType) moveAbsolute(row, col uint16) error {
	if row >= maxBoardSideRow || col >= maxBoardSideCol {
		return errors.New("moveAbsolute: location out of bounds")
	}

	origin.row = row
	origin.col = col

	return nil
}

func (origin *locationType) moveRelative(row, col int16) error {
	r1 := int16(origin.row) + row
	c1 := int16(origin.col) + col

	if r1 < 0 || c1 < 0 {
		return errors.New("moveRelative: negative location")
	}

	r2 := uint16(r1)
	c2 := uint16(c1)

	if r2 >= maxBoardSideRow || c2 >= maxBoardSideCol {
		return errors.New("moveRelative: location out of bounds")
	}

	origin.row = r2
	origin.col = c2

	return nil
}

func (origin *locationType) getDistance(destination *locationType) uint16 {
	var dx, dy float64

	if destination.col > origin.col {
		dx = float64(destination.col - origin.col)
	} else {
		dx = float64(origin.col - destination.col)
	}

	if destination.row > origin.row {
		dy = float64(destination.row - origin.row)
	} else {
		dy = float64(origin.row - destination.row)
	}

	result := math.Hypot(dx, dy)
	return uint16(result)
}

/*
   map origin lower left 0, 0

   1 2 3 (gate indices and relative locations)
   4 5 6
   7 8 9
*/
// return index of target location relative to origin location
func (origin *locationType) testForAdjacency(target *locationType) uint16 {
	var row, col uint16
	for ndx := uint16(1); ndx < 10; ndx++ {
		switch ndx {
		case 1:
			col = origin.col - 1
			row = origin.row + 1
		case 2:
			col = origin.col
			row = origin.row + 1
		case 3:
			col = origin.col + 1
			row = origin.row + 1
		case 4:
			col = origin.col - 1
			row = origin.row
		case 5:
			col = origin.col
			row = origin.row
		case 6:
			col = origin.col + 1
			row = origin.row
		case 7:
			col = origin.col - 1
			row = origin.row - 1
		case 8:
			col = origin.col
			row = origin.row - 1
		case 9:
			col = origin.col + 1
			row = origin.row - 1
		}

		if col == target.col && row == target.row {
			return ndx
		}
	}

	return 0
}
