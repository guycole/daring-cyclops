// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"strings"
	"testing"

	shared "github.com/guycole/daring-cyclops/shared"
)

func TestScheduleOperations(t *testing.T) {
	sugarLog := shared.ZapSetup(true)
	sugarLog.Info("game schedule test")

	pt1key := newPlayerKey(testPlayer1)
	pt2key := newPlayerKey(testPlayer2)

	ct1a := newCommand(userCommand, pt1key)
	ct1b := newCommand(timeCommand, pt1key)

	ct2a := newCommand(userCommand, pt2key)
	ct2b := newCommand(timeCommand, pt2key)

	st := newSchedule()
	if st.length != 0 {
		t.Error("length failure 0")
	}

	st.scheduleAdd(ct1a)
	st.scheduleAdd(ct2a)
	st.scheduleAdd(ct1b)
	st.scheduleAdd(ct2b)
	if st.length != 4 {
		t.Error("length failure 4")
	}

	st.scheduleDumper(sugarLog)

	// consume in middle
	selected := st.scheduleSelect(pt2key)
	if selected == nil {
		t.Error("selection error nil")
	}

	if strings.Compare(selected.sourcePlayerKey.key, pt2key.key) != 0 {
		t.Error("selection error with bad player key")
	}

	if st.length != 3 {
		t.Error("length failure 3")
	}

	// consume at root
	selected = st.scheduleSelect(pt1key)
	if selected == nil {
		t.Error("selection error nil")
	}

	// consume at tail
	selected = st.scheduleSelect(pt2key)
	if selected == nil {
		t.Error("selection error nil")
	}

	// consume last element
	selected = st.scheduleSelect(pt1key)
	if selected == nil {
		t.Error("selection error nil")
	}

	st.scheduleDumper(sugarLog)
}
