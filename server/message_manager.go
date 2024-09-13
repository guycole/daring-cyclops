// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"go.uber.org/zap"
)

type messageManagerType struct {
	sugarLog *zap.SugaredLogger
}

// convenience factory
func newMessageManager(playerManager *PlayerManagerType, sugarLog *zap.SugaredLogger) *messageManagerType {
	return nil
	// gmt := GameManagerType{playerManager: playerManager, sugarLog: sugarLog}
	// gmt.gameMaps = make(map[string]*gameType)
	// return &gmt
}
