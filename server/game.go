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
	queueOut []*commandType
	rank     rankEnum
	score    scoreType
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
type gamePlayerArrayType [maxGamePlayers]*gamePlayerType

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

type scoreType uint64
type turnCounterType uint64

type gameType struct {
	acheronFlag   bool                // true means acheron active
	activeFlag    bool                // false means game is over
	currentTurn   turnCounterType     // monotonic turn counter
	queueIn       commandArrayType    // input commands sorted by turn counter
	key           *gameKeyType        // unique game identifier
	playerArray   gamePlayerArrayType // current players and ships in game
	scheduleArray scheduleArrayType   // commands sorted by turn counter
	sleepSeconds  uint16              // delay between turns
	sugarLog      *zap.SugaredLogger  // logger
}

func newGame(sleepSeconds uint16, sugarLog *zap.SugaredLogger) (*gameType, error) {
	gt := gameType{acheronFlag: true, activeFlag: true, key: newGameKey(""), sleepSeconds: sleepSeconds, sugarLog: sugarLog}

	for ndx, _ := range gt.playerArray {
		gt.playerArray[ndx] = nil
	}

	for ndx, _ := range gt.scheduleArray {
		gt.scheduleArray[ndx] = nil
	}

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

		// schedule fresh commands
		for len(gt.queueIn) > 0 {
			candidate := gt.queueIn[0]
			gt.schedule(candidate)
			gt.queueIn = gt.queueIn[1:]
		}

		// consume commands for this turn

		gt.currentTurn++
		time.Sleep(time.Duration(gt.sleepSeconds) * time.Second)
	}
}

func (gt *gameType) enqueue(ct *commandType) {
	gt.queueIn = append(gt.queueIn, ct)
	//gt.inQueue[ct.key.key] = ct
	//gt.inQueue = append(gt.inQueue, ct)
}

func (gt *gameType) findPlayerByKey(key *playerKeyType) *gamePlayerType {
	//	result := gt.playerMap[key.key]
	//	return result
	return nil
}

func (gt *gameType) findPlayerByName(name string) *gamePlayerType {
	/*
		for _, val := range gt.playerMap {
			if val.name == name {
				return val
			}
		}
	*/

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
	// ensure there are no stale entries

	var blue_population, red_population uint16

	for ndx, val := range gt.playerArray {
		// no duplicate players
		if val != nil && val.key.key == pt.key.key {
			gt.sugarLog.Info("duplicate player key")
		}

		// no duplicate ships
		if val != nil && val.ship == ship {
			gt.sugarLog.Info("duplicate ship")
			gt.playerArray[ndx] = nil
		}

		// count team population
		if val != nil && val.team == redTeam {
			red_population++
		} else {
			blue_population++
		}
	}

	// enforce team size limits
	if team == blueTeam {
		if blue_population >= maxGameTeamPlayers {
			gt.sugarLog.Info("blue team full")
		}
	} else {
		if red_population >= maxGameTeamPlayers {
			gt.sugarLog.Info("red team full")
		}
	}

	for ndx, val := range gt.playerArray {
		if val == nil {
			gt.playerArray[ndx] = newGamePlayer(pt.key, pt.name, pt.rank, ship, team, gt.currentTurn)
			break
		}
	}
}

func (gt *gameType) getOutput(pkt *playerKeyType) commandArrayType {
	return nil
}
