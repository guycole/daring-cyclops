// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

// a summary is produced for for each game
// all games are displayed on the splash page for the user to choose a game

type gameSummaryType struct {
	age       turnCounterType
	blueScore scoreType
	blueShips uint16
	key       *tokenKeyType
	redScore  scoreType
	redShips  uint16
}

func (gt *gameType) newGameSummary() *gameSummaryType {
	gst := gameSummaryType{age: gt.currentTurn, key: gt.key}

	for _, gpat := range gt.playerArray {
		if gpat != nil {
			switch gpat.team {
			case blueTeam:
				gst.blueShips++
				gst.blueScore += gpat.score
			case redTeam:
				gst.redShips++
				gst.redScore += gpat.score
			}
		}
	}

	return &gst
}
