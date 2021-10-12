// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import "log"

type requestNodeType struct {
	payload *RequestType
	next    *requestNodeType
}

func newRequestNode(rt *RequestType) *requestNodeType {
	return &requestNodeType{payload: rt}
}

type requestQueueType struct {
	head *requestNodeType
	tail *requestNodeType
	size int
}

func newRequestQueue() *requestQueueType {
	return &requestQueueType{}
}

func (rqt *requestQueueType) dump() {
	log.Println("=-=-=-= Request Queue Dump =-=-=-=")

	current := rqt.head
	for {
		if current == nil || current.payload == nil {
			break
		}

		log.Println(current.payload)
		current = current.next
	}

	log.Println("=-=-=-= Request Queue Dump =-=-=-=")
}

func (rqt *requestQueueType) isEmpty() bool {
	if rqt.size == 0 {
		return true
	}

	return false
}

func (rqt *requestQueueType) dequeue() *RequestType {
	if rqt.size <= 0 {
		return nil
	}

	result := rqt.head.payload

	rqt.head = rqt.head.next
	rqt.size--

	return result
}

func (rqt *requestQueueType) enqueue(payload *RequestType) {
	node := newRequestNode(payload)

	if rqt.size <= 0 {
		rqt.head = node
		rqt.tail = node
		rqt.size = 1
	} else {
		rqt.tail.next = node
		rqt.tail = node
		rqt.size++
	}
}
