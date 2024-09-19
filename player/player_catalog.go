// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package player

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
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
