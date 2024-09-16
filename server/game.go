// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// player for this game
type gamePlayerType struct {
	active   bool
	joinedAt turnCounterType
	key      *playerKeyType
	name     string
	rank     rankEnum
	score    uint64
	ship     shipNameEnum
	team     teamEnum
}

// convenience factory
func newGamePlayer(key *playerKeyType, name string, rank rankEnum, ship shipNameEnum, team teamEnum, tc turnCounterType) *gamePlayerType {
	gpt := gamePlayerType{active: true, key: key, name: name, rank: rank, score: 0, ship: ship, team: team}
	gpt.joinedAt = tc
	return &gpt
}

const (
	defaultSleepSeconds uint16 = 3

	maxGameTeams       uint16 = 2
	maxGameTeamPlayers uint16 = 5
	maxGamePlayers     uint16 = maxGameTeamPlayers * maxGameTeams
)

// playerArrayType contains all active players
//type gamePlayerArrayType [maxGamePlayers]*gamePlayerType

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

type turnCounterType uint64

type gameType struct {
	activeFlag   bool
	currentTurn  turnCounterType
	inQueue      []*commandType
	key          *gameKeyType
	playerMap    map[string]*gamePlayerType
	removeGame   bool
	shipMap      map[shipNameEnum]*shipType
	sleepSeconds uint16
	sugarLog     *zap.SugaredLogger
}

func newGame(sleepSeconds uint16, sugarLog *zap.SugaredLogger) (*gameType, error) {
	players := make(map[string]*gamePlayerType)
	ships := make(map[shipNameEnum]*shipType)
	gt := gameType{activeFlag: true, key: newGameKey(""), removeGame: false, playerMap: players, shipMap: ships, sleepSeconds: sleepSeconds, sugarLog: sugarLog}
	//gt.inQueue = make(map[string]*commandType)

	if sleepSeconds > 0 {
		sugarLog.Infof("fresh game %s with thread", gt.key.key)
		go gt.eclectic()
	} else {
		sugarLog.Infof("fresh game %s no thread", gt.key.key)
	}

	return &gt, nil
}

func (gt *gameType) eclectic() {
	gt.sleepSeconds = 5

	for {
		gt.sugarLog.Info("eclectic:", gt.currentTurn)

		for len(gt.inQueue) > 0 {
			gt.sugarLog.Info("eclectic dispatch")
			break
			//element := gt.inQueue[0]
			//gt.sugarLog.Info(element)
			//gt.inQueue = gt.inQueue[1:]
		}

		gt.currentTurn++
		time.Sleep(time.Duration(gt.sleepSeconds) * time.Second)
	}
}

func (gt *gameType) enqueue(ct *commandType) {
	gt.inQueue = append(gt.inQueue, ct)
	//gt.inQueue[ct.key.key] = ct
	//gt.inQueue = append(gt.inQueue, ct)
}

type gameSummaryType struct {
	age       turnCounterType
	blueScore uint64
	blueShips uint16
	key       *gameKeyType
	redScore  uint64
	redShips  uint16
}

func newGameSummary(gt *gameType) *gameSummaryType {
	gst := gameSummaryType{age: gt.currentTurn, key: gt.key}

	for _, val := range gt.playerMap {
		if val != nil {
			switch val.team {
			case blueTeam:
				gst.blueShips++
				gst.blueScore += val.score
			case redTeam:
				gst.redShips++
				gst.redScore += val.score
			}
		}
	}

	return &gst
}

func (gt *gameType) findPlayerByKey(key *playerKeyType) *gamePlayerType {
	result := gt.playerMap[key.key]
	return result
}

func (gt *gameType) findPlayerByName(name string) *gamePlayerType {
	for _, val := range gt.playerMap {
		if val.name == name {
			return val
		}
	}

	return nil
}

func (gt *gameType) findShipByName(name shipNameEnum) *shipType {
	/*
		for _, val := range gt.shipMap {
			if val.name == name {
				return val
			}
		}
	*/

	return nil
}

func (gt *gameType) addPlayerToGame(pt *playerType, ship shipNameEnum, team teamEnum) {
	//ensure there are no stale entries
	gpt := gt.findPlayerByKey(pt.key)
	if gpt != nil {
		gt.sugarLog.Info("delete duplicate player key")
		delete(gt.playerMap, pt.key.key)
	}

	st := gt.findShipByName(ship)
	if st != nil {
		gt.sugarLog.Info("delete duplicate ship")
	}

	// convert playerType to gamePlayerType
	gpt = newGamePlayer(pt.key, pt.name, pt.rank, ship, team, gt.currentTurn)
	gt.playerMap[pt.key.key] = gpt

	st, _ = newShip(ship)
	gt.shipMap[ship] = st
}
