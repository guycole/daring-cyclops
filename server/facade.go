// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"go.uber.org/zap"
)

type FacadeType struct {
	featureFlags uint32
	gameManager  *GameManagerType
	sugarLog     *zap.SugaredLogger
}

func newFacade(featureFlags uint32, gameManager *GameManagerType, sugarLog *zap.SugaredLogger) *FacadeType {
	gameManager.runAllGames()

	return &FacadeType{featureFlags: featureFlags, gameManager: gameManager, sugarLog: sugarLog}
}

func (ft *FacadeType) addPlayerToGame(gameKey *GameKeyType, playerKey *PlayerKeyType, playerShip string, playerTeam teamEnum) {
	ft.gameManager.addPlayerToGame(gameKey, playerKey, playerShip, playerTeam)
}

func (ft *FacadeType) findGame(key *GameKeyType) *gameType {
	result := ft.gameManager.findGame(key)
	return result
}

func (ft *FacadeType) gameSummary() gameSummaryArrayType {
	return ft.gameManager.gameSummary()
}

// add to global player map
func (ft *FacadeType) playerAdd(name string) *playerType {
	return ft.gameManager.playerManager.addFreshPlayer(name)
}

// select from global player map
func (ft *FacadeType) playerGet(key *PlayerKeyType) *playerType {
	return ft.gameManager.playerManager.playerGet(key)
}
