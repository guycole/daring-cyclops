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
	//	planets   planetArrayType
	players playerArrayType
	ships   shipArrayType
	//	stars     starArrayType
	//	starGates starGateArrayType

	commandQueue *commandQueueType

	//turnEventQueue turnEventQueueType
	//eventQueueNdx  int // current queue index
	turnCounter int // current game turn

	//outQueue outputType

	rdb *redis.Client

	uuid string // game identifier
}

func newGame(id string, boardType boardTypeEnum) *gameType {
	log.Println("new game:", id, boardType.string())

	rand.Seed(time.Now().UnixNano())

	gt := gameType{uuid: id, boardType: boardType}
	gt.commandQueue = newCommandQueue()
	gt.rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	gt.board = newBoard()
	gt.boardGenerator()

	return &gt
}

func (gt *gameType) boardGenerator() {
	switch gt.boardType {
	case emptyBoard:
		log.Println("generating empty board")
	case standardBoard:
		log.Println("generating standard board")
		//		starGatesAdd(&gt.starGates, &gt.board)
		//		starsAdd(&gt.stars, &gt.board)
		//		planetsAdd(&gt.planets, &gt.board)
	default:
		log.Println("unsupported boardType in boardGenerator")
	}
}

func (gt *gameType) testShip1() *shipType {
	position := randomShipLocation(gt.board)
	ns1, _ := newShip(testShipName1, testPlayerName1, position)
	ns1.uuid = testShipUuid1
	return ns1
}

func (gt *gameType) testShip2() *shipType {
	position := randomShipLocation(gt.board)
	ns2, _ := newShip(testShipName2, testPlayerName2, position)
	ns2.uuid = testShipUuid2
	return ns2
}

func (gt *gameType) scheduleTurnEvent(tnt *turnNodeType) {
	playerNdx := gt.players.playerFind(tnt.name)
	if playerNdx < 0 {
		log.Println("skipping command w/unknown player")
		return
	}

	// schedule event for now
	tqn := gt.turnCounter

	if gt.players[playerNdx].turnQueueNdx >= gt.turnCounter {
		// schedle event for later
		tqn = gt.turnCounter + gt.players[playerNdx].turnQueueNdx
	}

	ndx := tqn % maxTurnQueueArray

	gt.players[playerNdx].turnQueue[ndx].enqueue(tnt)

	gt.players[playerNdx].turnQueueNdx = tqn + tnt.duration
}

func (gt *gameType) serviceCommandStack() {
	for {
		// process fresh messages from manager
		ct := gt.commandQueue.dequeue()
		if ct == nil {
			break
		} else {
			tnt := newTurnNode(ct)

			// process admin commands immediately
			switch tnt.command {
			case playerCreateCommand:
				commandPlayerCreate(tnt, &gt.players)
			case playerDeleteCommand:
				commandPlayerDelete(tnt, &gt.players)
			case shipCreateCommand:
				log.Println("ship create")
			case shipDeleteCommand:
				log.Println("ship delete")
			}

			// schedule player commands for future execution
			gt.scheduleTurnEvent(tnt)
		}
	}
}

func (gt *gameType) serviceEventQueue() {
	// FIXME commands for this turn
}

func (gt *gameType) turnManager() {
	gt.turnCounter++
	//gt.eventQueueNdx = gt.turnCounter % maxTurnEventQueue
	//log.Printf("starting turn:%d %d", gt.turnCounter, gt.eventQueueNdx)

	gt.serviceCommandStack()
	gt.serviceEventQueue()

	log.Printf("ending turn:%d", gt.turnCounter)
}
