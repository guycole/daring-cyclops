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
	boardDump(gt.board)

	for ndx := 0; ndx < maxStarGates; ndx++ {
		log.Println(gt.starGates[ndx])
	}
}
