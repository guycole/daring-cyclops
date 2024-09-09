// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"strings"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	maxTeams       uint16 = 2
	maxTeamPlayers uint16 = 5
	maxPlayers     uint16 = maxTeamPlayers * maxTeams
)

// playerArrayType contains all active players
type playerArrayType [maxPlayers]*playerType

type GameKeyType struct {
	key string
}

// convenience factory
func newGameKey(key string) *GameKeyType {
	var result GameKeyType

	temp := strings.TrimSpace(key)
	if len(temp) < 36 {
		result = GameKeyType{key: uuid.NewString()}
	} else {
		result = GameKeyType{key: temp}
	}

	return &result
}

type gameType struct {
	age         uint64
	blueScore   uint64
	blueShips   uint16
	key         *GameKeyType
	players     playerArrayType
	bluePlayers string
	redPlayers  string
	redScore    uint64
	redShips    uint16
	removeGame  bool
	sugarLog    *zap.SugaredLogger
}

func newGame(sugarLog *zap.SugaredLogger) (*gameType, error) {
	players := playerArrayType{}
	for ndx := uint16(0); ndx < maxPlayers; ndx++ {
		players[ndx] = nil
	}

	result := gameType{key: newGameKey(""), age: 0, removeGame: false, players: players, sugarLog: sugarLog}
	return &result, nil
}

type gameSummaryType struct {
	age       uint64
	blueScore uint64
	blueShips uint16
	key       *GameKeyType
	redScore  uint64
	redShips  uint16
}

func newGameSummary(gt *gameType) *gameSummaryType {
	results := gameSummaryType{age: gt.age, blueScore: gt.blueScore, blueShips: gt.blueShips, key: gt.key, redScore: gt.redScore, redShips: gt.redShips}
	return &results
}
