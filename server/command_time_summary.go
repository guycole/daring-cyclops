// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

type userTimeType struct {
	age turnCounterType
}

func (gt *gameType) timeCommand(ct *commandType) *userTimeType {
	gt.sugarLog.Debug(ct.command)

	utt := userTimeType{age: gt.currentTurn}

	return &utt
}
