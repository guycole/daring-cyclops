// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"go.uber.org/zap"
)

const (
	maxGames uint16 = 5
)

type gameArrayType [maxGames]*gameType

type GameManagerType struct {
	GameArray gameArrayType
	SugarLog  *zap.SugaredLogger
}

func newGameManager(sugarLog *zap.SugaredLogger) (*GameManagerType, error) {
	result := GameManagerType{SugarLog: sugarLog}

	for ndx := uint16(0); ndx < maxGames; ndx++ {
		result.GameArray[ndx] = nil
	}

	return &result, nil
}

func (gmt *GameManagerType) startAllGames() {
	var err error

	for ndx := uint16(0); ndx < maxGames; ndx++ {
		if gmt.GameArray[ndx] == nil {
			gmt.GameArray[ndx], err = newGame("", gmt.SugarLog)
			if err != nil {
				gmt.SugarLog.Info(err)
			}
		}
	}
}
