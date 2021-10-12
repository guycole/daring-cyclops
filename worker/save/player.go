// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"errors"
	"log"
	"strings"
)

// rankEnum
type rankEnum int

// must match order for legalRanks
const (
	unknownRank rankEnum = iota
	cadetRank
	lieutenantRank
	captainRank
	admiralRank
)

// must match order for rankEnum
var legalRanks = [...]string{
	"unknown",
	"cadet",
	"lieutenant",
	"captain",
	"admiral",
}

// must match order for rankEnum
func (re rankEnum) string() string {
	return [...]string{"unknown", "cadet", "lieutenant", "captain", "admiral"}[re]
}

func findRank(arg string) rankEnum {
	for ndx := 0; ndx < len(legalRanks); ndx++ {
		if strings.Compare(legalRanks[ndx], arg) == 0 {
			return rankEnum(ndx)
		}
	}

	return rankEnum(unknownRank)
}

// teamEnum
type teamEnum int

// must match order for legalTeams
const (
	unknownTeam teamEnum = iota
	neutralTeam
	blueTeam
	redTeam
	acheronTeam
)

// must match order for teamEnum
var legalTeams = [...]string{
	"unknown",
	"neutral",
	"blue",
	"red",
	"acheron",
}

// must match order for teamEnum
func (te teamEnum) string() string {
	return [...]string{"unknown", "neutral", "blue", "red", "acheron"}[te]
}

func findTeam(arg string) teamEnum {
	for ndx := 0; ndx < len(legalTeams); ndx++ {
		if strings.Compare(legalTeams[ndx], arg) == 0 {
			return teamEnum(ndx)
		}
	}

	return teamEnum(unknownTeam)
}

type playerType struct {
	name string
	rank rankEnum
	team teamEnum

	turnQueueNdx int
	turnQueue    turnQueueArrayType
}

const maxTeamPlayers = 5
const maxPlayers = maxTeamPlayers * 2

// playerArrayType contains all active players
type playerArrayType [maxPlayers]*playerType

// newPlayer convenience function to populate struct
func newPlayer(name string, rank string, team string) (*playerType, error) {
	if len(name) < 1 {
		return nil, errors.New("bad player name")
	}

	result := playerType{name: name}
	playerRank := findRank(rank)
	if playerRank == unknownRank {
		return nil, errors.New("unknown rank")
	}

	result.rank = playerRank

	playerTeam := findTeam(team)
	if playerTeam == unknownTeam {
		return nil, errors.New("unknown team")
	}

	result.team = playerTeam

	for ndx := 0; ndx < maxTurnQueueArray; ndx++ {
		result.turnQueue[ndx] = newTurnQueue()
	}

	return &result, nil
}

func testPlayer1() *playerType {
	np1, _ := newPlayer(testPlayerName1, "cadet", "blue")
	return np1
}

func testPlayer2() *playerType {
	np2, _ := newPlayer(testPlayerName2, "admiral", "red")
	return np2
}

func (pat *playerArrayType) add(pt *playerType) int {
	log.Printf("playerAdd:%s", pt.name)

	for ndx := 0; ndx < maxPlayers; ndx++ {
		if pat[ndx] == nil {
			pat[ndx] = pt
			return ndx
		}
	}

	return -1
}

func (pat playerArrayType) census() (int, int) {
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

func (pat *playerArrayType) delete(target string) int {
	log.Printf("playerDelete:%s", target)

	for ndx := 0; ndx < maxPlayers; ndx++ {
		if pat[ndx] != nil {
			if strings.Compare(pat[ndx].name, target) == 0 {
				pat[ndx] = nil
				return ndx
			}
		}
	}

	return -1
}

func (pat playerArrayType) dump() {
	log.Println("=-=-=-= playerDump =-=-=-=")

	for ndx := 0; ndx < maxPlayers; ndx++ {
		if pat[ndx] == nil {
			log.Printf("%d nil", ndx)
		} else {
			rank := pat[ndx].rank.string()
			team := pat[ndx].team.string()

			log.Printf("%d %s %s %s", ndx, pat[ndx].name, rank, team)
		}
	}

	log.Println("=-=-=-= playerDump =-=-=-=")
}

func (pat playerArrayType) find(target string) int {
	for ndx := 0; ndx < maxPlayers; ndx++ {
		if pat[ndx] != nil {
			if strings.Compare(pat[ndx].name, target) == 0 {
				return ndx
			}
		}
	}

	return -1
}

func commandPlayerCreate(tnt *turnNodeType, pat *playerArrayType) {
	log.Println("commandPlayerCreate")

	rawRank := tnt.commands[1]
	rawTeam := tnt.commands[2]

	np, err := newPlayer(tnt.name, rawRank, rawTeam)
	if err != nil {
		log.Println("skipping commandPlayerCreate w/newPlayer error")
		return
	}

	duplicate := pat.find(tnt.name)
	if duplicate >= 0 {
		log.Println("skipping commandPlayerCreate w/duplicate player")
		return
	}

	bluePopulation, redPopulation := pat.census()

	switch np.team {
	case blueTeam:
		if bluePopulation >= maxTeamPlayers {
			log.Println("skipping commandPlayerCreate w/max blue team")
			return
		}
	case redTeam:
		if redPopulation >= maxTeamPlayers {
			log.Println("skipping commandPlayerCreate w/max red team")
			return
		}
	default:
		log.Println("skipping commandPlayerCreate w/unknown team")
		return
	}

	pat.add(np)
}

func commandPlayerDelete(tnt *turnNodeType, bat *boardArrayType, sat *shipArrayType, pat *playerArrayType) {
	log.Println("commandPlayerDelete")

	ndx := sat.findByOwner(tnt.name)
	if ndx >= 0 {
		sat.delete(sat[ndx].uuid, bat)
	}

	pat.delete(tnt.name)
}
