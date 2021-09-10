// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found
// in the LICENSE file.
package main

import (
	"log"
	"testing"
)

func TestGameBoard1(t *testing.T) {
	gt := newGame("testGame", standardBoard)

	//for ndx := 0; ndx < maxStarGates; ndx++ {
	//	log.Println(gt.starGates[ndx])
	//}

	//planetDump(gt.planets)
	//starDump(gt.stars)

	ns1 := testShip1(gt)
	if ns1 == nil {
		t.Error("testShip1 returns nil")
	}

	ns2 := testShip2(gt)
	if ns2 == nil {
		t.Error("testShip2 returns nil")
	}

	log.Println("------------")
	log.Println(ns1.position)
	log.Println(ns2.position)
	log.Println("------------")

	boardDump(gt.board)
}
