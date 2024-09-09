// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"testing"

	shared "github.com/guycole/daring-cyclops/shared"
)

func TestFacadeGamePlayer(t *testing.T) {
	sugarLog := shared.ZapSetup(true)

	pmt := newPlayerManager(sugarLog)
	gmt := newGameManager(pmt, sugarLog)
	ft := newFacade(0, gmt, sugarLog)

	pit1 := ft.playerAdd("player1")
	pit2 := ft.playerAdd("player2")

	// add player to game

	gsat := ft.gameSummary()
	gt := ft.findGame(gsat[0].key)

	gmt.addPlayerToGame(gt.key, pit1.key, blueTeam)
	gmt.addPlayerToGame(gt.key, pit2.key, redTeam)
}
