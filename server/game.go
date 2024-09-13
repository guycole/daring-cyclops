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
	active bool
	key    *playerKeyType
	name   string
	rank   rankEnum
	score  uint64
	ship   string
	team   teamEnum
}

// convenience factory
func newGamePlayer(key *playerKeyType, name string, rank rankEnum, ship string, team teamEnum) *gamePlayerType {
	result := gamePlayerType{active: true, key: key, name: name, rank: rank, score: 0, ship: ship, team: team}
	return &result
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

type gameType struct {
	age        uint64
	key        *gameKeyType
	playerMap  map[string]*gamePlayerType
	removeGame bool
	sugarLog   *zap.SugaredLogger
}

func newGame(sugarLog *zap.SugaredLogger) (*gameType, error) {
	players := make(map[string]*gamePlayerType)
	result := gameType{key: newGameKey(""), age: 0, removeGame: false, playerMap: players, sugarLog: sugarLog}
	return &result, nil
}

type gameSummaryType struct {
	age       uint64
	blueScore uint64
	blueShips uint16
	key       *gameKeyType
	redScore  uint64
	redShips  uint16
}

func newGameSummary(gt *gameType) *gameSummaryType {
	gst := gameSummaryType{age: gt.age, key: gt.key}

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
