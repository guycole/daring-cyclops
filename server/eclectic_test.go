// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"testing"

	shared "github.com/guycole/daring-cyclops/shared"
)

func TestEclectic(t *testing.T) {
	sugarLog := shared.ZapSetup(true)
	sugarLog.Info("eclectic test start xoxoxoxoxoxoxo")

	const maxGames = uint16(1)
	gmt := newGameManager(maxGames, sugarLog)

	// start game
	const sleepSeconds = uint16(0)
	gmt.runAllGames(sleepSeconds)

	// add players to game
	gmt.playerManager.seedTestUsers()

	// test game
	gt := gmt.pickGame()
	sugarLog.Info("user command test:", gt.key.key)

	pt1 := gmt.playerManager.findPlayerByKey(newPlayerKey(testPlayer1))
	gt.addPlayerToGame(pt1, roninShipName, blueTeam)

	pt2 := gmt.playerManager.findPlayerByKey(newPlayerKey(testPlayer2))
	gt.addPlayerToGame(pt2, tritonShipName, redTeam)

	ct := newCommand(usersCommand)
	//gt.usersCommand(ct)
	gt.enqueue(ct)
	gt.eclectic()

	//	for _, val := range ust {
	//		sugarLog.Infof("%s %s %t %d %d %s", val.name, val.rank.string(), val.active, val.age, val.score, val.team.string())
	//	}

	sugarLog.Info("eclectic test end xoxoxoxoxoxoxo")
}
