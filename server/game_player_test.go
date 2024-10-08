// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"testing"

	shared "github.com/guycole/daring-cyclops/shared"
)

func TestGamePlayer(t *testing.T) {
	sugarLog := shared.ZapSetup(true)

	const maxGames = uint16(1)
	gmt := newGameManager(maxGames, sugarLog)

	// start game without thread
	const sleepSeconds = uint16(0)
	gmt.runAllGames(sleepSeconds)

	// add players to game
	gmt.playerManager.seedTestUsers()

	// test game
	gt := gmt.pickGame()

	pt1 := gmt.playerManager.findPlayerByKey(newTokenKey(testPlayer1))
	gpt1 := gt.addPlayerToGame(pt1, roninShipName, blueTeam)
	gpt1.activeFlag = true

	pt2 := gmt.playerManager.findPlayerByKey(newTokenKey(testPlayer2))
	gpt2 := gt.addPlayerToGame(pt2, tritonShipName, redTeam)

	gpt2.activeFlag = false
	gt.testForPlayerEviction()

	// gt.gamePlayerArrayDumper()
}
