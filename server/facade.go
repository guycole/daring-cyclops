// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"go.uber.org/zap"
)

type facadeType struct {
	featureFlags uint32
	gameManager  *GameManagerType
	sugarLog     *zap.SugaredLogger
}

func newFacade(featureFlags uint32, gameManager *GameManagerType, sugarLog *zap.SugaredLogger) *facadeType {
	gameManager.runAllGames()

	return &facadeType{featureFlags: featureFlags, gameManager: gameManager, sugarLog: sugarLog}
}

func (ft *facadeType) addPlayerToGame(gameKey *gameKeyType, playerKey *playerKeyType, playerShip string, playerTeam teamEnum) {
	ft.gameManager.addPlayerToGame(gameKey, playerKey, playerShip, playerTeam)
}

func (ft *facadeType) findGame(key *gameKeyType) *gameType {
	result := ft.gameManager.findGame(key)
	return result
}

func (ft *facadeType) gameSummary() gameSummaryArrayType {
	return ft.gameManager.gameSummary()
}

// add to global player map
func (ft *facadeType) playerAdd(name string) *playerType {
	return ft.gameManager.playerManager.addFreshPlayer(name)
}

// select from global player map
func (ft *facadeType) playerGet(key *playerKeyType) *playerType {
	return ft.gameManager.playerManager.playerGet(key)
}
