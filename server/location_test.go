// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"testing"
)

func TestLocationAdjacency(t *testing.T) {
	const originRow = uint16(4)
	const originCol = uint16(4)

	lt0 := newLocation(originRow+0, originCol+0)

	lt1 := newLocation(originRow+1, originCol-1)
	if lt0.testForAdjacency(lt1) != 1 {
		t.Error("bad adjacency")
	}

	lt2 := newLocation(originRow+1, originCol+0)
	if lt0.testForAdjacency(lt2) != 2 {
		t.Error("bad adjacency")
	}

	lt3 := newLocation(originRow+1, originCol+1)
	if lt0.testForAdjacency(lt3) != 3 {
		t.Error("bad adjacency")
	}

	lt4 := newLocation(originRow, originCol-1)
	if lt0.testForAdjacency(lt4) != 4 {
		t.Error("bad adjacency")
	}

	lt5 := newLocation(originRow, originCol+0)
	if lt0.testForAdjacency(lt5) != 5 {
		t.Error("bad adjacency")
	}

	lt6 := newLocation(originRow, originCol+1)
	if lt0.testForAdjacency(lt6) != 6 {
		t.Error("bad adjacency")
	}

	lt7 := newLocation(originRow-1, originCol-1)
	if lt0.testForAdjacency(lt7) != 7 {
		t.Error("bad adjacency")
	}

	lt8 := newLocation(originRow-1, originCol+0)
	if lt0.testForAdjacency(lt8) != 8 {
		t.Error("bad adjacency")
	}

	lt9 := newLocation(originRow-1, originCol+1)
	if lt0.testForAdjacency(lt9) != 9 {
		t.Error("bad adjacency")
	}
}

func TestLocationMove(t *testing.T) {
	const originRow = uint16(4)
	const originCol = uint16(4)

	lt0 := newLocation(originRow+0, originCol+0)

	lt0.moveAbsolute(1, 2)
	if lt0.row != 1 && lt0.col != 2 {
		t.Error("bad move")
	}

	err := lt0.moveAbsolute(maxBoardSideRow, maxBoardSideCol)
	if err == nil {
		t.Error("bad move")
	}

	lt0.moveRelative(1, 1)
	if lt0.row != 2 && lt0.col != 3 {
		t.Error("bad move")
	}

	lt0.moveRelative(-2, -3)
	if lt0.row != 0 && lt0.col != 0 {
		t.Error("bad move")
	}

	err = lt0.moveRelative(-2, -3)
	if err == nil {
		t.Error("bad move")
	}
}

func TestLocationDistance(t *testing.T) {
	const originRow = uint16(4)
	const originCol = uint16(4)

	lt0 := newLocation(originRow+0, originCol+0)
	if lt0.getDistance(lt0) != 0 {
		t.Error("bad distance")
	}

	lt1 := newLocation(originRow+0, originCol+4)
	if lt0.getDistance(lt1) != 4 {
		t.Error("bad distance")
	}

	lt2 := newLocation(originRow+4, originCol+4)
	if lt0.getDistance(lt2) != 5 {
		t.Error("bad distance")
	}

	lt3 := newLocation(originRow+4, originCol+0)
	if lt0.getDistance(lt3) != 4 {
		t.Error("bad distance")
	}

	lt4 := newLocation(originRow+4, originCol-4)
	if lt0.getDistance(lt4) != 5 {
		t.Error("bad distance")
	}

	lt5 := newLocation(originRow-0, originCol-4)
	if lt0.getDistance(lt5) != 4 {
		t.Error("bad distance")
	}

	lt6 := newLocation(originRow-4, originCol-4)
	if lt0.getDistance(lt6) != 5 {
		t.Error("bad distance")
	}

	lt7 := newLocation(originRow-4, originCol-0)
	if lt0.getDistance(lt7) != 4 {
		t.Error("bad distance")
	}

	lt8 := newLocation(originRow+4, originCol-4)
	if lt0.getDistance(lt8) != 5 {
		t.Error("bad distance")
	}
}
