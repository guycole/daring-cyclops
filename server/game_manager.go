// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"go.uber.org/zap"
)

const (
	maxGames uint16 = 5
)

type GameManagerType struct {
	GameMaps      map[string]*gameType // all active games
	SugarLog      *zap.SugaredLogger
	PlayerManager *PlayerManagerType
}

// convenience factory
func newGameManager(playerManager *PlayerManagerType, sugarLog *zap.SugaredLogger) *GameManagerType {
	gmt := GameManagerType{PlayerManager: playerManager, SugarLog: sugarLog}
	gmt.GameMaps = make(map[string]*gameType)
	return &gmt
}

// ensure there are always maxGames running
func (gmt *GameManagerType) runAllGames() {
	for key, val := range gmt.GameMaps {
		if val.removeGame {
			gmt.SugarLog.Infof("runAllGames: removing %s", key)
			delete(gmt.GameMaps, key)
		}
	}

	for len(gmt.GameMaps) < int(maxGames) {
		gt, err := newGame(gmt.SugarLog)

		if err == nil {
			gmt.SugarLog.Infof("runAllGames: adding %s", gt.key.key)
			gmt.GameMaps[gt.key.key] = gt
		} else {
			gmt.SugarLog.Info(err)
		}
	}
}

func (gmt *GameManagerType) findGame(key *GameKeyType) *gameType {
	result := gmt.GameMaps[key.key]
	return result
}

// supports gRPC message
type gameSummaryArrayType [maxGames]*gameSummaryType

func (gmt *GameManagerType) gameSummary() gameSummaryArrayType {
	var ndx int
	var results gameSummaryArrayType

	for _, val := range gmt.GameMaps {
		results[ndx] = newGameSummary(val)
		ndx++
	}

	return results
}

func (gmt *GameManagerType) addPlayerToGame(gameKey *GameKeyType, playerKey *PlayerKeyType, playerTeam teamEnum) {
	// switch array to map
	/*
		game := gmt.GameMaps[gameKey.key]

		switch playerTeam {
		case blueTeam:
			game.blue_players = playerKey.key
		case redTeam:
			game.red_players = playerKey.key
		default:
			gmt.SugarLog.Infof("addPlayerToGame: unknown team %s", playerTeam.string())
		}
	*/
}
