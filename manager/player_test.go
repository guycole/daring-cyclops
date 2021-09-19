// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found
// in the LICENSE file.
package main

import (
	"log"
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
	result, err := newPlayer(testPlayerEmail1, testPlayerName1, "cadet", "blue")
	if err != nil {
		t.Errorf("newPlayer error:%s", err)
	}

	if result != nil {
		if result.Email != testPlayerEmail1 {
			t.Error("newPlayer email failure")
		}

		if result.Name != testPlayerName1 {
			t.Error("newPlayer name failure")
		}

		if result.Rank != cadetRank {
			t.Error("newPlayer rank failure")
		}

		if result.Team != blueTeam {
			t.Error("newPlayer team failure")
		}
	} else {
		t.Error("newPlayer returns nil")
	}
}

func TestNewBadPlayer01(t *testing.T) {
	result, err := newPlayer("", testPlayerName1, "cadet", "blue")

	if err == nil {
		t.Error("newPlayer error:expecting bad player")
	}

	if result != nil {
		t.Error("newPlayer error expecting nil")
	}
}

func TestNewBadPlayer03(t *testing.T) {
	result, err := newPlayer(testPlayerEmail1, "", "cadet", "blue")

	if err == nil {
		t.Error("newPlayer error:expecting bad id")
	}

	if result != nil {
		t.Error("newPlayer error expecting nil")
	}
}

func TestNewBadPlayer04(t *testing.T) {
	result, err := newPlayer(testPlayerEmail1, testPlayerName1, "", "blue")

	if err == nil {
		t.Error("newPlayer error:expecting bad rank")
	}

	if result != nil {
		t.Error("newPlayer error expecting nil")
	}
}

func TestNewBadPlayer05(t *testing.T) {
	result, err := newPlayer(testPlayerEmail1, testPlayerName1, "cadet", "")

	if err == nil {
		t.Error("newPlayer error:expecting bad team")
	}

	if result != nil {
		t.Error("newPlayer error expecting nil")
	}
}

func TestNewRank(t *testing.T) {
	tp1 := testPlayer1()
	rankChange(tp1, captainRank)
	if tp1.Rank != captainRank {
		t.Error("bad rank")
	}
}

func TestNewTeam(t *testing.T) {
	tp1 := testPlayer1()
	teamChange(tp1, blueTeam)
	if tp1.Team != blueTeam {
		t.Error("bad team")
	}
}

func TestRedis01(t *testing.T) {
	gmt := newManager()
	log.Println(gmt)

	tp1 := testPlayer1()
	setPlayer(gmt.rdb, tp1)

	tp2 := testPlayer2()
	setPlayer(gmt.rdb, tp2)

	selected1 := getPlayer(gmt.rdb, testPlayerName1)
	log.Println(selected1)

	if selected1.Name != testPlayerName1 {
		t.Error("selected1 player name failure")
	}

	selected2 := getPlayer(gmt.rdb, testPlayerName2)
	log.Println(selected2)

	if selected2.Name != testPlayerName2 {
		t.Error("selected2 player name failure")
	}
}
