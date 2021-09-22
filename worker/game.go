// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found
// in the LICENSE file.
package main

import (
	"log"
	"math/rand"
	"time"

	redis "github.com/go-redis/redis/v8"
)

const maxEventQueue = 10

// gameType main game structure, only one instance per game
type gameType struct {
	// TODO creation time
	board     boardArrayType
	boardType boardTypeEnum
	planets   planetArrayType
	players   playerArrayType
	ships     shipArrayType
	stars     starArrayType
	starGates starGateArrayType

	//eventQueue    [maxEventQueue]turnEventType
	eventQueueNdx int // current queue index
	turnCounter   int // current game turn

	//outQueue outputType

	rdb *redis.Client

	uuid string // game identifier
}

/*
type turnEventType struct {
	payload *commandType // single linked list of commands in execution order
}

type outputType struct {
	player  string // player uuid
	request string // request uuid

	args []string

	next *outputType
}
*/

/*
func eventQueueSimulator(gt *gameType) {
	log.Println("simulator")

	message1 := `{"player":"player1uuid", "requestId":"request1uuid", "command":["users"]}`
	log.Println(message1)
	temp1 := parseJsonCommand(message1, gt.turnCounter)
	if temp1 != nil {
		eventQueuePush(*temp1, gt)
	}

	message2 := `{"player":"player1uuid", "requestId":"request1uuid", "command":["createPlayer", "playerName1", "captain", "blue"]}`
	log.Println(message2)
	temp2 := parseJsonCommand(message2, gt.turnCounter)
	if temp2 != nil {
		eventQueuePush(*temp2, gt)
	}

	message3 := `{"player":"player2uuid", "requestId":"request2uuid", "command":["createPlayer", "playerName2", "admiral", "red"]}`
	log.Println(message3)
	temp3 := parseJsonCommand(message3, gt.turnCounter)
	if temp3 != nil {
		eventQueuePush(*temp3, gt)
	}

	message4 := `{"player":"player1uuid", "requestId":"request1uuid", "command":["createShip", "nimrod"]}`
	log.Println(message4)
	temp4 := parseJsonCommand(message4, gt.turnCounter)
	if temp4 != nil {
		eventQueuePush(*temp4, gt)
	}

	message5 := `{"player":"player1uuid", "requestId":"request1uuid", "command":["move", "3", "3"]}`
	log.Println(message5)
	temp5 := parseJsonCommand(message5, gt.turnCounter)
	if temp5 != nil {
		eventQueuePush(*temp5, gt)
	}

	message6 := `{"player":"player1uuid", "requestId":"request1uuid", "command":["move", "3", "3"]}`
	log.Println(message6)
	temp6 := parseJsonCommand(message6, gt.turnCounter)
	if temp6 != nil {
		eventQueuePush(*temp6, gt)
	}
}
*/

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

/*
// eventQueuePush add event to queue
func eventQueuePush(ct commandType, gt *gameType) {
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
}
*/

func newGame(id string, boardType boardTypeEnum) *gameType {
	log.Println("new game:", id, boardType.string())

	rand.Seed(time.Now().UnixNano())

	gt := gameType{uuid: id, boardType: boardType}
	gt.board = newBoard()
	boardGenerator(&gt)

	gt.rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &gt
}

// turnManager schedules game play
func turnManager(gt *gameType) {
	gt.turnCounter++
	gt.eventQueueNdx = gt.turnCounter % maxEventQueue
	log.Printf("starting turn:%d %d", gt.turnCounter, gt.eventQueueNdx)

	/*
		serviceInboundQueue(gt)
		serviceEventQueue(gt)
		serviceOutboundQueue(gt)
	*/

	log.Printf("ending turn:%d", gt.turnCounter)
}

/*
// serviceInboundQueue read from RabbitMQ and add to event queue
func serviceInboundQueue(gt *gameType) {
	log.Printf("serviceInboundQueue:%d", gt.turnCounter)

	if gt.turnCounter == 1 {
		eventQueueSimulator(gt)
		eventQueueDump(*gt)
	}
}
*/

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

// serviceOutboundQueue by writing all pending traffic to manager
func serviceOutboundQueue(gt *gameType) {
	log.Printf("serviceOutboundQueue:%d", gt.turnCounter)
}
