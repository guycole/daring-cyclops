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
	age      uint64
	key      *GameKeyType
	players  playerArrayType
	sugarLog *zap.SugaredLogger
}

func newGame(id string, sugarLog *zap.SugaredLogger) (*gameType, error) {
	result := gameType{key: newGameKey(id), sugarLog: sugarLog}
	return &result, nil
}
