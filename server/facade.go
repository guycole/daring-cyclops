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

func newFacade(featureFlags uint32, gameManager *GameManagerType, sugarLog *zap.SugaredLogger) (*FacadeType, error) {
	gameManager.runAllGames()

	return &FacadeType{featureFlags: featureFlags, gameManager: gameManager, sugarLog: sugarLog}, nil
}

func (ft *FacadeType) findGame(key *GameKeyType) *gameType {
	result := ft.gameManager.findGame(key)
	return result
}

func (ft *FacadeType) gameSummary() gameSummaryArrayType {
	return ft.gameManager.gameSummary()
}

func (ft *FacadeType) playerIdentityNew(name string) (*playerIdentityType, error) {
	pit := ft.gameManager.addFreshPlayer(name)
	return pit, nil
}

func (ft *FacadeType) playerIdentityGet(key *PlayerKeyType) (*playerIdentityType, error) {
	pit := ft.gameManager.playerIdentityGet(key)
	return pit, nil
}
