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
	gameManager.startAllGames()

	return &FacadeType{featureFlags: featureFlags, gameManager: gameManager, sugarLog: sugarLog}, nil
}

func (ft *FacadeType) gameCatalog() gameArrayType {
	return ft.gameManager.GameArray
}

func (ft *FacadeType) gameNew() (*gameType, error) {
	gt, err := newGame("", ft.sugarLog)
	return gt, err
}

func (ft *FacadeType) playerNew(name string) (*playerType, error) {
	_, err := newPlayer(name, "", "", "")
	if err != nil {
		ft.sugarLog.Error("playerNew failure")
		return nil, err
	}

	return nil, nil
}
