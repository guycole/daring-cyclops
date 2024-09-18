// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

type userPointType struct {
	age turnCounterType
}

func (gt *gameType) pointCommand(ct *commandType) *userPointType {
	gt.sugarLog.Debug("pointsCommand")

	upt := userPointType{age: gt.currentTurn}

	return &upt
}
