// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"strings"
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

func TestTurnQueueSimpleOps(t *testing.T) {
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

func seedTestShips() (*CommandType, *CommandType) {
	var commands commandArrayType
	commands[0] = "shipCreate"
	commands[1] = "nike"

	ct1 := newCommand(testPlayerName1, "reqId3", 2, commands)

	commands[0] = "shipCreate"
	commands[1] = "triton"

	ct2 := newCommand(testPlayerName2, "reqId4", 2, commands)

	return ct1, ct2
}

func seedTestUsers() (*CommandType, *CommandType) {
	var commands commandArrayType
	commands[0] = "playerCreate"
	commands[1] = "captain"
	commands[2] = "blue"

	ct1 := newCommand(testPlayerName1, "reqId1", 3, commands)

	commands[0] = "playerCreate"
	commands[1] = "admiral"
	commands[2] = "red"

	ct2 := newCommand(testPlayerName2, "reqId2", 3, commands)

	return ct1, ct2
}

func TestTurnQueueOps(t *testing.T) {
	gt := newGame("testGame", emptyBoard)

	ct1, ct2 := seedTestUsers()
	gt.commandQueue.enqueue(ct1)
	gt.commandQueue.enqueue(ct2)

	ct3, ct4 := seedTestShips()
	gt.commandQueue.enqueue(ct3)
	gt.commandQueue.enqueue(ct4)

	gt.commandQueue.dump()

	gt.serviceCommandQueue()

	// users and ships created
	// now move ships to create future ops

	shipNdx := gt.ships.findByOwner(testPlayerName1)
	if shipNdx < 0 {
		t.Error("unable to discover ship by owner")
	}

	// move ship to known location
	newLocation := newLocation(20, 40)
	err := gt.ships.move(gt.ships[shipNdx].uuid, newLocation, &gt.board)
	if err != nil {
		t.Error("unable to move ship")
	}

	// verify ship struct
	position := gt.ships[shipNdx].position
	if position.yy != 20 || position.xx != 40 {
		t.Error("bad ship position")
	}

	// verify board cell
	bc := gt.board[20][40]
	if !bc.ship {
		t.Error("board cell ship fail")
	}

	if strings.Compare(bc.shipSymbol, gt.ships[shipNdx].symbol) != 0 {
		t.Error("boad cell symbol fail")
	}

	if strings.Compare(bc.tokenID, gt.ships[shipNdx].uuid) != 0 {
		t.Error("boad cell tokenID fail")
	}

	// ship move
	var commands commandArrayType
	commands[0] = "move"
	commands[1] = "25"
	commands[2] = "45"

	ct5 := newCommand(testPlayerName1, "reqId5", 3, commands)
	gt.commandQueue.enqueue(ct5)
}
