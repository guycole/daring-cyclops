// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
//
// player support
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
	return [...]string{"Unknown", "Neutral", "Blue", "Red"}[pte]
}

func findPlayerTeam(arg string) playerTeamEnum {
	for ndx := 0; ndx < len(legalPlayerTeams); ndx++ {
		if legalPlayerTeams[ndx] == arg {
			return playerTeamEnum(ndx)
		}
	}

	return playerTeamEnum(unknownTeam)
}

// playerType comment
type playerType struct {
	active bool
	name   string
	rank   playerRankEnum
	team   playerTeamEnum
	uuid   string
}

/*
func newPlayer(name, id string, rank playerRankEnum, team playerTeamEnum) playerType {
	result := playerType{active: true, name: name, rank: rank, team: team, uuid: id}
	return result
}
*/

func newPlayer(name, id string, rank string, team string) (*playerType, error) {
	if len(id) < 1 {
		return nil, errors.New("bad player id")
	}

	if len(name) < 1 {
		return nil, errors.New("bad player name")
	}

	result := playerType{active: true, name: name, uuid: id}
	playerRank := findPlayerRank(rank)
	if playerRank == unknownRank {
		log.Println("boo")
		return nil, errors.New("unknown rank")
	}

	result.rank = playerRank

	playerTeam := findPlayerTeam(team)
	if playerTeam == unknownTeam {
		log.Println("boo2")
		return nil, errors.New("unknown team")
	}

	result.team = playerTeam

	return &result, nil
}

func playerAdd(pt playerType, gt *gameType) int {
	for ndx := 0; ndx < maxPlayers; ndx++ {
		if gt.players[ndx].active == false {
			gt.players[ndx] = pt
			return ndx
		}
	}

	return -1
}

func playerCensus(gt gameType) (int, int) {
	blue := 0
	red := 0

	for ndx := 0; ndx < maxPlayers; ndx++ {
		if gt.players[ndx].active == true {
			switch gt.players[ndx].team {
			case blueTeam:
				blue++
			case redTeam:
				red++
			}
		}
	}

	return blue, red
}

func playerDelete(target string, gt *gameType) int {
	for ndx := 0; ndx < maxPlayers; ndx++ {
		if strings.Compare(gt.players[ndx].uuid, target) == 0 {
			gt.players[ndx].active = false
			return ndx
		}
	}

	return -1
}

func playerDump(gt gameType) {
	log.Println("=-=-=-= playerDump =-=-=-=")

	for ndx := 0; ndx < maxPlayers; ndx++ {
		rank := gt.players[ndx].rank.string()
		team := gt.players[ndx].team.string()

		log.Printf("%d %t %s %s %s %s", ndx, gt.players[ndx].active, gt.players[ndx].name, rank, team, gt.players[ndx].uuid)
	}

	log.Println("=-=-=-= playerDump =-=-=-=")
}

func playerFind(target string, gt *gameType) int {
	for ndx := 0; ndx < maxPlayers; ndx++ {
		if gt.players[ndx].active == true {
			if strings.Compare(gt.players[ndx].uuid, target) == 0 {
				return ndx
			}
		}
	}

	return -1
}

func commandCreatePlayer(command commandType, gt *gameType) {
	log.Println("create player")

	playerDump(*gt)

	// convert structure
	pt := playerType{active: true, name: command.args[1], uuid: command.player}
	pt.rank = findPlayerRank(command.args[2])
	pt.team = findPlayerTeam(command.args[3])

	log.Println(pt)

	blue, red := playerCensus(*gt)
	log.Println(blue)
	log.Println(red)

	// test for duplicate
	duplicate := playerFind(pt.uuid, gt)
	log.Println(duplicate)

	// ensure team size limits
	blue, red = playerCensus(*gt)
	log.Println(blue)
	log.Println(red)

	// add to player array
	ndx := playerAdd(pt, gt)
	log.Println(gt.players[ndx])
	playerDump(*gt)

	duplicate = playerFind(pt.uuid, gt)
	log.Println(duplicate)
}
