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
	gameMaps      map[string]*gameType // all active games
	sugarLog      *zap.SugaredLogger
	playerManager *PlayerManagerType
}

// convenience factory
func newGameManager(playerManager *PlayerManagerType, sugarLog *zap.SugaredLogger) *GameManagerType {
	gmt := GameManagerType{playerManager: playerManager, sugarLog: sugarLog}
	gmt.gameMaps = make(map[string]*gameType)
	return &gmt
}

// ensure there are always maxGames running
func (gmt *GameManagerType) runAllGames() {
	for key, val := range gmt.gameMaps {
		if val.removeGame {
			gmt.sugarLog.Infof("runAllGames: removing %s", key)
			delete(gmt.gameMaps, key)
		}
	}

	for len(gmt.gameMaps) < int(maxGames) {
		gt, err := newGame(gmt.sugarLog)

		if err == nil {
			gmt.sugarLog.Infof("runAllGames: adding %s", gt.key.key)
			gmt.gameMaps[gt.key.key] = gt
		} else {
			gmt.sugarLog.Info(err)
		}
	}
}

func (gmt *GameManagerType) findGame(key *GameKeyType) *gameType {
	result := gmt.gameMaps[key.key]
	return result
}

// supports gRPC message
type gameSummaryArrayType [maxGames]*gameSummaryType

func (gmt *GameManagerType) gameSummary() gameSummaryArrayType {
	var ndx int
	var results gameSummaryArrayType

	for _, val := range gmt.gameMaps {
		results[ndx] = newGameSummary(val)
		ndx++
	}

	return results
}

func (gmt *GameManagerType) addPlayerToGame(gameKey *GameKeyType, playerKey *PlayerKeyType, playerShip string, playerTeam teamEnum) {
	pm := gmt.gameMaps[gameKey.key].playerMap

	//ensure there are no stale entries
	delete(pm, playerKey.key)

	// convert to game player type
	pt := gmt.playerManager.playerGet(playerKey)
	gpt := newGamePlayer(playerKey, pt.name, pt.rank, playerShip, playerTeam)

	pm[playerKey.key] = gpt
}
