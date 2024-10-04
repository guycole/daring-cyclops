// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"strings"

	"go.uber.org/zap"
)

const (
	// circular array of scheduled commands indexed by turn
	maxScheduleArray uint16 = 100
)

type scheduleType struct {
	root, tail *commandType // single linked list of commands
	length     uint16       // list population
}

// convenience factory
func newSchedule() *scheduleType {
	result := scheduleType{root: nil, tail: nil, length: 0}
	return &result
}

// circular array of scheduled items
type scheduleArrayType [maxScheduleArray]*scheduleType

func newScheduleArray() scheduleArrayType {
	var sat scheduleArrayType

	for ndx, _ := range sat {
		sat[ndx] = newSchedule()
	}

	return sat
}

func (st *scheduleType) scheduleDumper(sugarLog *zap.SugaredLogger) {
	sugarLog.Debug("====> schedule dump <====")
	sugarLog.Debugf("length %d", st.length)

	var temp *commandType
	for temp = st.root; temp != nil; temp = temp.next {
		commandName := legalGameCommands[temp.command].longName
		sugarLog.Debugf("%s %s", commandName, temp.sourcePlayerKey.key)
	}

	sugarLog.Debug("====> schedule dump <====")
}

// add a fresh command to end of list
func (st *scheduleType) scheduleAdd(ct *commandType) {
	ct.next = nil

	if st.length < 1 {
		st.root = ct
	} else {
		st.tail.next = ct
	}

	st.tail = ct
	st.length++
}

// return first command in list which matches player
func (st *scheduleType) scheduleSelect(target *tokenKeyType) *commandType {
	var last, current *commandType

	last = nil
	for current = st.root; current != nil; current = current.next {
		if strings.Compare(current.sourcePlayerKey.key, target.key) == 0 {
			if last == nil {
				st.root = st.root.next
			} else {
				last.next = current.next
			}

			st.length = st.length - 1
			current.next = nil
			return current
		}

		last = current
	}

	return current
}
