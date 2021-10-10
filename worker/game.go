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

	shutDownFlag bool
	//turnEventQueue turnEventQueueType
	//eventQueueNdx  int // current queue index
	turnCounter int // current game turn

	//outQueue outputType

	rdb *redis.Client

	inboundQueue  string
	outboundQueue string

	uuid string // game identifier
}

func newGame(id string, boardType boardTypeEnum) *gameType {
	log.Println("new game:", id, boardType.string())

	rand.Seed(time.Now().UnixNano())

	gt := gameType{uuid: id, boardType: boardType}
	gt.inboundQueue = id + "m"
	gt.outboundQueue = id + "w"

	gt.commandQueue = newCommandQueue()

	gt.rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "bigSekret",
		DB:       0, // use default DB
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
	playerNdx := gt.players.find(tnt.name)
	if playerNdx < 0 {
		log.Println("skipping command w/unknown player")
		return
	}

	// schedule event for now
	tqn := gt.turnCounter
	if gt.players[playerNdx].turnQueueNdx >= gt.turnCounter {
		// schedule event for later
		tqn = gt.players[playerNdx].turnQueueNdx
	}

	ndx := tqn % maxTurnQueueArray

	gt.players[playerNdx].turnQueue[ndx].enqueue(tnt)
	gt.players[playerNdx].turnQueueNdx = tqn + tnt.duration

	//turnQueueArrayDump(gt.players[playerNdx].turnQueue)
}

func (gt *gameType) dispatchCommand(tnt *turnNodeType) {
	//var response *CommandType

	switch tnt.command {
	case moveCommand:
		response, err := commandShipMove(tnt, &gt.board, &gt.ships)
		log.Println(err)
		log.Println(response)
	case pingCommand:
		response, err := commandPing(tnt)
		log.Println(err)
		log.Println(response)
	case playerCreateCommand:
		commandPlayerCreate(tnt, &gt.players)
	case playerDeleteCommand:
		commandPlayerDelete(tnt, &gt.board, &gt.ships, &gt.players)
	case shipCreateCommand:
		commandShipCreate(tnt, &gt.board, &gt.ships)
	case shipDeleteCommand:
		commandShipDelete(tnt, &gt.board, &gt.ships)
	case shutDownCommand:
		log.Println("shutdown noted")
		gt.shutDownFlag = true
	}

	/*
		if response != nil {
			responseToManager(gt.outboundQueue, response)
		}
	*/
}

func (gt *gameType) serviceCommandQueue() {
	for {
		// process fresh messages from manager
		ct := gt.commandQueue.dequeue()
		if ct == nil {
			break
		} else {
			tnt := newTurnNode(ct)

			switch tnt.command {
			case pingCommand:
				gt.dispatchCommand(tnt)
			case playerCreateCommand:
				gt.dispatchCommand(tnt)
			case playerDeleteCommand:
				gt.dispatchCommand(tnt)
			case shipCreateCommand:
				gt.dispatchCommand(tnt)
			case shipDeleteCommand:
				gt.dispatchCommand(tnt)
			case shutDownCommand:
				gt.dispatchCommand(tnt)
			default:
				gt.scheduleTurnEvent(tnt)
			}
		}
	}
}

func (gt *gameType) servicePlayerTurnQueue(playerNdx int) {
	if gt.players[playerNdx] == nil {
		log.Printf("skipping nil player %d", playerNdx)
		return
	}

	ndx := gt.turnCounter % maxTurnQueueArray

	log.Printf("player %d turn %d %d", playerNdx, gt.turnCounter, ndx)

	if gt.players[playerNdx].turnQueue[ndx].size > 0 {
		tnt := gt.players[playerNdx].turnQueue[ndx].dequeue()
		log.Println(tnt.commands[0])
		gt.dispatchCommand(tnt)
	}
}

func (gt *gameType) serviceTurnQueue() {
	// need random option
	for ndx := 0; ndx < maxPlayers; ndx++ {
		gt.servicePlayerTurnQueue(ndx)
	}
}

func (gt *gameType) serviceResponseQueue() {
	log.Println("service response queue")
}

func (gt *gameType) turnManager() {
	gt.turnCounter++
	//gt.eventQueueNdx = gt.turnCounter % maxTurnEventQueue
	//log.Printf("starting turn:%d %d", gt.turnCounter, gt.eventQueueNdx)

	gt.serviceCommandQueue()

	gt.serviceTurnQueue()

	gt.serviceResponseQueue()

	log.Printf("ending turn:%d", gt.turnCounter)
}
