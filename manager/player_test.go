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
		if result.email != testPlayerEmail1 {
			t.Error("newPlayer email failure")
		}

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

func TestRedis01(t *testing.T) {
	gmt := newManager()
	log.Println(gmt)

	/*
		setPlayer(gmt.rdb)
		//log.Println(xx)
	*/

	tp1 := testPlayer1()
	setPlayer(gmt.rdb, tp1)

	xxxx := getPlayer(gmt.rdb, testPlayerName1)
	log.Println(xxxx)

	/*
		key := testPlayerName1
		log.Println(key)

		p := rdb.Get(context.Background(), key)
		log.Println(p)
	*/
}
