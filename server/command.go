// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

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
	stubCommand0 // command_test.go
	stubCommand1 // command_test.go
	stubCommand2 // command_test.go
	stubCommand3 // command_test.go
	timeCommand  // command_time.go
	torpedoCommand
	tractorCommand
	typeCommand
	unknownCommand
	userCommand // command_user_summary.go
)

type legalGameCommandType struct {
	longName  string
	shortName string
	duration  turnCounterType
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
	//{"ping", "", 0},
	{"planet", "", 1},
	//{"playerCreate", "playerCreate", 0},
	//{"playerDelete", "playerDelete", 0},
	{"point", "", 1},
	//{"pong", "", 0},
	{"radio", "", 1},
	{"repair", "", 1},
	{"scan", "", 1},
	{"set", "", 1},
	{"shields", "", 1},
	//{"shipCreate", "shipCreate", 0},
	//{"shipDelete", "shipDelete", 0},
	//{"shutDown", "shutDown", 0},
	{"status", "st", 1},
	{"summary", "", 1},
	{"target", "", 1},
	{"tell", "", 1},
	{"stub0", "", 0},
	{"stub1", "", 1},
	{"stub2", "", 2},
	{"stub3", "", 3},
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

type commandType struct {
	key                   *tokenKeyType
	command               commandGameEnum
	destinationPlayerKeys playerKeyArrayType
	next                  *commandType
	sourcePlayerKey       *tokenKeyType
	//pointsRequest         *pointsRequestType
	//pointsResponse        *pointsResponseType
	stubRequest  *stubRequestType
	stubResponse *stubResponseType
	timeRequest  *timeRequestType
	timeResponse *timeResponseType
	userRequest  *userRequestType
	userResponse *userResponseType
}

type commandArrayType []*commandType

// convenience factory
func newCommand(command commandGameEnum, playerKey *tokenKeyType) *commandType {
	result := commandType{command: command, key: newTokenKey(""), sourcePlayerKey: playerKey}
	return &result
}

func (gt *gameType) commandDispatch(ct *commandType) {
	switch ct.command {
	//case pointCommand:
	//	gt.pointCommand(ct)
	//case timeCommand:
	//	gt.timeCommand(ct)

	case stubCommand0:
		gt.stubCommand(ct)
	case stubCommand1:
		gt.stubCommand(ct)
	case stubCommand2:
		gt.stubCommand(ct)
	case stubCommand3:
		gt.stubCommand(ct)
	case userCommand:
		gt.userCommand(ct)
	default:
		gt.sugarLog.Infof("unknown command %s", legalGameCommands[ct.command].longName)
	}
}
