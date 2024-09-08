package server

import (
	"testing"

	shared "github.com/guycole/daring-cyclops/shared"
)

func TestFacadeGamePlayer(t *testing.T) {
	sugarLog := shared.ZapSetup(true)

	gmt, _ := newGameManager(sugarLog)
	ft, _ := newFacade(0, gmt, sugarLog)

	pit1, _ := ft.playerIdentityNew("player1")
	pit2, _ := ft.playerIdentityNew("player2")

	temp1, _ := ft.playerIdentityGet(pit1.key)
	temp2, _ := ft.playerIdentityGet(pit2.key)

	if temp1.name != "player1" {
		t.Error("player1")
	}

	if temp2.name != "player2" {
		t.Error("player2")
	}

	//	sugarLog.Info(pit1)
	//	sugarLog.Info(pit2)

	// add player to game

	gsat := ft.gameSummary()
	gt := ft.findGame(gsat[0].key)

	gmt.addPlayerToGame(gt.key, pit1.key, blueTeam)
	gmt.addPlayerToGame(gt.key, pit2.key, redTeam)

}
