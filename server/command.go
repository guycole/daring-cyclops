// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"strings"

	"github.com/google/uuid"
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
	//pingCommand
	planetsCommand
	//playerCreateCommand
	//playerDeleteCommand
	pointCommand // command_point.go
	//pongCommand
	radioCommand
	repairCommand
	scanCommand
	setCommand
	shieldsCommand
	//shipCreateCommand
	//shipDeleteCommand
	//shutDownCommand
	statusCommand
	summaryCommand
	targetsCommand
	tellCommand
	timeCommand // command_time.go
	torpedoCommand
	tractorCommand
	typeCommand
	unknownCommand
	userCommand // command_user_summary.go
)

type legalGameCommandType struct {
	longName  string
	shortName string
	duration  int
}

var legalGameCommands = [...]legalGameCommandType{
	{"bases", "ba", 1},
	{"build", "bu", 1},
	{"capture", "ca", 1},
	//	{"chronicles", ""},
	{"damages", "", 1},
	{"dock", "", 1},
	{"drop", "", 1},
	{"energy", "", 1},
	{"exit", "", 1},
	{"gate", "", 1},
	{"gripe", "", 1},
	{"help", "", 1},
	{"history", "", 1},
	//	{"honor", ""},
	{"impulse", "", 1},
	{"list", "", 1},
	{"move", "m", 3},
	{"news", "", 1},
	{"phasers", "", 1},
	{"ping", "", 0},
	{"planet", "", 1},
	{"playerCreate", "playerCreate", 0},
	{"playerDelete", "playerDelete", 0},
	{"point", "", 1},
	{"pong", "", 0},
	{"radio", "", 1},
	{"repair", "", 1},
	{"scan", "", 1},
	{"set", "", 1},
	{"shields", "", 1},
	{"shipCreate", "shipCreate", 0},
	{"shipDelete", "shipDelete", 0},
	{"shutDown", "shutDown", 0},
	{"status", "st", 1},
	{"summary", "", 1},
	{"target", "", 1},
	{"tell", "", 1},
	{"time", "", 1},
	{"torpedo", "", 1},
	{"tractor", "", 1},
	{"type", "", 1},
	{"unknownCommand", "unknownCommand", 1},
	{"user", "", 1},
}

func findGameCommand(arg string) commandGameEnum {
	for ndx := 0; ndx < len(legalGameCommands); ndx++ {
		if legalGameCommands[ndx].longName == arg || legalGameCommands[ndx].shortName == arg {
			return commandGameEnum(ndx)
		}
	}

	return commandGameEnum(unknownCommand)
}

type commandKeyType struct {
	key string
}

// convenience factory
func newCommandKey(key string) *commandKeyType {
	var result commandKeyType

	temp := strings.TrimSpace(key)
	if len(temp) < 36 {
		result = commandKeyType{key: uuid.NewString()}
	} else {
		result = commandKeyType{key: temp}
	}

	return &result
}

type commandType struct {
	key                   *commandKeyType
	command               commandGameEnum
	sourcePlayerKey       *playerKeyType
	destinationPlayerKeys playerKeyArrayType
	userRequest           *userRequestType
	userResponse          *userResponseType
}

type commandArrayType []*commandType

// convenience factory
func newCommand(command commandGameEnum, playerKey *playerKeyType) *commandType {
	result := commandType{command: command, key: newCommandKey(""), sourcePlayerKey: playerKey}
	return &result
}

func (gt *gameType) commandDispatch(ct *commandType) {
	switch ct.command {
	//case pointCommand:
	//	gt.pointCommand(ct)
	//case timeCommand:
	//	gt.timeCommand(ct)
	case userCommand:
		gt.userCommand(ct)
	default:
		gt.sugarLog.Info("unknown command")
	}
}
