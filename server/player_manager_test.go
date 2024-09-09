// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"testing"

	shared "github.com/guycole/daring-cyclops/shared"
)

func TestPlayerManager(t *testing.T) {
	sugarLog := shared.ZapSetup(true)

	pmt := newPlayerManager(sugarLog)
	pt1a := pmt.addFreshPlayer("player1")
	pt2a := pmt.addFreshPlayer("player2")

	pt1b := pmt.findPlayer(pt1a.key)
	pt2b := pmt.findPlayer(pt2a.key)

	if pt1a.name != pt1b.name {
		t.Error("player1")
	}

	if pt2a.name != pt2b.name {
		t.Error("player2")
	}
}
