package main

import (
	"encoding/json"
	"log"
)

type commandGameEnum int

// must match order for legalGameCommands
const (
	basesCommand commandGameEnum = iota
	buildCommand
	captureCommand
	//	chroniclesCommand
	createPlayerCommand
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
	unknownCommand
	usersCommand
)

// must match order for commandGameEnum
var legalGameCommands = [...][2]string{
	{"bases", "ba"},
	{"build", "bu"},
	{"capture", "ca"},
	//	{"chronicles", ""},
	{"createPlayer", "createPlayer"},
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

func dispatchCommand(command commandType, gt *gameType) {
	switch command.command {
	case createPlayerCommand:
		log.Println("create player noted")
		commandCreatePlayer(command, gt)
	default:
		log.Println("unknown command")
	}
}
