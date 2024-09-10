// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"go.uber.org/zap"
)

type PlayerManagerType struct {
	sugarLog  *zap.SugaredLogger
	playerMap map[string]*playerType // all known players
}

// convenience factory
func newPlayerManager(sugarLog *zap.SugaredLogger) *PlayerManagerType {
	pmt := PlayerManagerType{sugarLog: sugarLog}
	pmt.playerMap = make(map[string]*playerType)
	return &pmt
}

func (pmt *PlayerManagerType) findPlayer(key *PlayerKeyType) *playerType {
	result := pmt.playerMap[key.key]
	return result
}

func (pmt *PlayerManagerType) addFreshPlayer(name string) *playerType {
	pt, err := newPlayer(name, "", "")
	if err == nil {
		pmt.playerMap[pt.key.key] = pt
		return pt
	}

	return nil
}

func (pmt *PlayerManagerType) playerGet(key *PlayerKeyType) *playerType {
	result := pmt.playerMap[key.key]
	return result
}

func (pmt *PlayerManagerType) playerUpdate(pt *playerType) {
	pmt.playerMap[pt.key.key] = pt
}
