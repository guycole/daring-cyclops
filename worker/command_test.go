// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"log"
	"testing"
)

func TestCommandQueueEmpty(t *testing.T) {
	cq := newCommandQueue()

	flag := cq.isEmpty()
	if !flag {
		t.Error("queue should be reported empty")
	}

	ct := cq.dequeue()
	if ct != nil {
		t.Error("dequeue should return nil")
	}
}

func TestCommandQueueNotEmpty(t *testing.T) {
	var commands commandArrayType
	commands[0] = "stubCommand"

	command := newCommand(testPlayerName1, "reqId1", 1, commands)

	cq := newCommandQueue()

	cq.enqueue(command)

	flag := cq.isEmpty()
	if flag {
		t.Error("stack should NOT be reported empty")
	}

	ct := cq.dequeue()
	if ct == nil {
		t.Error("dequeue should NOT return nil")
	}
}

func TestCommandQueueOps(t *testing.T) {
	log.Println("xoxoxoxoxoxoxoxoxo")

	var commands commandArrayType
	commands[0] = "stubCommand"

	command1 := newCommand(testPlayerName1, "reqId1", 1, commands)
	command2 := newCommand(testPlayerName1, "reqId2", 1, commands)
	command3 := newCommand(testPlayerName1, "reqId3", 1, commands)

	cq := newCommandQueue()

	cq.enqueue(command1)
	cq.enqueue(command2)
	cq.enqueue(command3)

	cq.dump()

	if cq.size != 3 {
		t.Error("size returns bad value")
	}

	temp1 := cq.dequeue()
	if temp1.RequestId != "reqId1" {
		t.Error("dq returns bad value")
	}

	temp2 := cq.dequeue()
	if temp2.RequestId != "reqId2" {
		t.Error("dq returns bad value")
	}

	temp3 := cq.dequeue()
	if temp3.RequestId != "reqId3" {
		t.Error("dq returns bad value")
	}

	temp4 := cq.dequeue()
	if temp4 != nil {
		t.Error("dq did not return nil")
	}
}
