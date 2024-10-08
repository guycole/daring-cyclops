// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

type scoreType uint64

// player for this game
type gamePlayerType struct {
	activeFlag bool            // false means player should be removed
	joinedAt   turnCounterType // turn counter when player joined
	key        *tokenKeyType   // unique player identifier
	maxFuture  turnCounterType // future command turn
	name       string          // player name
	queueOut   []*commandType  // completed commands awaiting output
	rank       rankEnum        // player rank
	score      scoreType       // player score
	ship       *shipType       // player ship
	team       teamEnum        // player team
}

// convenience factory
func newGamePlayer(key *tokenKeyType, name string, rank rankEnum, ship shipNameEnum, team teamEnum, tc turnCounterType) *gamePlayerType {
	gpt := gamePlayerType{activeFlag: true, key: key, joinedAt: tc, name: name, rank: rank, score: 0, team: team}
	gpt.ship = newShip(ship)
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

func (gt *gameType) addPlayerToGame(pt *playerType, ship shipNameEnum, team teamEnum) *gamePlayerType {
	// ensure there are no stale entries

	var blue_population, red_population uint16

	for ndx, val := range gt.playerArray {
		// no duplicate players
		if val != nil && val.key.key == pt.key.key {
			gt.sugarLog.Info("duplicate player key")
		}

		// no duplicate ships
		if val != nil && val.ship.name == ship {
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

	// add new player to game
	gpt := newGamePlayer(pt.key, pt.name, pt.rank, ship, team, gt.currentTurn)
	for ndx, val := range gt.playerArray {
		if val == nil {
			gt.playerArray[ndx] = gpt
			gt.sugarLog.Infof("new player %s %s %s", gpt.name, gpt.team.string(), gpt.ship.name.string())
			break
		}
	}

	gt.moveTo(gpt.ship)

	return gpt
}

func (gt *gameType) evictPlayer(gpt *gamePlayerType) {
	gt.sugarLog.Infof("evicting player %s %s %s", gpt.name, gpt.team.string(), gpt.ship.name.string())

	gt.gameBoard[gpt.ship.location.row][gpt.ship.location.col].tokenKey = nil
	gt.gameBoard[gpt.ship.location.row][gpt.ship.location.col].tokenType = emptyToken
}

func (gt *gameType) testForPlayerEviction() {
	for ndx, gpt := range gt.playerArray {
		if gpt != nil && !gpt.activeFlag {
			gt.evictPlayer(gpt)
		}

		gt.playerArray[ndx] = nil
	}
}

func (gt *gameType) moveFrom(lt *locationType) {
	gt.gameBoard[lt.row][lt.col].tokenKey = nil
	gt.gameBoard[lt.row][lt.col].tokenType = emptyToken
}

func (gt *gameType) moveTo(st *shipType) {
	gt.sugarLog.Infof("moveTo %s row:%d col:%d", st.name.string(), st.location.row, st.location.col)
	gt.gameBoard[st.location.row][st.location.col].tokenKey = st.key
	gt.gameBoard[st.location.row][st.location.col].tokenType = shipToken
}

func (gt *gameType) getOutput(pkt *tokenKeyType) commandArrayType {
	return nil
}
