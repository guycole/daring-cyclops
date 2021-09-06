// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found
// in the LICENSE file.
package main

import (
	"errors"
	"log"
	"strings"
)

// playerRankEnum
type playerRankEnum int

// must match order for legalPlayerRanks
const (
	unknownRank playerRankEnum = iota
	cadetRank
	lieutenantRank
	captainRank
	admiralRank
)

// must match order for playerRankEnum
var legalPlayerRanks = [...]string{
	"unknown",
	"cadet",
	"lieutenant",
	"captain",
	"admiral",
}

// must match order for playerRankEnum
func (pre playerRankEnum) string() string {
	return [...]string{"unknown", "cadet", "lieutenant", "captain", "admiral"}[pre]
}

func findPlayerRank(arg string) playerRankEnum {
	for ndx := 0; ndx < len(legalPlayerRanks); ndx++ {
		if legalPlayerRanks[ndx] == arg {
			return playerRankEnum(ndx)
		}
	}

	return playerRankEnum(unknownRank)
}

// playerTeamEnum
type playerTeamEnum int

// must match order for legalPlayerTeams
const (
	unknownTeam playerTeamEnum = iota
	neutralTeam
	blueTeam
	redTeam
)

// must match order for playerTeamEnum
var legalPlayerTeams = [...]string{
	"unknown",
	"neutral",
	"blue",
	"red",
}

// must match order for playerTeamEnum
func (pte playerTeamEnum) string() string {
	return [...]string{"unknown", "neutral", "blue", "red"}[pte]
}

func findPlayerTeam(arg string) playerTeamEnum {
	for ndx := 0; ndx < len(legalPlayerTeams); ndx++ {
		if legalPlayerTeams[ndx] == arg {
			return playerTeamEnum(ndx)
		}
	}

	return playerTeamEnum(unknownTeam)
}

type playerType struct {
	name string
	rank playerRankEnum
	team playerTeamEnum
	uuid string
}

const maxTeamPlayers = 5
const maxPlayers = maxTeamPlayers * 2

// playerArrayType contains all active players
type playerArrayType [maxPlayers]*playerType

// newPlayer convenience function to populate struct
func newPlayer(name, id string, rank string, team string) (*playerType, error) {
	if len(id) < 1 {
		return nil, errors.New("bad player id")
	}

	if len(name) < 1 {
		return nil, errors.New("bad player name")
	}

	result := playerType{name: name, uuid: id}
	playerRank := findPlayerRank(rank)
	if playerRank == unknownRank {
		return nil, errors.New("unknown rank")
	}

	result.rank = playerRank

	playerTeam := findPlayerTeam(team)
	if playerTeam == unknownTeam {
		return nil, errors.New("unknown team")
	}

	result.team = playerTeam

	return &result, nil
}

const testPlayerID1 = "testId1"
const testPlayerName1 = "testName1"

// testPlayer1 returns test player1
func testPlayer1() *playerType {
	np1, _ := newPlayer(testPlayerName1, testPlayerID1, "cadet", "blue")
	return np1
}

const testPlayerID2 = "testId2"
const testPlayerName2 = "testName2"

// testPlayer2 returns test player2
func testPlayer2() *playerType {
	np2, _ := newPlayer(testPlayerName2, testPlayerID2, "admiral", "red")
	return np2
}

// playerAdd adds player to array
func playerAdd(pt *playerType, pat *playerArrayType) int {
	log.Printf("playerAdd:%s %s", pt.name, pt.uuid)

	for ndx := 0; ndx < maxPlayers; ndx++ {
		if pat[ndx] == nil {
			pat[ndx] = pt
			return ndx
		}
	}

	return -1
}

// playerCensus returns population of red/blue players
func playerCensus(pat playerArrayType) (int, int) {
	bluePopulation := 0
	redPopulation := 0

	for ndx := 0; ndx < maxPlayers; ndx++ {
		if pat[ndx] != nil {
			switch pat[ndx].team {
			case blueTeam:
				bluePopulation++
			case redTeam:
				redPopulation++
			}
		}
	}

	return bluePopulation, redPopulation
}

// playerDelete removes player from array
func playerDelete(target string, pat *playerArrayType) int {
	log.Printf("playerDelete:%s", target)

	for ndx := 0; ndx < maxPlayers; ndx++ {
		if pat[ndx] != nil {
			if strings.Compare(pat[ndx].uuid, target) == 0 {
				pat[ndx] = nil
				return ndx
			}
		}
	}

	return -1
}

// playerDump diagnostic
func playerDump(pat playerArrayType) {
	log.Println("=-=-=-= playerDump =-=-=-=")

	for ndx := 0; ndx < maxPlayers; ndx++ {
		if pat[ndx] == nil {
			log.Printf("%d nil", ndx)
		} else {
			rank := pat[ndx].rank.string()
			team := pat[ndx].team.string()

			log.Printf("%d %s %s %s %s", ndx, pat[ndx].name, rank, team, pat[ndx].uuid)
		}
	}

	log.Println("=-=-=-= playerDump =-=-=-=")
}

// playerFind returns array index for player by uuid
func playerFind(target string, pat playerArrayType) int {
	for ndx := 0; ndx < maxPlayers; ndx++ {
		if pat[ndx] != nil {
			if strings.Compare(pat[ndx].uuid, target) == 0 {
				return ndx
			}
		}
	}

	return -1
}

// commandPlayerCreate services command
func commandPlayerCreate(ct commandType, gt *gameType) error {
	duplicate := playerFind(ct.player, gt.players)
	if duplicate >= 0 {
		return errors.New("duplicate player id")
	}

	// TODO test for max players per side

	np, err := newPlayer(ct.args[1], ct.player, ct.args[2], ct.args[3])
	if err != nil {
		return errors.New("newPlayer creation failure")
	}

	if np == nil {
		return errors.New("newPlayer returns nil")
	}

	playerAdd(np, &gt.players)

	return nil
}

// commandPlayerDelete services command
func commandPlayerDelete(ct commandType, gt *gameType) {
	// TODO delete ship
	playerDelete(ct.player, &gt.players)
}
