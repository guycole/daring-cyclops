// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

type scoreType uint64

// player for this game
type gamePlayerType struct {
	active    bool            // false means player should be removed
	joinedAt  turnCounterType // turn counter when player joined
	key       *playerKeyType  // unique player identifier
	maxFuture turnCounterType // future command turn
	name      string          // player name
	queueOut  []*commandType  // completed commands awaiting output
	rank      rankEnum        // player rank
	score     scoreType       // player score
	ship      shipNameEnum    // player ship
	team      teamEnum        // player team
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

func newGamePlayerArray() gamePlayerArrayType {
	var gpat gamePlayerArrayType

	for ndx, _ := range gpat {
		gpat[ndx] = nil
	}

	return gpat
}

func (gt *gameType) gamePlayerArrayDumper() {
	gt.sugarLog.Debug("====> player array dump <====")

	for ndx, gpt := range gt.playerArray {
		if gpt == nil {
			gt.sugarLog.Debugf("%d nil", ndx)
		} else {
			gt.sugarLog.Debugf("%d %s %s", ndx, gpt.name, gpt.team.string())
		}
	}

	gt.sugarLog.Debug("====> player array dump <====")
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
		if val != nil {
			if val.team == redTeam {
				red_population++
			} else {
				blue_population++
			}
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

	// add to catalog
}

func (gt *gameType) getOutput(pkt *playerKeyType) commandArrayType {
	return nil
}
