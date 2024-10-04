// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"errors"

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

func (pmt *PlayerManagerType) findPlayerByKey(key *tokenKeyType) *playerType {
	result := pmt.playerMap[key.key]
	return result
}

func (pmt *PlayerManagerType) findPlayerByName(name string) *playerType {
	for _, val := range pmt.playerMap {
		if val.name == name {
			return val
		}
	}

	return nil
}

func (pmt *PlayerManagerType) addFreshPlayer(name string) (*playerType, error) {
	temp := pmt.findPlayerByName(name)
	if temp != nil {
		return nil, errors.New("duplicate player name")
	}

	pt, err := newPlayer(name, "", "")
	if err != nil {
		return nil, err
	}

	pmt.playerMap[pt.key.key] = pt
	return pt, nil
}

func (pmt *PlayerManagerType) playerGet(key *tokenKeyType) *playerType {
	result := pmt.playerMap[key.key]
	return result
}

func (pmt *PlayerManagerType) playerUpdate(pt *playerType) {
	pmt.playerMap[pt.key.key] = pt
}

func (pmt *PlayerManagerType) storeMessage(message *messageType) {
	destination := pmt.playerGet(message.destination)
	message.next = destination.messages
	destination.messages = message
}

func (pmt *PlayerManagerType) retrieveMessage(key *tokenKeyType) *messageType {
	player := pmt.playerGet(key)
	result := player.messages
	player.messages = nil
	return result
}

func (pmt *PlayerManagerType) seedTestUsers() {
	pt1, _ := newPlayer("testPlayer1", "lieutenant", testPlayer1)
	pmt.playerMap[pt1.key.key] = pt1
	pt2, _ := newPlayer("testPlayer2", "captain", testPlayer2)
	pmt.playerMap[pt2.key.key] = pt2
}
