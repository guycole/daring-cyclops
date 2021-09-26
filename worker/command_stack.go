// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import "log"

// LIFO
type commandStackType struct {
	payload *CommandType
	next    *commandStackType
}

func newCommandStack() *commandStackType {
	return &commandStackType{}
}

func (root *commandStackType) dump() {
	log.Println("=-=-=-= stackDump =-=-=-=")

	current := root
	for {
		if current == nil || current.payload == nil {
			break
		}

		log.Println(current.payload)
		current = current.next
	}

	log.Println("=-=-=-= stackDump =-=-=-=")
}

func (root *commandStackType) isEmpty() bool {
	if root == nil || root.payload == nil {
		return true
	}

	return false
}

func (root *commandStackType) pop() *CommandType {
	if root == nil || root.payload == nil {
		return nil
	}

	response := root.payload

	if root.next == nil {
		// empty stack
		root.payload = nil
	} else {
		// next becomes top of stack
		root.payload = root.next.payload
		root.next = root.next.next
	}

	return response
}

func (root *commandStackType) push(payload *CommandType) {
	if root.payload == nil {
		// empty stack gets fresh payload
		root.payload = payload
	} else {
		// new command becomes list root
		next := &commandStackType{payload: root.payload, next: root.next}
		root.payload = payload
		root.next = next
	}
}
