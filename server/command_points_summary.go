// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

type userPointsType struct {
	age turnCounterType
}

func (gt *gameType) pointsCommand(ct *commandType) *userPointsType {
	gt.sugarLog.Debug("pointsCommand")

	upt := userPointsType{age: gt.currentTurn}

	return &upt
}
