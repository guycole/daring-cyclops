// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"errors"
	"strings"

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

type PlayerKeyType struct {
	key string
}

// convenience factory
func newPlayerKey(key string) *PlayerKeyType {
	var result PlayerKeyType

	temp := strings.TrimSpace(key)
	if len(temp) < 36 {
		result = PlayerKeyType{key: uuid.NewString()}
	} else {
		result = PlayerKeyType{key: temp}
	}

	return &result
}

type playerIdentityType struct {
	key    *PlayerKeyType
	name   string
	points uint64 // lifetime total
	rank   rankEnum
}

// convenience factory
func newPlayerIdentity(name string, rank string, uuid string) (*playerIdentityType, error) {
	result := playerIdentityType{key: newPlayerKey(uuid), points: 0}

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
