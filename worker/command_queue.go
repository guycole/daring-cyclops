// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import "log"

type commandNodeType struct {
	payload *CommandType
	next    *commandNodeType
}

type commandQueueType struct {
	head *commandNodeType
	tail *commandNodeType
	size int
}

func newCommandQueue() *commandQueueType {
	return &commandQueueType{}
}

func (cqt *commandQueueType) dump() {
	log.Println("=-=-=-= Command Queue Dump =-=-=-=")

	current := cqt.head
	for {
		if current == nil || current.payload == nil {
			break
		}

		log.Println(current.payload)
		current = current.next
	}

	log.Println("=-=-=-= Command Queue Dump =-=-=-=")
}

func (cqt *commandQueueType) isEmpty() bool {
	if cqt.size == 0 {
		return true
	}

	return false
}

func (cqt *commandQueueType) dequeue() *CommandType {
	if cqt.size <= 0 {
		return nil
	}

	result := cqt.head.payload

	cqt.head = cqt.head.next
	cqt.size--

	return result
}

func (cqt *commandQueueType) enqueue(payload *CommandType) {
	node := commandNodeType{payload: payload}

	if cqt.size <= 0 {
		cqt.head = &node
		cqt.tail = &node
		cqt.size = 1
	} else {
		cqt.tail.next = &node
		cqt.tail = &node
		cqt.size++
	}
}
