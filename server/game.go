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
type playerArrayType [maxPlayers]*playerIdentityType

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
	age          uint64
	blue_score   uint64
	blue_ships   uint16
	key          *GameKeyType
	players      playerArrayType
	blue_players string
	red_players  string
	red_score    uint64
	red_ships    uint16
	removeGame   bool
	sugarLog     *zap.SugaredLogger
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
	age        uint64
	blue_score uint64
	blue_ships uint16
	key        *GameKeyType
	red_score  uint64
	red_ships  uint16
}

func newGameSummary(gt *gameType) *gameSummaryType {
	results := gameSummaryType{age: gt.age, blue_score: gt.blue_score, blue_ships: gt.blue_ships, key: gt.key, red_score: gt.red_score, red_ships: gt.red_ships}
	return &results
}
