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

	ct := newCommand(userCommand, pt1.key)
	gt.userCommand(ct)

	if ct.destinationPlayerKeys == nil {
		t.Error("destinationPlayerKeys failure")
	}

	if len(ct.destinationPlayerKeys) != 1 {
		t.Error("destinationPlayerKeys length failure:", len(ct.destinationPlayerKeys))
	}

	if ct.destinationPlayerKeys[0] != pt1.key {
		t.Error("destinationPlayerKeys key failure:", ct.destinationPlayerKeys[0])
	}

	ur := ct.userResponse
	if ur == nil {
		t.Error("userResponse failure")
	}

	if len(ur.ust) != 2 {
		t.Error("userSummaryType length failure:", len(ur.ust))
	}
}
