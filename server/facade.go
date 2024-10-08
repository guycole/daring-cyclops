// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"go.uber.org/zap"
)

type facadeType struct {
	featureFlags uint32
	gameManager  *gameManagerType
	sugarLog     *zap.SugaredLogger
}

func newFacade(featureFlags uint32, gameManager *gameManagerType, sugarLog *zap.SugaredLogger) *facadeType {
	gameManager.runAllGames(defaultSleepSeconds)

	return &facadeType{featureFlags: featureFlags, gameManager: gameManager, sugarLog: sugarLog}
}

func (ft *facadeType) addPlayerToGame(gameKey *tokenKeyType, playerKey *tokenKeyType, playerShip shipNameEnum, playerTeam teamEnum) {
	gt := ft.gameManager.findGame(gameKey)
	pt := ft.gameManager.playerManager.findPlayerByKey(playerKey)
	gt.addPlayerToGame(pt, playerShip, playerTeam)
}

func (ft *facadeType) findGame(key *tokenKeyType) *gameType {
	result := ft.gameManager.findGame(key)
	return result
}

func (ft *facadeType) gameSummary() gameSummaryArrayType {
	return ft.gameManager.gameSummary()
}

// add to global player map
func (ft *facadeType) playerAdd(name string) (*playerType, error) {
	return ft.gameManager.playerManager.addFreshPlayer(name)
}

// select from global player map
func (ft *facadeType) playerGet(key *tokenKeyType) *playerType {
	return ft.gameManager.playerManager.playerGet(key)
}
