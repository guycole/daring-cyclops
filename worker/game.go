// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"log"
	"math/rand"
	"time"

	redis "github.com/go-redis/redis/v8"
)

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

	commandStack *commandStackType

	turnEventQueue turnEventQueueType
	eventQueueNdx  int // current queue index
	turnCounter    int // current game turn

	//outQueue outputType

	rdb *redis.Client

	uuid string // game identifier
}

func newGame(id string, boardType boardTypeEnum) *gameType {
	log.Println("new game:", id, boardType.string())

	rand.Seed(time.Now().UnixNano())

	gt := gameType{uuid: id, boardType: boardType}
	gt.commandStack = newCommandStack()
	gt.rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	gt.board = newBoard()
	boardGenerator(&gt)

	return &gt
}

func (gt *gameType) serviceCommandStack() {
	for {
		// process fresh messages from manager
		temp := gt.commandStack.pop()
		if temp == nil {
			break
		} else {
			newEvent := newTurnEvent(temp)
			if newEvent == nil {
				log.Println("skipping bad event:", temp)
			} else {
				log.Println("must schedule:", newEvent)
				turnEventQueuePush(newEvent, gt.turnEventQueue)
			}
		}
	}
}

// turnManager schedules game play
func (gt *gameType) turnManager() {
	gt.turnCounter++
	gt.eventQueueNdx = gt.turnCounter % maxTurnEventQueue
	log.Printf("starting turn:%d %d", gt.turnCounter, gt.eventQueueNdx)

	gt.serviceCommandStack()

	/*
		serviceEventQueue(gt)
	*/

	serviceOutboundQueue(gt)

	log.Printf("ending turn:%d", gt.turnCounter)
}

// serviceOutboundQueue by writing all pending traffic to manager
func serviceOutboundQueue(gt *gameType) {
	log.Printf("serviceOutboundQueue:%d", gt.turnCounter)
}
