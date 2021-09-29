// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import "log"

type turnNodeType struct {
	name    string // player name
	request string // request uuid

	duration int // command duration (in turns)
	turn     int // turn counter for command execution

	commandSize int
	commands    commandArrayType

	command commandGameEnum

	next *turnNodeType
}

func newTurnNode(ct *CommandType) *turnNodeType {
	result := turnNodeType{name: ct.Name, request: ct.RequestId, commandSize: ct.CommandSize, commands: ct.Commands}

	result.command = findGameCommand(ct.Commands[0])
	result.duration = findGameCommandDuration(result.command)

	if result.command == unknownCommand {
		log.Println("unknown unknown")
	} else {
		log.Println("not unknown")
	}

	return &result
}

type turnQueueType struct {
	head *turnNodeType
	tail *turnNodeType
	size int
}

func newTurnQueue() *turnQueueType {
	return &turnQueueType{}
}

func (tqt *turnQueueType) isEmpty() bool {
	if tqt.size == 0 {
		return true
	}

	return false
}

func (tqt *turnQueueType) dequeue() *turnNodeType {
	if tqt.size <= 0 {
		return nil
	}

	result := tqt.head

	tqt.head = tqt.head.next
	tqt.size--

	return result
}

func (tqt *turnQueueType) enqueue(node *turnNodeType) {
	if tqt.size <= 0 {
		tqt.head = node
		tqt.tail = node
		tqt.size = 1
	} else {
		tqt.tail.next = node
		tqt.tail = node
		tqt.size++
	}
}

func (tqt *turnQueueType) dump() {
	log.Println("=-=-=-= Turn Queue Dump =-=-=-=")

	current := tqt.head
	for {
		if current == nil {
			break
		}

		log.Println(current)
		current = current.next
	}

	log.Println("=-=-=-= Turn Queue Dump =-=-=-=")
}

const maxTurnQueueArray = 10

type turnQueueArrayType [maxTurnQueueArray]*turnQueueType

func turnQueueArrayDump(tqat turnQueueArrayType) {
	log.Println("=-=-=-= Turn Queue Array Dump =-=-=-=")

	for ndx := 0; ndx < maxTurnQueueArray; ndx++ {
		tqt := tqat[ndx]

		for {
			log.Printf("%d %d", ndx, tqt.size)
			if tqt.size > 0 {
				tqt.dump()
			}
		}
	}

	log.Println("=-=-=-= Turn Queue Array Dump =-=-=-=")
}
