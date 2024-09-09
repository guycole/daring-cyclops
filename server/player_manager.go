// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"go.uber.org/zap"
)

type PlayerManagerType struct {
	SugarLog   *zap.SugaredLogger
	PlayerMaps map[string]*playerType // all known players
}

// convenience factory
func newPlayerManager(sugarLog *zap.SugaredLogger) *PlayerManagerType {
	pmt := PlayerManagerType{SugarLog: sugarLog}
	pmt.PlayerMaps = make(map[string]*playerType)
	return &pmt
}

func (pmt *PlayerManagerType) findPlayer(key *PlayerKeyType) *playerType {
	result := pmt.PlayerMaps[key.key]
	return result
}

func (pmt *PlayerManagerType) addFreshPlayer(name string) *playerType {
	pt, err := newPlayer(name, "", "")
	if err == nil {
		pmt.PlayerMaps[pt.key.key] = pt
		return pt
	}

	return nil
}

func (pmt *PlayerManagerType) playerGet(key *PlayerKeyType) *playerType {
	result := pmt.PlayerMaps[key.key]
	return result
}

func (pmt *PlayerManagerType) playerUpdate(pt *playerType) {
	pmt.PlayerMaps[pt.key.key] = pt
}
