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
	creation  time.Time
	board     boardArrayType
	boardType boardTypeEnum
	//	planets   planetArrayType
	players playerArrayType
	//ships   shipArrayType
	//	stars     starArrayType
	//	starGates starGateArrayType

	requestQueue *requestQueueType

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

type boardTypeEnum int

const (
	unknownBoardType boardTypeEnum = iota
	emptyBoard                     // no mines, planets, ships, stargates, voids
	standardBoard                  // standard game map w/stars, planets, stargates and voids
)

// must match order for boardTypeEnum
func (bte boardTypeEnum) string() string {
	return [...]string{"unknown", "empty", "standard"}[bte]
}

func newGame(id string, boardType boardTypeEnum) *gameType {
	log.Println("new game:", id, boardType.string())

	rand.Seed(time.Now().UnixNano())

	gt := gameType{uuid: id, boardType: boardType}

	gt.creation = time.Now()

	gt.inboundQueue = id + "m"
	gt.outboundQueue = id + "w"

	gt.requestQueue = newRequestQueue()

	// TODO get these arguments from secrets
	gt.rdb = redis.NewClient(&redis.Options{
		Addr:     "cyclops-redis-master:6379",
		Password: "bigSekret",
		DB:       0, // use default DB
	})

	gt.boardGenerator()

	return &gt
}

func (gt *gameType) boardGenerator() {
	gt.board = newBoard()

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

/*
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
*/

func (gt *gameType) scheduleTurnEvent(tnt *turnNodeType) {
	/*
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
	*/
}

func (gt *gameType) dispatchCommand(tnt *turnNodeType) {
	log.Println("dispatchCommand")
	log.Println(tnt)

	var err error
	var response *ResponseType

	switch tnt.requestCommand {
	case pingRequest:
		response, err = pingReqRes(tnt)
	case playerCreateRequest:
		response, err = commandPlayerCreate(tnt, &gt.players)
	case playerDeleteRequest:
		response, err = commandPlayerDelete(tnt, &gt.players)
		//response, err = commandPlayerDelete(tnt, &gt.board, &gt.ships, &gt.players)
		/*
			case shipCreateCommand:
				commandShipCreate(tnt, &gt.board, &gt.ships)
			case shipDeleteCommand:
				commandShipDelete(tnt, &gt.board, &gt.ships)
		*/

	case shutDownRequest:
		log.Println("shutdown noted")
		gt.shutDownFlag = true
	default:
		log.Println("unknown command")
		log.Println(tnt)
	}

	if err != nil {
		log.Println("error error")
		log.Println(err)
	}

	if response == nil {
		log.Println("xxx xxx xxx nil response")
	} else {
		log.Println("xxx xxx xxx write response")
		responseToManager(gt.outboundQueue, response)
	}
}

func (gt *gameType) serviceRequestQueue() {
	log.Println("serviceRequestQueue entry")

	for {
		// process fresh messages from manager
		rt := gt.requestQueue.dequeue()
		if rt == nil {
			break
		} else {
			tnt := newTurnNode(rt)

			switch tnt.requestCommand {
			case pingRequest:
				gt.dispatchCommand(tnt)
			case playerCreateRequest:
				gt.dispatchCommand(tnt)
			case playerDeleteRequest:
				gt.dispatchCommand(tnt)
			case shipCreateRequest:
				gt.dispatchCommand(tnt)
			case shipDeleteRequest:
				gt.dispatchCommand(tnt)
			case shutDownRequest:
				gt.dispatchCommand(tnt)
			default:
				gt.scheduleTurnEvent(tnt)
			}
		}
	}

	log.Println("serviceRequestQueue exit")
}

func (gt *gameType) servicePlayerTurnQueue(playerNdx int) {
	/*
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
	*/
}

func (gt *gameType) serviceTurnQueue() {
	// need random option

	for ndx := 0; ndx < maxPlayers; ndx++ {
		gt.servicePlayerTurnQueue(ndx)
	}
}

func (gt *gameType) turnManager() {
	gt.turnCounter++
	//gt.eventQueueNdx = gt.turnCounter % maxTurnEventQueue
	//log.Printf("starting turn:%d %d", gt.turnCounter, gt.eventQueueNdx)

	gt.serviceRequestQueue()

	gt.serviceTurnQueue()

	log.Printf("ending turn:%d", gt.turnCounter)
}
