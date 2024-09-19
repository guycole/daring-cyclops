// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

const (
	maxScheduleArray uint16 = 100
)

type scheduleType struct {
	ct   *commandType
	next *scheduleType
}

type scheduleArrayType [maxScheduleArray]*scheduleType

func (gt *gameType) schedule(ct *commandType) {

}
