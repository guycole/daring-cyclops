package main

import (
	"log"
)

// commandType single linked list of commands (for each event)
type commandType struct {
	payload string // original json
	player  string // player uuid
	turn    int
	next    *commandType
}

func newCommand(payload, player string, currentTurn int) *commandType {
	result := commandType{payload: payload, player: player}
	return &result
}

type commandGameEnum int

// must match order for legalGameCommands
const (
	basesCommand commandGameEnum = iota
	buildCommand
	captureCommand
	chroniclesCommand
	createUserCommand
	damagesCommand
	dockCommand
	dropCommand
	energyCommand
	exitCommand
	gateCommand
	gripeCommand
	helpCommand
	historyCommand
	honorCommand
	impulseCommand
	listCommand
	moveCommand
	newsCommand
	phasersCommand
	planetsCommand
	pointsCommand
	radioCommand
	repairCommand
	scanCommand
	setCommand
	shieldsCommand
	statusCommand
	summaryCommand
	targetsCommand
	tellCommand
	timeCommand
	torpedoCommand
	tractorCommand
	typeCommand
	usersCommand
)

// must match order for commandGameEnum
var legalGameCommands = [...][2]string{
	{"bases", "ba"},
	{"build", "bu"},
	{"capture", "ca"},
	{"chronicles", ""},
	{"createUser", "createUser"},
	{"damages", ""},
	{"dock", ""},
	{"drop", ""},
	{"energy", ""},
	{"exit", ""},
	{"gate", ""},
	{"gripe", ""},
	{"help", ""},
	{"history", ""},
	{"honor", ""},
	{"impulse", ""},
	{"list", ""},
	{"move", "m"},
	{"news", ""},
	{"phasers", ""},
	{"planet", ""},
	{"points", ""},
	{"radio", ""},
	{"repair", ""},
	{"scan", ""},
	{"set", ""},
	{"shields", ""},
	{"status", "st"},
	{"summary", ""},
	{"target", ""},
	{"tell", ""},
	{"time", ""},
	{"torpedo", ""},
	{"tractor", ""},
	{"type", ""},
	{"users", ""},
}

// CommandGameType ryryr
type commandGameType struct {
	command  commandGameEnum
	duration int
	user     string
}

func findGameCommand(arg string) int {
	for ndx := 0; ndx < len(legalGameCommands); ndx++ {
		if legalGameCommands[ndx][0] == arg || legalGameCommands[ndx][1] == arg {
			return ndx
		}
	}

	return -1
}

// NewJsonCommand ryryr
func newJsonCommand(command string) *commandGameType {
	log.Println(command)
	return nil

	/*
		commandNdx := findGameCommand(command)
		if commandNdx < 0 {
			log.Println("error error error")
		}

		result := CommandGameType{user: user}
		result.command = commandGameEnum(commandNdx)

		return &result
	*/
}

/*
// NewTextCommand ryryry
func NewTextCommand(command, user string) *CommandType {
	commandNdx := findCommand(command)

	if commandNdx < 0 {
		log.Println("error error error")
	}

	result := CommandType{user: user}
	result.command = commandEnum(commandNdx)

	return &result
}
*/

func unknownCommand() {
	log.Println("unknownCommand")
}

func dispatchCommand(command commandType, gt *gameType) {
	log.Println(command)
}

/*
func dispatchCommand2(command string, gt gameType) {
	log.Println(command)

	var result map[string]interface{}
	json.Unmarshal([]byte(command), &result)
	args := result["command"]

	log.Println(result)
	log.Println(result["command"])
	log.Println(args)
	//log.Println(args[0])
}
*/

/*
// DispatchCommand ryryry
func DispatchCommand(command *CommandType, game WorkerType) {
	log.Println("dispatch command")

	switch command.command {
	case 0: // bases
		unknownCommand()
	case 1: // build
		unknownCommand()
	case 2: // capture
		unknownCommand()
	case 3: // chronicles
		unknownCommand()
	case 4: // damages
		unknownCommand()
	case 5: // dock
		unknownCommand()
	case 6: // drop
		unknownCommand()
	case 7: // energy
		unknownCommand()
	case 8: // exit
		unknownCommand()
	case 9: // gate
		unknownCommand()
	case 10: // gripe
		unknownCommand()
	case 11: // help
		unknownCommand()
	case 12: // history
		unknownCommand()
	case 13: // honor
		unknownCommand()
	case 14: // impulse
		unknownCommand()
	case 15: // list
		unknownCommand()
	case 16: // move
		unknownCommand()
	case 17: // news
		unknownCommand()
	case 18: // phasers
		unknownCommand()
	case 19: // planet
		unknownCommand()
	case 20: // points
		unknownCommand()
	case 21: // radio
		unknownCommand()
	case 22: // repair
		unknownCommand()
	case 23: // scan
		unknownCommand()
	case 24: // set
		unknownCommand()
	case 25: // shields
		unknownCommand()
	case 26: // status
		unknownCommand()
	case 27: // summary
		unknownCommand()
	case 28: // target
		unknownCommand()
	case 29: // tell
		unknownCommand()
	case 30: // time
		unknownCommand()
	case 31: // torpedo
		unknownCommand()
	case 32: // tractor
		unknownCommand()
	case 33: // type
		unknownCommand()
	case 34: // users
		unknownCommand()
	default:
		unknownCommand()
	}
}
*/
