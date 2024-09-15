// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"testing"

	shared "github.com/guycole/daring-cyclops/shared"
)

func TestGameManager(t *testing.T) {
	sugarLog := shared.ZapSetup(true)

	gmt := newGameManager(sugarLog)

	gmt.runAllGames()

	// ensure all games are running
	gsat := gmt.gameSummary()
	if len(gsat) != int(maxGames) {
		t.Error("gameSummary length failure:", gsat)
	}

	sugarLog.Info(gsat)

	// game with keys
	for _, gst := range gsat {
		if len(gst.key.key) < 36 {
			t.Error("gameSummary key failure:", gst)
		}
	}

	// add players to game
	gmt.playerManager.seedTestUsers()

	// single game select
	target := gsat[0].key
	gt := gmt.findGame(target)

	if gt.key.key != target.key {
		t.Error("gameSelect failure:", gt)
	}

	pt1 := gmt.playerManager.findPlayerByKey(newPlayerKey(testPlayer1))
	gt.addPlayerToGame(pt1, roninShipName, blueTeam)

	pt2 := gmt.playerManager.findPlayerByKey(newPlayerKey(testPlayer2))
	gt.addPlayerToGame(pt2, tritonShipName, redTeam)

	if len(gt.playerMap) != 2 {
		t.Error("playerMap length failure:", gt.playerMap)
	}

	if len(gt.shipMap) != 2 {
		t.Error("shipMap length failure:", gt.shipMap)
	}
}
