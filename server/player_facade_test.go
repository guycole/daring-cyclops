// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"testing"

	shared "github.com/guycole/daring-cyclops/shared"
)

func TestFacadeGamePlayer(t *testing.T) {
	sugarLog := shared.ZapSetup(true)

	gmt := newGameManager(sugarLog)
	if gmt.playerManager == nil {
		t.Error("player manager failure")
	}

	//	gmt.runAllGames()

	ft := newFacade(0, gmt, sugarLog)

	ft.playerAdd("player1")
	//pit2a := ft.playerAdd("player2")

	/*
	   pit1b := ft.playerGet(pit1a.key)
	   pit2b := ft.playerGet(pit2a.key)

	   	if pit1a.key.key != pit1b.key.key {
	   		t.Error("player key failure:", pit1a.key.key, pit1b.key.key)
	   	}

	   	if pit2a.key.key != pit2b.key.key {
	   		t.Error("player key failure:", pit2a.key.key, pit2b.key.key)
	   	}
	*/
}
