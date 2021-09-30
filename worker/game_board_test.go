// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"testing"
)

func TestGameBoard1(t *testing.T) {
	gt := newGame("testGame", emptyBoard)

	//planetDump(gt.planets)
	//starDump(gt.stars)

	ns1 := gt.testShip1()
	if ns1 == nil {
		t.Error("testShip1 returns nil")
	}

	/*
		shipAdd(ns1, &gt.ships, &gt.board)

		ns2 := gt.testShip2()
		if ns2 == nil {
			t.Error("testShip2 returns nil")
		}

		shipAdd(ns2, &gt.ships, &gt.board)

		log.Println("------------")
		log.Println(ns1.position)
		log.Println(ns2.position)
		log.Println("------------")

		newLoc := newLocation(5, 5)
		shipMove(testShipUuid1, *newLoc, &gt.ships, &gt.board)

		newLoc = newLocation(3, 3)
		shipMove(testShipUuid1, *newLoc, &gt.ships, &gt.board)

		gt.board.boardDump()
	*/
}
