package game

import (
	"log"
)

const maxTeamPlayers = 5
const maxPlayers = maxTeamPlayers * 2

const maxEventQueue = 100

// WorkerType main game structure
type WorkerType struct {
	//gameBoard   GameBoardType
	//players     [maxPlayers]*PlayerType
	turnCounter int
	//eventQueue  *turnEventType
	eventQueue    [maxEventQueue]turnEventType
	eventQueueNdx int
	uuid          string
}

// turnEventType must be sorted by turn
type turnEventType struct {
	payload *commandType
}

func eventQueuePop(wt *WorkerType) *commandType {
	return nil
	/*
			payload := wt.eventQueue.payload
			if payload == nil {
				// all commands for this turn have been consumed
				if wt.eventQueue.next == nil {
					wt.eventQueue = nil
				} else {
					wt.eventQueue = wt.eventQueue.next
				}
			} else {
				wt.eventQueue.payload = payload.next
				payload.next = nil
			}

		return payload
	*/
}

func eventQueuePush(ct *commandType, wt *WorkerType) {
	xxx := ct.turn % maxEventQueue
	log.Println(xxx)

	/*
		if wt.eventQueue == nil {
			eq := newTurnEvent(ct.turn)
			eq.payload = ct
			wt.eventQueue = eq
			return
		}
	*/

	/*
		if ct.turn < wt.eventQueue.turn {
			eq = newTurnEvent(ct.turn)
			eq.payload = ct
			eq.next = wt.eventQueue
			wt.eventQueue = eq.next
		}

		for ndx := 0; xxx; ndx++ {
		}
	*/
	log.Println("kill me")
}

// NewGame creates new game (should be init)
func NewWorker(id string) *WorkerType {
	log.Println("new game:", id)

	wt := WorkerType{uuid: id}
	//	result.gameBoard = freshGameBoard()

	for ndx := 0; ndx < maxEventQueue; ndx++ {
		wt.eventQueue[ndx] = turnEventType{}
	}

	return &wt
}

// TurnManager manage game play
func TurnManager(wt *WorkerType) {
	wt.turnCounter += 1
	wt.eventQueueNdx = wt.turnCounter % maxEventQueue
	log.Printf("starting turn:%d %d", wt.turnCounter, wt.eventQueueNdx)

	serviceInboundQueue(wt)
	serviceEventQueue(wt)

	log.Printf("ending turn:%d", wt.turnCounter)
}

// serviceEventQueue dispatch events
func serviceEventQueue(wt *WorkerType) {
	log.Printf("serviceEventQueue:%d", wt.turnCounter)

	/*
		for ndx := 0; ndx <= wt.turnCounter; ndx++ {
			if wt.eventQueue == nil || wt.eventQueue.turn > wt.turnCounter {
				log.Println("skipping empty event queue")
				break
			} else {
				command := eventQueuePop(wt)
				log.Println(command)
				break
			}
		}
	*/
}

// serviceInboundQueue read from RabbitMQ and add to event queue
func serviceInboundQueue(wt *WorkerType) {
	log.Printf("serviceInboundQueue:%d", wt.turnCounter)

	ct := newCommand("aaa", "bbb", 3)
	eventQueuePush(ct, wt)
}

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
