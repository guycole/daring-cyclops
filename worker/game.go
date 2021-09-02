package main

import (
	"log"
)

const maxTeamPlayers = 5
const maxPlayers = maxTeamPlayers * 2

const maxTeamShips = 5
const maxShips = maxTeamShips * 2

const maxEventQueue = 10

// gameType main game structure, only one instance per game
type gameType struct {
	//gameBoard   GameBoardType

	players [maxPlayers]playerType
	ships   [maxShips]shipType

	eventQueue    [maxEventQueue]turnEventType
	eventQueueNdx int // current queue index
	turnCounter   int // current game turn

	uuid string // game identifier
}

type turnEventType struct {
	payload *commandType // single linked list of commands in execution order
}

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
}

// eventQueueDump writes event queue to stdout
func eventQueueDump(gt gameType) {
	log.Println("=-=-=-= eventQueueDump =-=-=-=")

	for ndx := 0; ndx < maxEventQueue; ndx++ {
		temp := gt.eventQueue[ndx].payload

		for {
			if temp == nil {
				log.Printf("%d nil", ndx)
				break
			} else {
				log.Printf("%d %s %s", ndx, temp.player, temp.raw)
				temp = temp.next
			}
		}
	}

	log.Println("=-=-=-= eventQueueDump =-=-=-=")
}

// eventQueuePop consume event from queue
func eventQueuePop(gt *gameType) *commandType {
	payload := gt.eventQueue[gt.eventQueueNdx].payload
	if payload != nil {
		gt.eventQueue[gt.eventQueueNdx].payload = payload.next
	}

	return payload
}

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

func newGame(id string) *gameType {
	log.Println("new game:", id)

	gt := gameType{uuid: id}

	//	result.gameBoard = freshGameBoard()

	playerDump(gt)

	return &gt
}

// turnManager schedules game play
func turnManager(gt *gameType) {
	gt.turnCounter++
	gt.eventQueueNdx = gt.turnCounter % maxEventQueue
	log.Printf("starting turn:%d %d", gt.turnCounter, gt.eventQueueNdx)

	serviceInboundQueue(gt)
	serviceEventQueue(gt)

	log.Printf("ending turn:%d", gt.turnCounter)
}

// serviceInboundQueue read from RabbitMQ and add to event queue
func serviceInboundQueue(gt *gameType) {
	log.Printf("serviceInboundQueue:%d", gt.turnCounter)

	if gt.turnCounter == 1 {
		eventQueueSimulator(gt)
		eventQueueDump(*gt)
	}
}

// serviceEventQueue dispatch events
func serviceEventQueue(gt *gameType) {
	log.Printf("serviceEventQueue:%d", gt.turnCounter)

	for {
		current := eventQueuePop(gt)
		if current == nil {
			break
		} else {
			dispatchCommand(*current, gt)
		}
	}
}

/////////// kill below

/*
// PlayerAdd add fresh player to game
func PlayerAdd(pt *PlayerType, gt *GameType) {
	log.Println(pt)

	switch pt.team {
	case BlueTeam:
		for ndx := 0; ndx < maxTeamPlayers; ndx++ {
			if gt.playersBlue[ndx] == nil {
				log.Println("assign", ndx)
				gt.playersBlue[ndx] = pt
				return
			}
		}

		// TOODO throw
		log.Println("team full blue")
	case RedTeam:
		for ndx := 0; ndx < maxTeamPlayers; ndx++ {
			if gt.playersRed[ndx] == nil {
				log.Println("assign", ndx)
				gt.playersRed[ndx] = pt
				return
			}
		}

		//TODO throw
		log.Println("team full red")
	default:
		log.Println("unkown team")
	}
}
*/

/*
// PlayerDelete remove player from game
func PlayerDelete(target string, gt *GameType) {
	for ndx := 0; ndx < maxTeamPlayers; ndx++ {
		if strings.Compare(gt.playersBlue[ndx].uuid, target) == 0 {
			gt.playersBlue[ndx] = nil
			return
		}
	}

	log.Println("no match blue")

	for ndx := 0; ndx < maxTeamPlayers; ndx++ {
		if strings.Compare(gt.playersRed[ndx].uuid, target) == 0 {
			gt.playersRed[ndx] = nil
			return
		}
	}

	log.Println("no match red")
}
*/

/*
// PlayerFind remove player from game
func PlayerFind(target string, gt *GameType) *PlayerType {
	for ndx := 0; ndx < maxTeamPlayers; ndx++ {
		if strings.Compare(gt.playersBlue[ndx].uuid, target) == 0 {
			return gt.playersBlue[ndx]
		}
	}

	log.Println("no match blue")

	for ndx := 0; ndx < maxTeamPlayers; ndx++ {
		if strings.Compare(gt.playersRed[ndx].uuid, target) == 0 {
			return gt.playersRed[ndx]
		}
	}

	log.Println("no match red")
	return nil
}
*/

/*
// ShipAdd add fresh ship to game
func ShipAdd(st *ShipType, gt *GameType) {
	// find player
	// discover if legal ship type for rank
	// discover if legal ship population
	// discover if ship name available
	// add ship to hashmap
}

// ShipDelete remove ship from game
func ShipDelete(target string, gt *GameType) {
}
*/
