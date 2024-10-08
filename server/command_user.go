// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

type userRequestType struct{}

// convenience factory
func newUserRequest(playerKey *tokenKeyType) *commandType {
	ct := commandType{command: userCommand, sourcePlayerKey: playerKey}
	ct.userRequest = &userRequestType{}
	return &ct
}

type userSummaryType struct {
	active bool
	age    turnCounterType
	name   string
	rank   rankEnum
	score  scoreType
	team   teamEnum
}

type userResponseType struct {
	ust []*userSummaryType
}

func (gt *gameType) userCommand(ct *commandType) *commandType {
	gt.sugarLog.Debug("usersCommand")

	results := []*userSummaryType{}

	for _, val := range gt.playerArray {
		if val == nil {
			gt.sugarLog.Debug("skipping nil player")
			continue
		}

		gt.sugarLog.Debug("not nil player")

		//gt.sugarLog.Info(val.name)
		temp := userSummaryType{active: val.activeFlag, name: val.name, rank: val.rank, score: val.score, team: val.team}
		temp.age = gt.currentTurn - val.joinedAt
		results = append(results, &temp)
		gt.sugarLog.Debug(results)
	}

	response := userResponseType{ust: results}
	gt.sugarLog.Debug(response)

	ct.userResponse = &response

	ct.destinationPlayerKeys = append(ct.destinationPlayerKeys, ct.sourcePlayerKey)

	return ct
}
