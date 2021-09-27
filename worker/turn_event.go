// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import "log"

type turnEventType struct {
	name    string // player name
	request string // request uuid

	duration int // command duration (in turns)
	turn     int // turn counter for command execution

	commandSize int
	commands    commandArrayType

	command commandGameEnum

	next *turnEventType
}

const maxTurnEventQueue = 10

type turnEventQueueType [maxTurnEventQueue]*turnEventType

func newTurnEvent(ct *CommandType) *turnEventType {
	result := turnEventType{name: ct.Name, request: ct.RequestId, commandSize: ct.CommandSize, commands: ct.Commands}

	result.command = findGameCommand(ct.Commands[0])
	result.duration = findGameCommandDuration(result.command)

	if result.command == unknownCommand {
		log.Println("unknown unknown")
	} else {
		log.Println("not unknown")
	}

	return &result
}

// eventQueueDump writes event queue to stdout
func eventQueueDump(gt gameType) {
	log.Println("=-=-=-= eventQueueDump =-=-=-=")

	/*
		for ndx := 0; ndx < maxEventQueue; ndx++ {
			temp := gt.eventQueue[ndx].payload

			for {
				if temp == nil {
					log.Printf("%d %d nil", ndx, gt.eventQueueNdx)
					break
				} else {
					log.Printf("%d %d %s %s", ndx, gt.eventQueueNdx, temp.player, temp.raw)
					temp = temp.next
				}
			}
		}
	*/

	log.Println("=-=-=-= eventQueueDump =-=-=-=")
}

/*
// eventQueuePop consume event from queue
func eventQueuePop(gt *gameType) *commandType {
	payload := gt.eventQueue[gt.eventQueueNdx].payload
	if payload != nil {
		gt.eventQueue[gt.eventQueueNdx].payload = payload.next
	}

	return payload
}
*/

func turnEventQueuePush(tet *turnEventType, teqt turnEventQueueType) {

	/*
		ndx := ct.turn % maxEventQueue

		if gt.eventQueue[ndx].payload == nil {
			// first command
			gt.eventQueue[ndx].payload = &ct
		} else {
			// new tail
			current := gt.eventQueue[ndx].payload
			for ; current.next != nil; current = current.next {
				// empty
			}

			current.next = &ct
		}
	*/
}

// serviceEventQueue dispatch events
func serviceEventQueue(gt *gameType) {
	log.Printf("serviceEventQueue:%d", gt.turnCounter)
	/*
		for {
			current := eventQueuePop(gt)
			if current == nil {
				break
			} else {
				dispatchCommand(*current, gt)
			}
		}
	*/
}
