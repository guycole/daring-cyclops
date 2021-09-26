// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

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
	pingCommand
	planetsCommand
	playerCreateCommand
	playerDeleteCommand
	pointsCommand
	pongCommand
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
	{"pingCommand", ""},
	{"planet", ""},
	{"playerCreate", "playerCreate"},
	{"playerDelete", "playerDelete"},
	{"points", ""},
	{"pongCommand", ""},
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
