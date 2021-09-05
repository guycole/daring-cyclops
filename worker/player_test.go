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
	const testId = "testId"
	const testName = "testName"

	result, err := newPlayer(testName, testId, "cadet", "blue")
	if err != nil {
		t.Errorf("newPlayer error:%s", err)
	}

	if result != nil {
		if result.active != true {
			t.Error("newPlayer active failure")
		}

		if result.name != testName {
			t.Error("newPlayer name failure")
		}

		if result.rank != cadetRank {
			t.Error("newPlayer rank failure")
		}

		if result.team != blueTeam {
			t.Error("newPlayer team failure")
		}

		if result.uuid != testId {
			t.Error("newPlayer id failure")
		}
	} else {
		t.Error("newPlayer returns nil")
	}
}

func TestNewBadPlayer01(t *testing.T) {
	const testId = "testId"

	result, err := newPlayer("", testId, "cadet", "blue")
	if err == nil {
		t.Errorf("newPlayer error:expecting bad player")
	}

	if result != nil {
		t.Error("newPlayer error expecting nil")
	}
}

func TestNewBadPlayer02(t *testing.T) {
	const testName = "testName"

	result, err := newPlayer(testName, "", "cadet", "blue")
	if err == nil {
		t.Errorf("newPlayer error:expecting bad id")
	}

	if result != nil {
		t.Error("newPlayer error expecting nil")
	}
}

func TestNewBadPlayer03(t *testing.T) {
	const testId = "testId"
	const testName = "testName"

	result, err := newPlayer(testName, testId, "", "blue")
	if err == nil {
		t.Errorf("newPlayer error:expecting bad rank")
	}

	if result != nil {
		t.Error("newPlayer error expecting nil")
	}
}

func TestNewBadPlayer04(t *testing.T) {
	const testId = "testId"
	const testName = "testName"

	result, err := newPlayer(testName, testId, "cadet", "")
	if err == nil {
		t.Errorf("newPlayer error:expecting bad team")
	}

	if result != nil {
		t.Error("newPlayer error expecting nil")
	}
}
