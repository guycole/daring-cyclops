package server

import (
	"testing"

	shared "github.com/guycole/daring-cyclops/shared"
)

func TestFacadeGameSummary(t *testing.T) {
	sugarLog := shared.ZapSetup(true)

	gmt, _ := newGameManager(sugarLog)
	ft, _ := newFacade(0, gmt, sugarLog)

	gsat := ft.gameSummary()

	if len(gsat) != int(maxGames) {
		t.Error("gameSummary length failure:", gsat)
	}

	for _, gst := range gsat {
		if len(gst.key.key) < 36 {
			t.Error("gameSummary key failure:", gst)
		}
	}

	// now test game select

	target := gsat[0].key
	candidate := ft.findGame(target)

	if candidate.key.key != target.key {
		t.Error("gameSelect failure:", candidate)
	}
}
