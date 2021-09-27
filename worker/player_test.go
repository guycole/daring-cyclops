// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"log"
	//"strings"
	"testing"
)

func TestPlayerRank(t *testing.T) {
	tests := []struct {
		candidate string
		answer    rankEnum
	}{
		{"cadet", cadetRank},
		{"admiral", admiralRank},
		{"bogus", unknownRank},
	}
	for _, ndx := range tests {
		result := findRank(ndx.candidate)
		if result != ndx.answer {
			t.Errorf("findRank(%s) failure", ndx.candidate)
		}
	}
}

func TestPlayerTeam(t *testing.T) {
	tests := []struct {
		candidate string
		answer    teamEnum
	}{
		{"neutral", neutralTeam},
		{"red", redTeam},
		{"bogus", unknownTeam},
	}
	for _, ndx := range tests {
		result := findTeam(ndx.candidate)
		if result != ndx.answer {
			t.Errorf("findTeam(%s) failure", ndx.candidate)
		}
	}
}

func TestNewOkPlayer(t *testing.T) {
	result, err := newPlayer(testPlayerName1, "cadet", "blue")
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
	} else {
		t.Error("newPlayer returns nil")
	}
}

func TestNewBadPlayer01(t *testing.T) {
	result, err := newPlayer("", "cadet", "blue")
	if err == nil {
		t.Error("newPlayer error:expecting bad player")
	}

	if result != nil {
		t.Error("newPlayer error expecting nil")
	}
}

func TestNewBadPlayer02(t *testing.T) {
	result, err := newPlayer(testPlayerName2, "", "blue")
	if err == nil {
		t.Error("newPlayer error:expecting bad rank")
	}

	if result != nil {
		t.Error("newPlayer error expecting nil")
	}
}

func TestNewBadPlayer03(t *testing.T) {
	result, err := newPlayer(testPlayerName2, "cadet", "")
	if err == nil {
		t.Error("newPlayer error:expecting bad team")
	}

	if result != nil {
		t.Error("newPlayer error expecting nil")
	}
}

func TestPlayerArray(t *testing.T) {
	var pat playerArrayType

	bluePopulation, redPopulation := pat.playerCensus()
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
	ndx := pat.playerAdd(np1)
	if ndx != 0 {
		t.Errorf("playerAdd returns wrong index %d", ndx)
	}

	// add player to player array, should be second array element
	ndx = pat.playerAdd(np2)
	if ndx != 1 {
		t.Errorf("playerAdd returns wrong index %d", ndx)
	}

	bluePopulation, redPopulation = pat.playerCensus()
	if bluePopulation != 1 && redPopulation != 1 {
		t.Errorf("playerCensus error:%d %d", bluePopulation, redPopulation)
	}

	pat.playerDump()

	ndx = pat.playerFind(testPlayerName1)
	if ndx != 0 {
		t.Errorf("playerFind returns wrong index %d", ndx)
	}

	ndx = pat.playerFind(testPlayerName2)
	if ndx != 1 {
		t.Errorf("playerFind returns wrong index %d", ndx)
	}

	ndx = pat.playerFind("bogus")
	if ndx >= 0 {
		t.Errorf("playerFind returns wrong index %d", ndx)
	}

	ndx = pat.playerDelete(testPlayerName1)
	if ndx != 0 {
		t.Errorf("playerDelete returns wrong index %d", ndx)
	}

	bluePopulation, redPopulation = pat.playerCensus()
	if bluePopulation != 0 && redPopulation != 1 {
		t.Errorf("playerCensus error:%d %d", bluePopulation, redPopulation)
	}
}

func TestCommandCreateDeletePlayer(t *testing.T) {
	// create player
	var commands1 commandArrayType
	commands1[0] = "playerCreate"
	commands1[1] = testPlayerName1
	commands1[2] = "captain"
	commands1[3] = "blue"

	nc1 := newCommand(testPlayerName1, "reqId1", 4, commands1)
	log.Println(nc1)
	/*
		// load command
		gt := newGame("testGame", emptyBoard)
		gt.commandStack.push(nc1)

		// run command
		gt.serviceCommandStack()

		// test for player add
		ndx := gt.players.playerFind(testPlayerName1)
		if ndx < 0 {
			t.Errorf("playerFind returns bad ndx:%d", ndx)
		}

		// verify existance in player array
		if strings.Compare(gt.players[ndx].name, testPlayerName1) != 0 {
			t.Error("missing player")
		}

		// delete
		var commands2 commandArrayType
		commands2[0] = "playerDelete"
		commands2[1] = testPlayerName1

		nc2 := newCommand(testPlayerName1, "reqId2", 2, commands2)
		log.Println(nc2)

		// load and run command
		gt.commandStack.push(nc2)
		gt.serviceCommandStack()

		// test for player delete
		ndx = gt.players.playerFind(testPlayerName1)
		if ndx >= 0 {
			t.Errorf("playerFind returns bad ndx:%d", ndx)
		}
	*/
}
