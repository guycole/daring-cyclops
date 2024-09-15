// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"testing"

	shared "github.com/guycole/daring-cyclops/shared"
)

func TestPlayerCreation(t *testing.T) {
	sugarLog := shared.ZapSetup(true)

	pmt := newPlayerManager(sugarLog)
	pmt.seedTestUsers()

	// test creation and selection
	pt1 := pmt.findPlayerByKey(newPlayerKey(testPlayer1))
	if pt1 == nil {
		t.Error("player1 find failure")
	}

	if pt1 != nil && pt1.name != "testPlayer1" {
		t.Error("player1 name")
	}

	if pt1 != nil && pt1.rank != lieutenantRank {
		t.Error("player1 rank")
	}

	pt1 = pmt.findPlayerByName("testPlayer1")
	if pt1 == nil {
		t.Error("player1 find failure")
	}

	pt2 := pmt.findPlayerByKey(newPlayerKey(testPlayer2))
	if pt2 == nil {
		t.Error("player2 find failure")
	}

	if pt2 != nil && pt2.name != "testPlayer2" {
		t.Error("player2")
	}

	if pt2 != nil && pt2.rank != captainRank {
		t.Error("player2 rank")
	}

	// test for duplicates

	_, err := pmt.addFreshPlayer("testPlayer1")
	if err == nil {
		t.Error("should be duplicate name failure")
	}
}
