package main

import (
	"context"
	"encoding/json"
	"log"

	redis "github.com/go-redis/redis/v8"
)

type commandGameEnum int

// must match order for legalGameCommands
const (
	basesCommand commandGameEnum = iota
	buildCommand
	captureCommand
	//	chroniclesCommand
	damagesCommand
	dockCommand
	dropCommand
	energyCommand
	exitCommand
	gateCommand
	gripeCommand
	helpCommand
	historyCommand
	//	honorCommand
	impulseCommand
	listCommand
	moveCommand
	newsCommand
	phasersCommand
	planetsCommand
	playerCreateCommand
	playerDeleteCommand
	pointsCommand
	radioCommand
	repairCommand
	scanCommand
	setCommand
	shieldsCommand
	shipCreateCommand
	shipDeleteCommand
	statusCommand
	summaryCommand
	targetsCommand
	tellCommand
	timeCommand
	torpedoCommand
	tractorCommand
	typeCommand
	unknownCommand
	usersCommand
)

// must match order for commandGameEnum
var legalGameCommands = [...][2]string{
	{"bases", "ba"},
	{"build", "bu"},
	{"capture", "ca"},
	//	{"chronicles", ""},
	{"damages", ""},
	{"dock", ""},
	{"drop", ""},
	{"energy", ""},
	{"exit", ""},
	{"gate", ""},
	{"gripe", ""},
	{"help", ""},
	{"history", ""},
	//	{"honor", ""},
	{"impulse", ""},
	{"list", ""},
	{"move", "m"},
	{"news", ""},
	{"phasers", ""},
	{"planet", ""},
	{"playerCreate", "playerCreate"},
	{"playerDelete", "playerDelete"},
	{"points", ""},
	{"radio", ""},
	{"repair", ""},
	{"scan", ""},
	{"set", ""},
	{"shields", ""},
	{"shipCreate", "shipCreate"},
	{"shipDelete", "shipDelete"},
	{"status", "st"},
	{"summary", ""},
	{"target", ""},
	{"tell", ""},
	{"time", ""},
	{"torpedo", ""},
	{"tractor", ""},
	{"type", ""},
	{"unknownCommand", "unknownCommand"},
	{"users", ""},
}

func findGameCommand(arg string) commandGameEnum {
	for ndx := 0; ndx < len(legalGameCommands); ndx++ {
		if legalGameCommands[ndx][0] == arg || legalGameCommands[ndx][1] == arg {
			return commandGameEnum(ndx)
		}
	}

	return commandGameEnum(unknownCommand)
}

func findCommandDuration(arg commandGameEnum) int {
	// TODO
	return 1
}

///////////////
// process fresh command
///////////////

type eventType struct {
	name    string // player name
	request string // request uuid

	duration int // command duration (in turns)
	turn     int // turn counter for command execution

	commands commandArrayType

	command commandGameEnum

	next *eventType
}

func newEvent(ct CommandType) (*eventType, error) {
	result := eventType{name: ct.Name, request: ct.RequestId}

	return &result, nil
}

func eventQueueAdapter(ct CommandType) {
	log.Println(ct)
	et, err := newEvent(ct)
	if err != nil {
		log.Println("err err err err 333")
		log.Println(err)
	}
	log.Println(et)
}

///////////////
// command from manager
///////////////

const maxCommandArguments = 5

type commandArrayType [maxCommandArguments]string

type CommandType struct {
	Name        string
	RequestId   string
	CommandSize int
	Commands    commandArrayType
}

func commandFromManager(channelName string) {
	log.Println("commandFromManager entry")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	topic := rdb.Subscribe(context.Background(), channelName)

	for {
		// blocking read
		msg, err := topic.ReceiveMessage(context.Background())
		if err != nil {
			log.Println("err err err err")
			log.Println(err)
			continue
		}

		var ct CommandType
		err = json.Unmarshal([]byte(msg.Payload), &ct)
		if err != nil {
			log.Println("err err err err 222")
			log.Println(err)
			continue
		}

		eventQueueAdapter(ct)
	}

	log.Println("commandFromManager exit")
}

/*
// commandType single linked list of commands (for each event)
type commandType struct {
	player  string // player uuid
	raw     string // original json
	request string // request uuid

	duration int // command duration (in turns)
	turn     int // turn counter for command execution

	args    []string
	command commandGameEnum

	next *commandType
}
*/

/*
// parseJsonCommand parse fresh command
func parseJsonCommand(raw string, tc int) *commandType {
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(raw), &jsonMap)

	player := jsonMap["player"].(string)
	request := jsonMap["requestId"].(string)

	tempArray := jsonMap["command"].([]interface{})
	argArray := make([]string, len(tempArray))
	for key, value := range tempArray {
		argArray[key] = value.(string)
	}

	command := findGameCommand(argArray[0])
	if command == unknownCommand {
		log.Printf("unknownCommand:%s", argArray[0])
		return nil
	}

	result := commandType{raw: raw, player: player, request: request}
	result.args = argArray
	result.command = command
	result.duration = findCommandDuration(command)
	result.turn = result.duration + tc

	return &result
}
*/

/*
func dispatchCommand(command commandType, gt *gameType) {
	switch command.command {
	case moveCommand:
		log.Println("move noted")
		commandMoveShip(command, gt)
	case playerCreateCommand:
		log.Println("create player noted")
		commandPlayerCreate(command, gt)
	case playerDeleteCommand:
		log.Println("delete player noted")
		commandPlayerDelete(command, gt)
	case shipCreateCommand:
		log.Println("create ship noted")
		commandShipCreate(command, gt)
	case shipDeleteCommand:
		log.Println("delete ship noted")
		commandShipDelete(command, gt)
	default:
		log.Println("unknown command")
	}
}
*/
