// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"testing"

	shared "github.com/guycole/daring-cyclops/shared"
)

func TestMinorCommands(t *testing.T) {
	sugarLog := shared.ZapSetup(true)

	pmt := newPlayerManager(sugarLog)
	pmt.seedTestUsers()

	pt1 := pmt.findPlayerByKey(newPlayerKey(testPlayer1))
	pt2 := pmt.findPlayerByKey(newPlayerKey(testPlayer2))

	// start a game
	gt, _ := newGame(0, sugarLog)
	sugarLog.Info(gt)

	gt.addPlayerToGame(pt1, roninShipName, blueTeam)
	gt.addPlayerToGame(pt2, tritonShipName, redTeam)

	ct := newCommand(usersCommand)
	ust := gt.usersCommand(ct)
	sugarLog.Info(ust)
	if len(ust) != 2 {
		t.Error("userSummaryType length failure:", len(ust))
	}

	ct = newCommand(timeCommand)
	utt := gt.timeCommand(ct)
	sugarLog.Info(utt)
}
