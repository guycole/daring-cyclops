// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

type userSummaryType struct {
	active bool
	age    turnCounterType
	name   string
	rank   rankEnum
	score  uint64
	team   teamEnum
}

func (gt *gameType) usersCommand(ct *commandType) []*userSummaryType {
	gt.sugarLog.Debug("usersCommand")

	results := []*userSummaryType{}

	for _, val := range gt.playerMap {
		//gt.sugarLog.Info(val.name)
		temp := userSummaryType{active: val.active, name: val.name, rank: val.rank, score: val.score, team: val.team}
		temp.age = gt.currentTurn - val.joinedAt
		results = append(results, &temp)
	}

	return results
}
