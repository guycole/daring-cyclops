// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"strings"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// player for this game
type gamePlayerType struct {
	active   bool
	joinedAt turnCounterType
	key      *playerKeyType
	inbound  map[string]*commandType
	name     string
	rank     rankEnum
	score    uint64
	ship     shipNameEnum
	team     teamEnum
}

func (gpt *gamePlayerType) addInboundCommand(ct *commandType) {
	// TODO
}

// convenience factory
func newGamePlayer(key *playerKeyType, name string, rank rankEnum, ship shipNameEnum, team teamEnum, tc turnCounterType) *gamePlayerType {
	gpt := gamePlayerType{active: true, key: key, name: name, rank: rank, score: 0, ship: ship, team: team}
	gpt.inbound = make(map[string]*commandType)
	gpt.joinedAt = tc
	return &gpt
}

const (
	maxGameTeams       uint16 = 2
	maxGameTeamPlayers uint16 = 5
	maxGamePlayers     uint16 = maxGameTeamPlayers * maxGameTeams
)

// playerArrayType contains all active players
//type gamePlayerArrayType [maxGamePlayers]*gamePlayerType

type gameKeyType struct {
	key string
}

// convenience factory
func newGameKey(key string) *gameKeyType {
	var result gameKeyType

	temp := strings.TrimSpace(key)
	if len(temp) < 36 {
		result = gameKeyType{key: uuid.NewString()}
	} else {
		result = gameKeyType{key: temp}
	}

	return &result
}

type turnCounterType uint64

type gameType struct {
	currentTurn turnCounterType
	key         *gameKeyType
	playerMap   map[string]*gamePlayerType
	removeGame  bool
	shipMap     map[shipNameEnum]*shipType
	sugarLog    *zap.SugaredLogger
}

func newGame(sugarLog *zap.SugaredLogger) (*gameType, error) {
	players := make(map[string]*gamePlayerType)
	ships := make(map[shipNameEnum]*shipType)
	result := gameType{key: newGameKey(""), removeGame: false, playerMap: players, shipMap: ships, sugarLog: sugarLog}
	return &result, nil
}

type gameSummaryType struct {
	age       turnCounterType
	blueScore uint64
	blueShips uint16
	key       *gameKeyType
	redScore  uint64
	redShips  uint16
}

func newGameSummary(gt *gameType) *gameSummaryType {
	gst := gameSummaryType{age: gt.currentTurn, key: gt.key}

	for _, val := range gt.playerMap {
		if val != nil {
			switch val.team {
			case blueTeam:
				gst.blueShips++
				gst.blueScore += val.score
			case redTeam:
				gst.redShips++
				gst.redScore += val.score
			}
		}
	}

	return &gst
}

func (gt *gameType) findPlayerByKey(key *playerKeyType) *gamePlayerType {
	result := gt.playerMap[key.key]
	return result
}

func (gt *gameType) findPlayerByName(name string) *gamePlayerType {
	for _, val := range gt.playerMap {
		if val.name == name {
			return val
		}
	}

	return nil
}

func (gt *gameType) findShipByName(name shipNameEnum) *shipType {
	/*
		for _, val := range gt.shipMap {
			if val.name == name {
				return val
			}
		}
	*/

	return nil
}

func (gt *gameType) addPlayerToGame(pt *playerType, ship shipNameEnum, team teamEnum) {
	//ensure there are no stale entries
	gpt := gt.findPlayerByKey(pt.key)
	if gpt != nil {
		gt.sugarLog.Info("delete duplicate player key")
		delete(gt.playerMap, pt.key.key)
	}

	st := gt.findShipByName(ship)
	if st != nil {
		gt.sugarLog.Info("delete duplicate ship")
	}

	// convert playerType to gamePlayerType
	gpt = newGamePlayer(pt.key, pt.name, pt.rank, ship, team, gt.currentTurn)
	gt.playerMap[pt.key.key] = gpt

	st, _ = newShip(ship)
	gt.shipMap[ship] = st
}
