// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found
// in the LICENSE file.
package main

import (
	"testing"
)

func TestPlayerRank(t *testing.T) {
	tests := []struct {
		candidate string
		answer    playerRankEnum
	}{
		{"cadet", cadetRank},
		{"admiral", admiralRank},
		{"bogus", unknownRank},
	}
	for _, ndx := range tests {
		result := findPlayerRank(ndx.candidate)
		if result != ndx.answer {
			t.Errorf("findPlayerRank(%s) failure", ndx.candidate)
		}
	}
}

func TestPlayerTeam(t *testing.T) {
	tests := []struct {
		candidate string
		answer    playerTeamEnum
	}{
		{"neutral", neutralTeam},
		{"red", redTeam},
		{"bogus", unknownTeam},
	}
	for _, ndx := range tests {
		result := findPlayerTeam(ndx.candidate)
		if result != ndx.answer {
			t.Errorf("findPlayerTeam(%s) failure", ndx.candidate)
		}
	}
}

func TestNewOkPlayer(t *testing.T) {
	result, err := newPlayer(testPlayerName1, testPlayerID1, "cadet", "blue")
	if err != nil {
		t.Errorf("newPlayer error:%s", err)
	}

	if result != nil {
		if result.name != testPlayerName1 {
			t.Error("newPlayer name failure")
		}

		if result.rank != cadetRank {
			t.Error("newPlayer rank failure")
		}

		if result.team != blueTeam {
			t.Error("newPlayer team failure")
		}

		if result.uuid != testPlayerID1 {
			t.Error("newPlayer id failure")
		}
	} else {
		t.Error("newPlayer returns nil")
	}
}

func TestNewBadPlayer01(t *testing.T) {
	result, err := newPlayer("", testPlayerID2, "cadet", "blue")
	if err == nil {
		t.Error("newPlayer error:expecting bad player")
	}

	if result != nil {
		t.Error("newPlayer error expecting nil")
	}
}

func TestNewBadPlayer02(t *testing.T) {
	result, err := newPlayer(testPlayerName2, "", "cadet", "blue")
	if err == nil {
		t.Error("newPlayer error:expecting bad id")
	}

	if result != nil {
		t.Error("newPlayer error expecting nil")
	}
}

func TestNewBadPlayer03(t *testing.T) {
	result, err := newPlayer(testPlayerName2, testPlayerID2, "", "blue")
	if err == nil {
		t.Error("newPlayer error:expecting bad rank")
	}

	if result != nil {
		t.Error("newPlayer error expecting nil")
	}
}

func TestNewBadPlayer04(t *testing.T) {
	result, err := newPlayer(testPlayerName2, testPlayerID2, "cadet", "")
	if err == nil {
		t.Error("newPlayer error:expecting bad team")
	}

	if result != nil {
		t.Error("newPlayer error expecting nil")
	}
}

func TestPlayerArray(t *testing.T) {
	var pat playerArrayType

	bluePopulation, redPopulation := playerCensus(pat)
	if bluePopulation != 0 && redPopulation != 0 {
		t.Errorf("playerCensus error:%d %d", bluePopulation, redPopulation)
	}

	np1 := testPlayer1()
	if np1 == nil {
		t.Error("testPlayer1 returns nil")
	}

	np2 := testPlayer2()
	if np2 == nil {
		t.Error("testPlayer2 returns nil")
	}

	// add player to player array, should be first array element
	ndx := playerAdd(np1, &pat)
	if ndx != 0 {
		t.Errorf("playerAdd returns wrong index %d", ndx)
	}

	// add player to player array, should be second array element
	ndx = playerAdd(np2, &pat)
	if ndx != 1 {
		t.Errorf("playerAdd returns wrong index %d", ndx)
	}

	bluePopulation, redPopulation = playerCensus(pat)
	if bluePopulation != 1 && redPopulation != 1 {
		t.Errorf("playerCensus error:%d %d", bluePopulation, redPopulation)
	}

	playerDump(pat)

	ndx = playerFind(testPlayerID1, pat)
	if ndx != 0 {
		t.Errorf("playerFind returns wrong index %d", ndx)
	}

	ndx = playerFind(testPlayerID2, pat)
	if ndx != 1 {
		t.Errorf("playerFind returns wrong index %d", ndx)
	}

	ndx = playerFind("bogus", pat)
	if ndx >= 0 {
		t.Errorf("playerFind returns wrong index %d", ndx)
	}

	ndx = playerDelete(testPlayerID1, &pat)
	if ndx != 0 {
		t.Errorf("playerDelete returns wrong index %d", ndx)
	}

	bluePopulation, redPopulation = playerCensus(pat)
	if bluePopulation != 0 && redPopulation != 1 {
		t.Errorf("playerCensus error:%d %d", bluePopulation, redPopulation)
	}
}

func TestCreateDeletePlayer(t *testing.T) {
	ct := commandType{player: testPlayerID2, request: "requestId"}
	ct.args = []string{"playerCreate", testPlayerName1, "captain", "blue"}
	ct.command = playerCreateCommand

	gt := gameType{uuid: "gameId"}

	err := commandPlayerCreate(ct, &gt)
	if err != nil {
		t.Errorf("commandCreatePlayer error:%s", err)
	}

	//playerDump(gt.players)

	bluePopulation, redPopulation := playerCensus(gt.players)
	if bluePopulation != 1 && redPopulation != 0 {
		t.Errorf("playerCensus error:%d %d", bluePopulation, redPopulation)
	}

	err = commandPlayerCreate(ct, &gt)
	if err == nil {
		t.Errorf("commandCreatePlayer should have duplicate error")
	}

	// playerDump(gt.players)

	ct = commandType{player: testPlayerID2, request: "requestId"}
	ct.args = []string{"playerDelete"}
	ct.command = playerDeleteCommand

	commandPlayerDelete(ct, &gt)

	// playerDump(gt.players)
}
