package main

import (
	"log"
)

const maxTeamPlayers = 5
const maxPlayers = maxTeamPlayers * 2

const maxEventQueue = 10

// gameType main game structure, only one instance per game
type gameType struct {
	//gameBoard   GameBoardType
	//players     [maxPlayers]*PlayerType
	turnCounter int
	//eventQueue  *turnEventType
	eventQueue    [maxEventQueue]turnEventType
	eventQueueNdx int
	uuid          string
}

type turnEventType struct {
	payload *commandType // single linked list of commands in execution order
}

func eventQueueSimulator(gt *gameType) {
	log.Println("simulator")

	temp1 := newCommand("tell all byte me 111", player1, 0)
	log.Println(temp1)
	eventQueuePush(*temp1, gt)

	temp2 := newCommand("tell all byte me 222", player1, 0)
	log.Println(temp2)
	eventQueuePush(*temp2, gt)

	temp3 := newCommand("tell all byte me 333", player1, 0)
	log.Println(temp3)
	eventQueuePush(*temp3, gt)

	message1 := `{"command":[ "commandUuid", "player1uuid", "createUser", "CaptainRank", "BlueTeam"]}`
	log.Println(message1)
	temp4 := newJsonCommand(message1)
	if temp4 == nil {

	}

	message2 := `{"command":[ "commandUuid", "player1uuid", "createShip", "player1uuid", "nimrod"]}`
	log.Println(message2)
	temp5 := newJsonCommand(message2)
	if temp5 == nil {

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
				log.Printf("%d %s %s", ndx, temp.player, temp.payload)
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
