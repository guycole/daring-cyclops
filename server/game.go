// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"strings"
	"time"

	"github.com/google/uuid"

	"go.uber.org/zap"
)

type turnCounterType uint64

type gameKeyType struct {
	key string
}

// convenience factory
func newGameKey(key string) *gameKeyType {
	var result gameKeyType

	temp := strings.TrimSpace(key)
	if len(temp) < 36 {
		result = gameKeyType{key: uuid.NewString()}
	} else {
		result = gameKeyType{key: temp}
	}

	return &result
}

type gameType struct {
	acheronFlag   bool                    // true means acheron active
	activeFlag    bool                    // false means game is over
	catalogMap    map[string]*catalogType // all objects on gameBoard
	currentTurn   turnCounterType         // monotonic turn counter
	gameBoard     *boardArrayType         // game board
	key           *gameKeyType            // unique game identifier
	playerArray   gamePlayerArrayType     // current players and ships in game
	scheduleArray scheduleArrayType       // commands sorted by turn counter
	sleepSeconds  uint16                  // delay between turns
	sugarLog      *zap.SugaredLogger      // logger
}

func newGame(sleepSeconds uint16, sugarLog *zap.SugaredLogger) (*gameType, error) {
	gt := gameType{acheronFlag: true, activeFlag: true, key: newGameKey(""), sleepSeconds: sleepSeconds, sugarLog: sugarLog}

	gt.catalogMap = make(map[string]*catalogType)
	gt.gameBoard = newGameBoard(emptyBoard)
	gt.playerArray = newGamePlayerArray()
	gt.scheduleArray = newScheduleArray()

	if sleepSeconds > 0 {
		sugarLog.Infof("fresh game %s with thread", gt.key.key)
		go gt.eclectic()
	} else {
		sugarLog.Infof("fresh game %s no thread", gt.key.key)
	}

	return &gt, nil
}

func (gt *gameType) eclectic() {
	for gt.activeFlag {
		gt.playTurn()
		time.Sleep(time.Duration(gt.sleepSeconds) * time.Second)
	}
}

func (gt *gameType) playTurn() {
	ndx := gt.currentTurn % turnCounterType(maxScheduleArray)
	gt.sugarLog.Debugf("playTurn %d %d:", gt.currentTurn, ndx)

	st := gt.scheduleArray[ndx]
	gt.currentTurn++

	gt.sugarLog.Debugf("len:%d", st.length)

	for current := st.root; current != nil; current = current.next {
		gt.commandDispatch(current)
	}

	st.length = 0
	st.root = nil
	st.tail = nil
}

func (gt *gameType) findPlayerByKey(key *playerKeyType) *gamePlayerType {
	for _, gpt := range gt.playerArray {
		if strings.Compare(gpt.key.key, key.key) == 0 {
			return gpt
		}
	}

	return nil
}

func (gt *gameType) findPlayerByName(name string) *gamePlayerType {
	for _, gpt := range gt.playerArray {
		if strings.Compare(gpt.name, name) == 0 {
			return gpt
		}
	}

	return nil
}

func (gt *gameType) findPlayerByShip(name shipNameEnum) *gamePlayerType {
	for _, gpt := range gt.playerArray {
		if gpt.ship.name == name {
			return gpt
		}
	}

	return nil
}

func (gt *gameType) enqueue(ct *commandType) {
	gpt := gt.findPlayerByKey(ct.sourcePlayerKey)
	if gpt == nil {
		gt.sugarLog.Info("rejecting message from unknown player")
		return
	}

	if gpt.maxFuture < gt.currentTurn {
		gpt.maxFuture = gt.currentTurn
	}

	ndx := gpt.maxFuture % turnCounterType(maxScheduleArray)
	gt.sugarLog.Infof("schedule command at %d %s %s %s", ndx, legalGameCommands[ct.command].longName, gpt.name, gpt.team.string())

	st := gt.scheduleArray[ndx]
	st.scheduleAdd(ct)

	gpt.maxFuture = gpt.maxFuture + legalGameCommands[ct.command].duration
}
