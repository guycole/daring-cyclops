// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"go.uber.org/zap"
)

type gameManagerType struct {
	gameMaps      map[string]*gameType // all active games
	maxGames      uint16               // game population limit
	playerManager *PlayerManagerType   // all known players
	sugarLog      *zap.SugaredLogger
}

// convenience factory
func newGameManager(maxGames uint16, sugarLog *zap.SugaredLogger) *gameManagerType {
	gmt := gameManagerType{maxGames: maxGames, playerManager: newPlayerManager(sugarLog), sugarLog: sugarLog}
	gmt.gameMaps = make(map[string]*gameType)
	return &gmt
}

// ensure there are always maxGames running
func (gmt *gameManagerType) runAllGames(sleepSeconds uint16) {
	for key, val := range gmt.gameMaps {
		if val.removeGame {
			gmt.sugarLog.Infof("runAllGames: removing %s", key)
			delete(gmt.gameMaps, key)
		}
	}

	for len(gmt.gameMaps) < int(gmt.maxGames) {
		gt, err := newGame(sleepSeconds, gmt.sugarLog)

		if err == nil {
			gmt.sugarLog.Infof("runAllGames: adding %s", gt.key.key)
			gmt.gameMaps[gt.key.key] = gt
		} else {
			gmt.sugarLog.Info(err)
		}
	}
}

func (gmt *gameManagerType) findGame(key *gameKeyType) *gameType {
	result := gmt.gameMaps[key.key]
	return result
}

func (gmt *gameManagerType) pickGame() *gameType {
	var gt *gameType

	for _, val := range gmt.gameMaps {
		gt = val
	}

	return gt
}

// supports gRPC message
type gameSummaryArrayType []*gameSummaryType

func (gmt *gameManagerType) gameSummary() gameSummaryArrayType {
	var results gameSummaryArrayType

	for _, val := range gmt.gameMaps {
		results = append(results, newGameSummary(val))
	}

	return results
}
