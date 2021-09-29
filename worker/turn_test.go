// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"testing"
)

func TestTurnQueueEmpty(t *testing.T) {
	tq := newTurnQueue()

	flag := tq.isEmpty()
	if !flag {
		t.Error("queue should be reported empty")
	}

	tqn := tq.dequeue()
	if tqn != nil {
		t.Error("dequeue should return nil")
	}
}

func TestTurnQueueNotEmpty(t *testing.T) {
	var commands commandArrayType
	commands[0] = "pingCommand"
	ct := newCommand(testPlayerName1, "reqId1", 1, commands)
	tnt := newTurnNode(ct)

	tq := newTurnQueue()
	tq.enqueue(tnt)

	flag := tq.isEmpty()
	if flag {
		t.Error("stack should NOT be reported empty")
	}

	tnt = tq.dequeue()
	if tnt == nil {
		t.Error("dequeue should NOT return nil")
	}
}

func TestTurnQueueOps(t *testing.T) {
	var commands commandArrayType
	commands[0] = "pingCommand"

	ct1 := newCommand(testPlayerName1, "reqId1", 1, commands)
	tnt1 := newTurnNode(ct1)
	ct2 := newCommand(testPlayerName1, "reqId2", 1, commands)
	tnt2 := newTurnNode(ct2)
	ct3 := newCommand(testPlayerName1, "reqId3", 1, commands)
	tnt3 := newTurnNode(ct3)

	tq := newTurnQueue()
	tq.enqueue(tnt1)
	tq.enqueue(tnt2)
	tq.enqueue(tnt3)

	tq.dump()

	if tq.size != 3 {
		t.Error("size returns bad value")
	}

	temp1 := tq.dequeue()
	if temp1.request != "reqId1" {
		t.Error("dq returns bad value")
	}

	temp2 := tq.dequeue()
	if temp2.request != "reqId2" {
		t.Error("dq returns bad value")
	}

	temp3 := tq.dequeue()
	if temp3.request != "reqId3" {
		t.Error("dq returns bad value")
	}

	temp4 := tq.dequeue()
	if temp4 != nil {
		t.Error("dq did not return nil")
	}
}
