// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

// rankEnum
type rankEnum int

// must match order for legalRanks
const (
	unknownRank rankEnum = iota
	cadetRank
	lieutenantRank
	captainRank
	admiralRank
)

// must match order for rankEnum
var legalRanks = [...]string{
	"unknown",
	"cadet",
	"lieutenant",
	"captain",
	"admiral",
}

// must match order for rankEnum
func (re rankEnum) string() string {
	return [...]string{"unknown", "cadet", "lieutenant", "captain", "admiral"}[re]
}

func findRank(arg string) rankEnum {
	for ndx := 0; ndx < len(legalRanks); ndx++ {
		if strings.Compare(legalRanks[ndx], arg) == 0 {
			return rankEnum(ndx)
		}
	}

	return rankEnum(unknownRank)
}

// teamEnum
type teamEnum int

// must match order for legalTeams
const (
	unknownTeam teamEnum = iota
	neutralTeam
	blueTeam
	redTeam
	acheronTeam
)

// must match order for teamEnum
var legalTeams = [...]string{
	"unknown",
	"neutral",
	"blue",
	"red",
	"acheron",
}

// must match order for teamEnum
func (te teamEnum) string() string {
	return [...]string{"unknown", "neutral", "blue", "red", "acheron"}[te]
}

func findTeam(arg string) teamEnum {
	for ndx := 0; ndx < len(legalTeams); ndx++ {
		if strings.Compare(legalTeams[ndx], arg) == 0 {
			return teamEnum(ndx)
		}
	}

	return teamEnum(unknownTeam)
}

const (
	testPlayer1 = "87277d2e-86f8-4dbd-8b7c-4ae3bdbad703"
	testPlayer2 = "daa10cdd-14ce-4996-9518-370b692a059f"
)

type playerKeyType struct {
	key string
}

// convenience factory
func newPlayerKey(key string) *playerKeyType {
	var result playerKeyType

	temp := strings.TrimSpace(key)
	if len(temp) < 36 {
		result = playerKeyType{key: uuid.NewString()}
	} else {
		result = playerKeyType{key: temp}
	}

	return &result
}

type playerType struct {
	key              *playerKeyType
	lastOn           time.Time
	name             string
	cumulativePoints uint64    // lifetime total
	highPoints       uint64    // single game high
	highPointsTime   time.Time // time of highPoints
	messages         *messageType
	sortie           uint64
	rank             rankEnum
}

// convenience factory
func newPlayer(name string, rank string, uuid string) (*playerType, error) {
	result := playerType{key: newPlayerKey(uuid)}

	temp := strings.TrimSpace(name)
	if len(temp) == 0 {
		return nil, errors.New("name failure")
	} else {
		result.name = temp
	}

	temp = strings.TrimSpace(rank)
	if len(temp) == 0 {
		result.rank = cadetRank
	} else {
		result.rank = findRank(temp)
	}

	return &result, nil
}
