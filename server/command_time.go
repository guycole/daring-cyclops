// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

type timeRequestType struct{}

// convenience factory
func newTimeRequest(playerKey *tokenKeyType) *commandType {
	ct := commandType{command: timeCommand, sourcePlayerKey: playerKey}
	ct.timeRequest = &timeRequestType{}
	return &ct
}

type timeType struct {
	age turnCounterType
}

type timeResponseType struct {
	tt *timeType
}

func (gt *gameType) timeCommand(ct *commandType) *commandType {
	gt.sugarLog.Debug("timeCommand")

	tt := timeType{age: gt.currentTurn}

	response := timeResponseType{tt: &tt}
	gt.sugarLog.Debug(response)

	ct.timeResponse = &response

	ct.destinationPlayerKeys = append(ct.destinationPlayerKeys, ct.sourcePlayerKey)

	return ct
}
