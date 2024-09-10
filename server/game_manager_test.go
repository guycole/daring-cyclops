// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"testing"

	shared "github.com/guycole/daring-cyclops/shared"
)

func TestGameManager(t *testing.T) {
	sugarLog := shared.ZapSetup(true)

	pmt := newPlayerManager(sugarLog)
	gmt := newGameManager(pmt, sugarLog)

	gmt.runAllGames()

	// ensure all games are running
	gsat := gmt.gameSummary()
	if len(gsat) != int(maxGames) {
		t.Error("gameSummary length failure:", gsat)
	}

	sugarLog.Info(gsat)

	for _, gst := range gsat {
		if len(gst.key.key) < 36 {
			t.Error("gameSummary key failure:", gst)
		}
	}

	// single game select
	target := gsat[0].key
	gt := gmt.findGame(target)

	if gt.key.key != target.key {
		t.Error("gameSelect failure:", gt)
	}

	// add user to game
	pt1 := pmt.addFreshPlayer("player1")
	pt2 := pmt.addFreshPlayer("player2")

	gmt.addPlayerToGame(target, pt1.key, "ship1", blueTeam)
	gmt.addPlayerToGame(target, pt2.key, "ship2", redTeam)

	// test all known players
	if len(gmt.playerManager.playerMap) != 2 {
		t.Error("playerMap length failure:", gmt.playerManager.playerMap)
	}

	// test players for game
	if len(gt.playerMap) != 2 {
		t.Error("playerMap length failure:", gt.playerMap)
	}
}
