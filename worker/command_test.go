// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"log"
	"testing"
)

func TestCommandStackEmpty(t *testing.T) {
	root := newCommandStack()

	flag := root.isEmpty()
	if !flag {
		t.Error("stack should be reported empty")
	}

	ct := root.pop()
	if ct != nil {
		t.Error("pop should return nil")
	}
}
func TestCommandStackNotEmpty(t *testing.T) {
	command := newCommand(testPlayerName1, "reqId1")

	root := newCommandStack()

	root.push(command)

	flag := root.isEmpty()
	if flag {
		t.Error("stack should NOT be reported empty")
	}

	ct := root.pop()
	if ct != nil {
		t.Error("pop should NOT return nil")
	}
}

func TestCommandStackOps(t *testing.T) {
	log.Println("xoxoxoxoxoxoxoxoxo")

	command1 := newCommand(testPlayerName1, "reqId1")
	command2 := newCommand(testPlayerName1, "reqId2")
	command3 := newCommand(testPlayerName1, "reqId3")

	root := newCommandStack()

	root.push(command1)
	root.push(command2)
	root.push(command3)

	root.dump()

	temp1 := root.pop()
	if temp1.RequestId != "reqId3" {
		t.Error("pop returns bad value")
	}

	temp2 := root.pop()
	if temp2.RequestId != "reqId2" {
		t.Error("pop returns bad value")
	}

	temp3 := root.pop()
	if temp3.RequestId != "reqId1" {
		t.Error("pop returns bad value")
	}

	temp4 := root.pop()
	if temp4 != nil {
		t.Error("pop did not return nil")
	}
}
